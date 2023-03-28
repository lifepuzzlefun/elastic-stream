package cluster

import (
	"context"
	"sort"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/AutoMQ/placement-manager/api/rpcfb/rpcfb"
	"github.com/AutoMQ/placement-manager/pkg/util/traceutil"
)

var (
	// ErrNotEnoughDataNodes is returned when there are not enough data nodes to allocate a range.
	ErrNotEnoughDataNodes = errors.New("not enough data nodes")
)

// Heartbeat updates DataNode's last active time, and save it to storage if its info changed.
func (c *RaftCluster) Heartbeat(ctx context.Context, node *rpcfb.DataNodeT) error {
	logger := c.lg.With(traceutil.TraceLogField(ctx))

	updated := c.cache.SaveDataNode(node)
	if updated {
		logger.Info("data node updated, start to save it", zap.Int32("node-id", node.NodeId), zap.String("advertise-addr", node.AdvertiseAddr))
		_, err := c.storage.SaveDataNode(ctx, node)
		logger.Info("finish saving data node", zap.Int32("node-id", node.NodeId), zap.Error(err))
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *RaftCluster) AllocateID(ctx context.Context) (int32, error) {
	logger := c.lg.With(traceutil.TraceLogField(ctx))

	id, err := c.dnAlloc.Alloc(ctx)
	if err != nil {
		logger.Error("failed to allocate data node id", zap.Error(err))
		return -1, err
	}

	return int32(id), nil
}

// chooseDataNodes selects `cnt` number of data nodes from the available data nodes for a range.
// Only DataNodeT.NodeId is filled in the returned ReplicaNodeT.
// It returns ErrNotEnoughDataNodes if there are not enough data nodes to allocate.
func (c *RaftCluster) chooseDataNodes(cnt int8) ([]*rpcfb.ReplicaNodeT, error) {
	if int(cnt) > c.cache.DataNodeCount() {
		return nil, errors.Wrapf(ErrNotEnoughDataNodes, "required %d, available %d", cnt, c.cache.DataNodeCount())
	}

	nodes := c.cache.DataNodes()
	// TODO more intelligent selection
	sort.Slice(nodes, func(i, j int) bool {
		return !nodes[i].LastActiveTime.Before(nodes[j].LastActiveTime)
	})

	replicaNodes := make([]*rpcfb.ReplicaNodeT, 0, cnt)
	for i := 0; i < int(cnt); i++ {
		replicaNodes = append(replicaNodes, &rpcfb.ReplicaNodeT{
			DataNode: &rpcfb.DataNodeT{
				NodeId: nodes[i].NodeId,
			},
		})
	}
	replicaNodes[0].IsPrimary = true

	return replicaNodes, nil
}

func (c *RaftCluster) eraseDataNodeInfo(node *rpcfb.DataNodeT) {
	node.AdvertiseAddr = ""
}

func (c *RaftCluster) fillDataNodeInfo(node *rpcfb.DataNodeT) {
	n := c.cache.DataNode(node.NodeId)
	if n == nil {
		c.lg.Warn("data node not found", zap.Int32("node-id", node.NodeId))
		return
	}
	node.AdvertiseAddr = n.AdvertiseAddr
}
