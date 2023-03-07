// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package header

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TrimStreamsResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsTrimStreamsResponse(buf []byte, offset flatbuffers.UOffsetT) *TrimStreamsResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TrimStreamsResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsTrimStreamsResponse(buf []byte, offset flatbuffers.UOffsetT) *TrimStreamsResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &TrimStreamsResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *TrimStreamsResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TrimStreamsResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TrimStreamsResponse) ThrottleTimeMs() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TrimStreamsResponse) MutateThrottleTimeMs(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *TrimStreamsResponse) TrimResponses(obj *TrimStreamResult, j int) bool {
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

func (rcv *TrimStreamsResponse) TrimResponsesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *TrimStreamsResponse) ErrorCode() ErrorCode {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return ErrorCode(rcv._tab.GetInt16(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *TrimStreamsResponse) MutateErrorCode(n ErrorCode) bool {
	return rcv._tab.MutateInt16Slot(8, int16(n))
}

func (rcv *TrimStreamsResponse) ErrorMessage() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func TrimStreamsResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func TrimStreamsResponseAddThrottleTimeMs(builder *flatbuffers.Builder, throttleTimeMs int32) {
	builder.PrependInt32Slot(0, throttleTimeMs, 0)
}
func TrimStreamsResponseAddTrimResponses(builder *flatbuffers.Builder, trimResponses flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(trimResponses), 0)
}
func TrimStreamsResponseStartTrimResponsesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TrimStreamsResponseAddErrorCode(builder *flatbuffers.Builder, errorCode ErrorCode) {
	builder.PrependInt16Slot(2, int16(errorCode), 0)
}
func TrimStreamsResponseAddErrorMessage(builder *flatbuffers.Builder, errorMessage flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(errorMessage), 0)
}
func TrimStreamsResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
