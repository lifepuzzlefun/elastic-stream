// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package header

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DescribeStreamsResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsDescribeStreamsResponse(buf []byte, offset flatbuffers.UOffsetT) *DescribeStreamsResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DescribeStreamsResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDescribeStreamsResponse(buf []byte, offset flatbuffers.UOffsetT) *DescribeStreamsResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DescribeStreamsResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DescribeStreamsResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DescribeStreamsResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DescribeStreamsResponse) ThrottleTimeMs() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DescribeStreamsResponse) MutateThrottleTimeMs(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *DescribeStreamsResponse) DescribeResponses(obj *DescribeStreamResult, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *DescribeStreamsResponse) DescribeResponsesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *DescribeStreamsResponse) ErrorCode() ErrorCode {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return ErrorCode(rcv._tab.GetInt16(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *DescribeStreamsResponse) MutateErrorCode(n ErrorCode) bool {
	return rcv._tab.MutateInt16Slot(8, int16(n))
}

func (rcv *DescribeStreamsResponse) ErrorMessage() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func DescribeStreamsResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func DescribeStreamsResponseAddThrottleTimeMs(builder *flatbuffers.Builder, throttleTimeMs int32) {
	builder.PrependInt32Slot(0, throttleTimeMs, 0)
}
func DescribeStreamsResponseAddDescribeResponses(builder *flatbuffers.Builder, describeResponses flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(describeResponses), 0)
}
func DescribeStreamsResponseStartDescribeResponsesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DescribeStreamsResponseAddErrorCode(builder *flatbuffers.Builder, errorCode ErrorCode) {
	builder.PrependInt16Slot(2, int16(errorCode), 0)
}
func DescribeStreamsResponseAddErrorMessage(builder *flatbuffers.Builder, errorMessage flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(errorMessage), 0)
}
func DescribeStreamsResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
