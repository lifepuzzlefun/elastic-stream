// Copyright 2016 TiKV Project Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cluster

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/AutoMQ/pd/api/rpcfb/rpcfb"
	sbpClient "github.com/AutoMQ/pd/pkg/sbp/client"
	"github.com/AutoMQ/pd/pkg/server/cluster/cache"
	"github.com/AutoMQ/pd/pkg/server/config"
	"github.com/AutoMQ/pd/pkg/server/id"
	"github.com/AutoMQ/pd/pkg/server/member"
	"github.com/AutoMQ/pd/pkg/server/model"
	"github.com/AutoMQ/pd/pkg/server/storage"
)

const (
	_streamIDAllocKey      = "stream"
	_streamIDStep          = 100
	_rangeServerIDAllocKey = "range-server"
	_rangeServerIDStep     = 1
	_objectIDAllocKey      = "object"
	_objectIDStep          = 100
)

// RaftCluster is used for metadata management.
type RaftCluster struct {
	ctx context.Context

	cfg *config.Cluster

	starting      atomic.Bool
	running       atomic.Bool
	runningCtx    context.Context
	runningCancel context.CancelFunc

	storage        storage.Storage
	sAlloc         id.Allocator // stream id allocator
	rsAlloc        id.Allocator // range server id allocator
	oAlloc         id.Allocator // object id allocator
	member         MemberService
	cache          *cache.Cache
	rangeServerIdx atomic.Uint64
	client         sbpClient.Client

	lg *zap.Logger
}

type MemberService interface {
	IsLeader() bool
	// ClusterInfo returns all members in the cluster.
	ClusterInfo(ctx context.Context) ([]*member.Info, error)
}

// Server is the interface for starting a RaftCluster.
type Server interface {
	Storage() storage.Storage
	IDAllocator(key string, start, step uint64) id.Allocator
	Member() MemberService
	SbpClient() sbpClient.Client
}

// NewRaftCluster creates a new RaftCluster.
func NewRaftCluster(ctx context.Context, cfg *config.Cluster, member MemberService, logger *zap.Logger) *RaftCluster {
	return &RaftCluster{
		ctx:    ctx,
		cfg:    cfg,
		member: member,
		cache:  cache.NewCache(),
		lg:     logger,
	}
}

// Start starts the RaftCluster.
func (c *RaftCluster) Start(s Server) error {
	logger := c.lg
	if c.IsRunning() {
		logger.Warn("raft cluster is already running")
		return nil
	}
	if c.starting.Swap(true) {
		logger.Warn("raft cluster is starting")
		return nil
	}
	defer c.starting.Store(false)

	logger.Info("starting raft cluster")

	c.storage = s.Storage()
	c.runningCtx, c.runningCancel = context.WithCancel(c.ctx)
	c.sAlloc = s.IDAllocator(_streamIDAllocKey, uint64(model.MinStreamID), _streamIDStep)
	c.rsAlloc = s.IDAllocator(_rangeServerIDAllocKey, uint64(model.MinRangeServerID), _rangeServerIDStep)
	c.oAlloc = s.IDAllocator(_objectIDAllocKey, uint64(model.MinObjectID), _objectIDStep)
	c.client = s.SbpClient()

	err := c.loadInfo()
	if err != nil {
		logger.Error("load cluster info failed", zap.Error(err))
		return errors.WithMessage(err, "load cluster info")
	}

	// start other background goroutines

	if c.running.Swap(true) {
		logger.Warn("raft cluster is already running")
		return nil
	}
	return nil
}

// loadInfo loads all info from storage into cache.
func (c *RaftCluster) loadInfo() error {
	logger := c.lg

	c.cache.Reset()

	// load range servers
	start := time.Now()
	err := c.storage.ForEachRangeServer(c.ctx, func(serverT *rpcfb.RangeServerT) error {
		updated, old := c.cache.SaveRangeServer(&cache.RangeServer{
			RangeServerT: *serverT,
		})
		if updated {
			logger.Warn("different range server in storage and cache", zap.Any("range-server-in-storage", serverT), zap.Any("range-server-in-cache", old))
		}
		return nil
	})
	if err != nil {
		return errors.WithMessage(err, "load range servers")
	}
	logger.Info("load range servers", zap.Int("count", c.cache.RangeServerCount()), zap.Duration("cost", time.Since(start)))

	return nil
}

// Stop stops the RaftCluster.
func (c *RaftCluster) Stop() error {
	logger := c.lg
	if !c.running.Swap(false) {
		logger.Warn("raft cluster has already been stopped")
		return nil
	}

	logger.Info("stopping cluster")
	c.runningCancel()

	logger.Info("cluster stopped")
	return nil
}

// IsRunning returns true if the RaftCluster is running.
func (c *RaftCluster) IsRunning() bool {
	return c.running.Load()
}

func (c *RaftCluster) IsLeader() bool {
	return c.member.IsLeader()
}

func (c *RaftCluster) ClusterInfo(ctx context.Context) ([]*member.Info, error) {
	return c.member.ClusterInfo(ctx)
}
