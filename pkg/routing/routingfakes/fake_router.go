// Code generated by counterfeiter. DO NOT EDIT.
package routingfakes

import (
	"context"
	"sync"

	"github.com/abdulhaseeb08/livekit-server/pkg/routing"
	"github.com/livekit/protocol/livekit"
)

type FakeRouter struct {
	ClearRoomStateStub        func(context.Context, livekit.RoomName) error
	clearRoomStateMutex       sync.RWMutex
	clearRoomStateArgsForCall []struct {
		arg1 context.Context
		arg2 livekit.RoomName
	}
	clearRoomStateReturns struct {
		result1 error
	}
	clearRoomStateReturnsOnCall map[int]struct {
		result1 error
	}
	DrainStub        func()
	drainMutex       sync.RWMutex
	drainArgsForCall []struct {
	}
	GetNodeForRoomStub        func(context.Context, livekit.RoomName) (*livekit.Node, error)
	getNodeForRoomMutex       sync.RWMutex
	getNodeForRoomArgsForCall []struct {
		arg1 context.Context
		arg2 livekit.RoomName
	}
	getNodeForRoomReturns struct {
		result1 *livekit.Node
		result2 error
	}
	getNodeForRoomReturnsOnCall map[int]struct {
		result1 *livekit.Node
		result2 error
	}
	GetRegionStub        func() string
	getRegionMutex       sync.RWMutex
	getRegionArgsForCall []struct {
	}
	getRegionReturns struct {
		result1 string
	}
	getRegionReturnsOnCall map[int]struct {
		result1 string
	}
	ListNodesStub        func() ([]*livekit.Node, error)
	listNodesMutex       sync.RWMutex
	listNodesArgsForCall []struct {
	}
	listNodesReturns struct {
		result1 []*livekit.Node
		result2 error
	}
	listNodesReturnsOnCall map[int]struct {
		result1 []*livekit.Node
		result2 error
	}
	OnNewParticipantRTCStub        func(routing.NewParticipantCallback)
	onNewParticipantRTCMutex       sync.RWMutex
	onNewParticipantRTCArgsForCall []struct {
		arg1 routing.NewParticipantCallback
	}
	OnRTCMessageStub        func(routing.RTCMessageCallback)
	onRTCMessageMutex       sync.RWMutex
	onRTCMessageArgsForCall []struct {
		arg1 routing.RTCMessageCallback
	}
	RegisterNodeStub        func() error
	registerNodeMutex       sync.RWMutex
	registerNodeArgsForCall []struct {
	}
	registerNodeReturns struct {
		result1 error
	}
	registerNodeReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveDeadNodesStub        func() error
	removeDeadNodesMutex       sync.RWMutex
	removeDeadNodesArgsForCall []struct {
	}
	removeDeadNodesReturns struct {
		result1 error
	}
	removeDeadNodesReturnsOnCall map[int]struct {
		result1 error
	}
	SetNodeForRoomStub        func(context.Context, livekit.RoomName, livekit.NodeID) error
	setNodeForRoomMutex       sync.RWMutex
	setNodeForRoomArgsForCall []struct {
		arg1 context.Context
		arg2 livekit.RoomName
		arg3 livekit.NodeID
	}
	setNodeForRoomReturns struct {
		result1 error
	}
	setNodeForRoomReturnsOnCall map[int]struct {
		result1 error
	}
	StartStub        func() error
	startMutex       sync.RWMutex
	startArgsForCall []struct {
	}
	startReturns struct {
		result1 error
	}
	startReturnsOnCall map[int]struct {
		result1 error
	}
	StartParticipantSignalStub        func(context.Context, livekit.RoomName, routing.ParticipantInit) (livekit.ConnectionID, routing.MessageSink, routing.MessageSource, error)
	startParticipantSignalMutex       sync.RWMutex
	startParticipantSignalArgsForCall []struct {
		arg1 context.Context
		arg2 livekit.RoomName
		arg3 routing.ParticipantInit
	}
	startParticipantSignalReturns struct {
		result1 livekit.ConnectionID
		result2 routing.MessageSink
		result3 routing.MessageSource
		result4 error
	}
	startParticipantSignalReturnsOnCall map[int]struct {
		result1 livekit.ConnectionID
		result2 routing.MessageSink
		result3 routing.MessageSource
		result4 error
	}
	StopStub        func()
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
	}
	UnregisterNodeStub        func() error
	unregisterNodeMutex       sync.RWMutex
	unregisterNodeArgsForCall []struct {
	}
	unregisterNodeReturns struct {
		result1 error
	}
	unregisterNodeReturnsOnCall map[int]struct {
		result1 error
	}
	WriteParticipantRTCStub        func(context.Context, livekit.RoomName, livekit.ParticipantIdentity, *livekit.RTCNodeMessage) error
	writeParticipantRTCMutex       sync.RWMutex
	writeParticipantRTCArgsForCall []struct {
		arg1 context.Context
		arg2 livekit.RoomName
		arg3 livekit.ParticipantIdentity
		arg4 *livekit.RTCNodeMessage
	}
	writeParticipantRTCReturns struct {
		result1 error
	}
	writeParticipantRTCReturnsOnCall map[int]struct {
		result1 error
	}
	WriteRoomRTCStub        func(context.Context, livekit.RoomName, *livekit.RTCNodeMessage) error
	writeRoomRTCMutex       sync.RWMutex
	writeRoomRTCArgsForCall []struct {
		arg1 context.Context
		arg2 livekit.RoomName
		arg3 *livekit.RTCNodeMessage
	}
	writeRoomRTCReturns struct {
		result1 error
	}
	writeRoomRTCReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRouter) ClearRoomState(arg1 context.Context, arg2 livekit.RoomName) error {
	fake.clearRoomStateMutex.Lock()
	ret, specificReturn := fake.clearRoomStateReturnsOnCall[len(fake.clearRoomStateArgsForCall)]
	fake.clearRoomStateArgsForCall = append(fake.clearRoomStateArgsForCall, struct {
		arg1 context.Context
		arg2 livekit.RoomName
	}{arg1, arg2})
	stub := fake.ClearRoomStateStub
	fakeReturns := fake.clearRoomStateReturns
	fake.recordInvocation("ClearRoomState", []interface{}{arg1, arg2})
	fake.clearRoomStateMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRouter) ClearRoomStateCallCount() int {
	fake.clearRoomStateMutex.RLock()
	defer fake.clearRoomStateMutex.RUnlock()
	return len(fake.clearRoomStateArgsForCall)
}

func (fake *FakeRouter) ClearRoomStateCalls(stub func(context.Context, livekit.RoomName) error) {
	fake.clearRoomStateMutex.Lock()
	defer fake.clearRoomStateMutex.Unlock()
	fake.ClearRoomStateStub = stub
}

func (fake *FakeRouter) ClearRoomStateArgsForCall(i int) (context.Context, livekit.RoomName) {
	fake.clearRoomStateMutex.RLock()
	defer fake.clearRoomStateMutex.RUnlock()
	argsForCall := fake.clearRoomStateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRouter) ClearRoomStateReturns(result1 error) {
	fake.clearRoomStateMutex.Lock()
	defer fake.clearRoomStateMutex.Unlock()
	fake.ClearRoomStateStub = nil
	fake.clearRoomStateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) ClearRoomStateReturnsOnCall(i int, result1 error) {
	fake.clearRoomStateMutex.Lock()
	defer fake.clearRoomStateMutex.Unlock()
	fake.ClearRoomStateStub = nil
	if fake.clearRoomStateReturnsOnCall == nil {
		fake.clearRoomStateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.clearRoomStateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) Drain() {
	fake.drainMutex.Lock()
	fake.drainArgsForCall = append(fake.drainArgsForCall, struct {
	}{})
	stub := fake.DrainStub
	fake.recordInvocation("Drain", []interface{}{})
	fake.drainMutex.Unlock()
	if stub != nil {
		fake.DrainStub()
	}
}

func (fake *FakeRouter) DrainCallCount() int {
	fake.drainMutex.RLock()
	defer fake.drainMutex.RUnlock()
	return len(fake.drainArgsForCall)
}

func (fake *FakeRouter) DrainCalls(stub func()) {
	fake.drainMutex.Lock()
	defer fake.drainMutex.Unlock()
	fake.DrainStub = stub
}

func (fake *FakeRouter) GetNodeForRoom(arg1 context.Context, arg2 livekit.RoomName) (*livekit.Node, error) {
	fake.getNodeForRoomMutex.Lock()
	ret, specificReturn := fake.getNodeForRoomReturnsOnCall[len(fake.getNodeForRoomArgsForCall)]
	fake.getNodeForRoomArgsForCall = append(fake.getNodeForRoomArgsForCall, struct {
		arg1 context.Context
		arg2 livekit.RoomName
	}{arg1, arg2})
	stub := fake.GetNodeForRoomStub
	fakeReturns := fake.getNodeForRoomReturns
	fake.recordInvocation("GetNodeForRoom", []interface{}{arg1, arg2})
	fake.getNodeForRoomMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRouter) GetNodeForRoomCallCount() int {
	fake.getNodeForRoomMutex.RLock()
	defer fake.getNodeForRoomMutex.RUnlock()
	return len(fake.getNodeForRoomArgsForCall)
}

func (fake *FakeRouter) GetNodeForRoomCalls(stub func(context.Context, livekit.RoomName) (*livekit.Node, error)) {
	fake.getNodeForRoomMutex.Lock()
	defer fake.getNodeForRoomMutex.Unlock()
	fake.GetNodeForRoomStub = stub
}

func (fake *FakeRouter) GetNodeForRoomArgsForCall(i int) (context.Context, livekit.RoomName) {
	fake.getNodeForRoomMutex.RLock()
	defer fake.getNodeForRoomMutex.RUnlock()
	argsForCall := fake.getNodeForRoomArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRouter) GetNodeForRoomReturns(result1 *livekit.Node, result2 error) {
	fake.getNodeForRoomMutex.Lock()
	defer fake.getNodeForRoomMutex.Unlock()
	fake.GetNodeForRoomStub = nil
	fake.getNodeForRoomReturns = struct {
		result1 *livekit.Node
		result2 error
	}{result1, result2}
}

func (fake *FakeRouter) GetNodeForRoomReturnsOnCall(i int, result1 *livekit.Node, result2 error) {
	fake.getNodeForRoomMutex.Lock()
	defer fake.getNodeForRoomMutex.Unlock()
	fake.GetNodeForRoomStub = nil
	if fake.getNodeForRoomReturnsOnCall == nil {
		fake.getNodeForRoomReturnsOnCall = make(map[int]struct {
			result1 *livekit.Node
			result2 error
		})
	}
	fake.getNodeForRoomReturnsOnCall[i] = struct {
		result1 *livekit.Node
		result2 error
	}{result1, result2}
}

func (fake *FakeRouter) GetRegion() string {
	fake.getRegionMutex.Lock()
	ret, specificReturn := fake.getRegionReturnsOnCall[len(fake.getRegionArgsForCall)]
	fake.getRegionArgsForCall = append(fake.getRegionArgsForCall, struct {
	}{})
	stub := fake.GetRegionStub
	fakeReturns := fake.getRegionReturns
	fake.recordInvocation("GetRegion", []interface{}{})
	fake.getRegionMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRouter) GetRegionCallCount() int {
	fake.getRegionMutex.RLock()
	defer fake.getRegionMutex.RUnlock()
	return len(fake.getRegionArgsForCall)
}

func (fake *FakeRouter) GetRegionCalls(stub func() string) {
	fake.getRegionMutex.Lock()
	defer fake.getRegionMutex.Unlock()
	fake.GetRegionStub = stub
}

func (fake *FakeRouter) GetRegionReturns(result1 string) {
	fake.getRegionMutex.Lock()
	defer fake.getRegionMutex.Unlock()
	fake.GetRegionStub = nil
	fake.getRegionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRouter) GetRegionReturnsOnCall(i int, result1 string) {
	fake.getRegionMutex.Lock()
	defer fake.getRegionMutex.Unlock()
	fake.GetRegionStub = nil
	if fake.getRegionReturnsOnCall == nil {
		fake.getRegionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getRegionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeRouter) ListNodes() ([]*livekit.Node, error) {
	fake.listNodesMutex.Lock()
	ret, specificReturn := fake.listNodesReturnsOnCall[len(fake.listNodesArgsForCall)]
	fake.listNodesArgsForCall = append(fake.listNodesArgsForCall, struct {
	}{})
	stub := fake.ListNodesStub
	fakeReturns := fake.listNodesReturns
	fake.recordInvocation("ListNodes", []interface{}{})
	fake.listNodesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRouter) ListNodesCallCount() int {
	fake.listNodesMutex.RLock()
	defer fake.listNodesMutex.RUnlock()
	return len(fake.listNodesArgsForCall)
}

func (fake *FakeRouter) ListNodesCalls(stub func() ([]*livekit.Node, error)) {
	fake.listNodesMutex.Lock()
	defer fake.listNodesMutex.Unlock()
	fake.ListNodesStub = stub
}

func (fake *FakeRouter) ListNodesReturns(result1 []*livekit.Node, result2 error) {
	fake.listNodesMutex.Lock()
	defer fake.listNodesMutex.Unlock()
	fake.ListNodesStub = nil
	fake.listNodesReturns = struct {
		result1 []*livekit.Node
		result2 error
	}{result1, result2}
}

func (fake *FakeRouter) ListNodesReturnsOnCall(i int, result1 []*livekit.Node, result2 error) {
	fake.listNodesMutex.Lock()
	defer fake.listNodesMutex.Unlock()
	fake.ListNodesStub = nil
	if fake.listNodesReturnsOnCall == nil {
		fake.listNodesReturnsOnCall = make(map[int]struct {
			result1 []*livekit.Node
			result2 error
		})
	}
	fake.listNodesReturnsOnCall[i] = struct {
		result1 []*livekit.Node
		result2 error
	}{result1, result2}
}

func (fake *FakeRouter) OnNewParticipantRTC(arg1 routing.NewParticipantCallback) {
	fake.onNewParticipantRTCMutex.Lock()
	fake.onNewParticipantRTCArgsForCall = append(fake.onNewParticipantRTCArgsForCall, struct {
		arg1 routing.NewParticipantCallback
	}{arg1})
	stub := fake.OnNewParticipantRTCStub
	fake.recordInvocation("OnNewParticipantRTC", []interface{}{arg1})
	fake.onNewParticipantRTCMutex.Unlock()
	if stub != nil {
		fake.OnNewParticipantRTCStub(arg1)
	}
}

func (fake *FakeRouter) OnNewParticipantRTCCallCount() int {
	fake.onNewParticipantRTCMutex.RLock()
	defer fake.onNewParticipantRTCMutex.RUnlock()
	return len(fake.onNewParticipantRTCArgsForCall)
}

func (fake *FakeRouter) OnNewParticipantRTCCalls(stub func(routing.NewParticipantCallback)) {
	fake.onNewParticipantRTCMutex.Lock()
	defer fake.onNewParticipantRTCMutex.Unlock()
	fake.OnNewParticipantRTCStub = stub
}

func (fake *FakeRouter) OnNewParticipantRTCArgsForCall(i int) routing.NewParticipantCallback {
	fake.onNewParticipantRTCMutex.RLock()
	defer fake.onNewParticipantRTCMutex.RUnlock()
	argsForCall := fake.onNewParticipantRTCArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRouter) OnRTCMessage(arg1 routing.RTCMessageCallback) {
	fake.onRTCMessageMutex.Lock()
	fake.onRTCMessageArgsForCall = append(fake.onRTCMessageArgsForCall, struct {
		arg1 routing.RTCMessageCallback
	}{arg1})
	stub := fake.OnRTCMessageStub
	fake.recordInvocation("OnRTCMessage", []interface{}{arg1})
	fake.onRTCMessageMutex.Unlock()
	if stub != nil {
		fake.OnRTCMessageStub(arg1)
	}
}

func (fake *FakeRouter) OnRTCMessageCallCount() int {
	fake.onRTCMessageMutex.RLock()
	defer fake.onRTCMessageMutex.RUnlock()
	return len(fake.onRTCMessageArgsForCall)
}

func (fake *FakeRouter) OnRTCMessageCalls(stub func(routing.RTCMessageCallback)) {
	fake.onRTCMessageMutex.Lock()
	defer fake.onRTCMessageMutex.Unlock()
	fake.OnRTCMessageStub = stub
}

func (fake *FakeRouter) OnRTCMessageArgsForCall(i int) routing.RTCMessageCallback {
	fake.onRTCMessageMutex.RLock()
	defer fake.onRTCMessageMutex.RUnlock()
	argsForCall := fake.onRTCMessageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRouter) RegisterNode() error {
	fake.registerNodeMutex.Lock()
	ret, specificReturn := fake.registerNodeReturnsOnCall[len(fake.registerNodeArgsForCall)]
	fake.registerNodeArgsForCall = append(fake.registerNodeArgsForCall, struct {
	}{})
	stub := fake.RegisterNodeStub
	fakeReturns := fake.registerNodeReturns
	fake.recordInvocation("RegisterNode", []interface{}{})
	fake.registerNodeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRouter) RegisterNodeCallCount() int {
	fake.registerNodeMutex.RLock()
	defer fake.registerNodeMutex.RUnlock()
	return len(fake.registerNodeArgsForCall)
}

func (fake *FakeRouter) RegisterNodeCalls(stub func() error) {
	fake.registerNodeMutex.Lock()
	defer fake.registerNodeMutex.Unlock()
	fake.RegisterNodeStub = stub
}

func (fake *FakeRouter) RegisterNodeReturns(result1 error) {
	fake.registerNodeMutex.Lock()
	defer fake.registerNodeMutex.Unlock()
	fake.RegisterNodeStub = nil
	fake.registerNodeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) RegisterNodeReturnsOnCall(i int, result1 error) {
	fake.registerNodeMutex.Lock()
	defer fake.registerNodeMutex.Unlock()
	fake.RegisterNodeStub = nil
	if fake.registerNodeReturnsOnCall == nil {
		fake.registerNodeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.registerNodeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) RemoveDeadNodes() error {
	fake.removeDeadNodesMutex.Lock()
	ret, specificReturn := fake.removeDeadNodesReturnsOnCall[len(fake.removeDeadNodesArgsForCall)]
	fake.removeDeadNodesArgsForCall = append(fake.removeDeadNodesArgsForCall, struct {
	}{})
	stub := fake.RemoveDeadNodesStub
	fakeReturns := fake.removeDeadNodesReturns
	fake.recordInvocation("RemoveDeadNodes", []interface{}{})
	fake.removeDeadNodesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRouter) RemoveDeadNodesCallCount() int {
	fake.removeDeadNodesMutex.RLock()
	defer fake.removeDeadNodesMutex.RUnlock()
	return len(fake.removeDeadNodesArgsForCall)
}

func (fake *FakeRouter) RemoveDeadNodesCalls(stub func() error) {
	fake.removeDeadNodesMutex.Lock()
	defer fake.removeDeadNodesMutex.Unlock()
	fake.RemoveDeadNodesStub = stub
}

func (fake *FakeRouter) RemoveDeadNodesReturns(result1 error) {
	fake.removeDeadNodesMutex.Lock()
	defer fake.removeDeadNodesMutex.Unlock()
	fake.RemoveDeadNodesStub = nil
	fake.removeDeadNodesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) RemoveDeadNodesReturnsOnCall(i int, result1 error) {
	fake.removeDeadNodesMutex.Lock()
	defer fake.removeDeadNodesMutex.Unlock()
	fake.RemoveDeadNodesStub = nil
	if fake.removeDeadNodesReturnsOnCall == nil {
		fake.removeDeadNodesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeDeadNodesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) SetNodeForRoom(arg1 context.Context, arg2 livekit.RoomName, arg3 livekit.NodeID) error {
	fake.setNodeForRoomMutex.Lock()
	ret, specificReturn := fake.setNodeForRoomReturnsOnCall[len(fake.setNodeForRoomArgsForCall)]
	fake.setNodeForRoomArgsForCall = append(fake.setNodeForRoomArgsForCall, struct {
		arg1 context.Context
		arg2 livekit.RoomName
		arg3 livekit.NodeID
	}{arg1, arg2, arg3})
	stub := fake.SetNodeForRoomStub
	fakeReturns := fake.setNodeForRoomReturns
	fake.recordInvocation("SetNodeForRoom", []interface{}{arg1, arg2, arg3})
	fake.setNodeForRoomMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRouter) SetNodeForRoomCallCount() int {
	fake.setNodeForRoomMutex.RLock()
	defer fake.setNodeForRoomMutex.RUnlock()
	return len(fake.setNodeForRoomArgsForCall)
}

func (fake *FakeRouter) SetNodeForRoomCalls(stub func(context.Context, livekit.RoomName, livekit.NodeID) error) {
	fake.setNodeForRoomMutex.Lock()
	defer fake.setNodeForRoomMutex.Unlock()
	fake.SetNodeForRoomStub = stub
}

func (fake *FakeRouter) SetNodeForRoomArgsForCall(i int) (context.Context, livekit.RoomName, livekit.NodeID) {
	fake.setNodeForRoomMutex.RLock()
	defer fake.setNodeForRoomMutex.RUnlock()
	argsForCall := fake.setNodeForRoomArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRouter) SetNodeForRoomReturns(result1 error) {
	fake.setNodeForRoomMutex.Lock()
	defer fake.setNodeForRoomMutex.Unlock()
	fake.SetNodeForRoomStub = nil
	fake.setNodeForRoomReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) SetNodeForRoomReturnsOnCall(i int, result1 error) {
	fake.setNodeForRoomMutex.Lock()
	defer fake.setNodeForRoomMutex.Unlock()
	fake.SetNodeForRoomStub = nil
	if fake.setNodeForRoomReturnsOnCall == nil {
		fake.setNodeForRoomReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setNodeForRoomReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) Start() error {
	fake.startMutex.Lock()
	ret, specificReturn := fake.startReturnsOnCall[len(fake.startArgsForCall)]
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
	}{})
	stub := fake.StartStub
	fakeReturns := fake.startReturns
	fake.recordInvocation("Start", []interface{}{})
	fake.startMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRouter) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeRouter) StartCalls(stub func() error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = stub
}

func (fake *FakeRouter) StartReturns(result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) StartReturnsOnCall(i int, result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	if fake.startReturnsOnCall == nil {
		fake.startReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.startReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) StartParticipantSignal(arg1 context.Context, arg2 livekit.RoomName, arg3 routing.ParticipantInit) (livekit.ConnectionID, routing.MessageSink, routing.MessageSource, error) {
	fake.startParticipantSignalMutex.Lock()
	ret, specificReturn := fake.startParticipantSignalReturnsOnCall[len(fake.startParticipantSignalArgsForCall)]
	fake.startParticipantSignalArgsForCall = append(fake.startParticipantSignalArgsForCall, struct {
		arg1 context.Context
		arg2 livekit.RoomName
		arg3 routing.ParticipantInit
	}{arg1, arg2, arg3})
	stub := fake.StartParticipantSignalStub
	fakeReturns := fake.startParticipantSignalReturns
	fake.recordInvocation("StartParticipantSignal", []interface{}{arg1, arg2, arg3})
	fake.startParticipantSignalMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3, fakeReturns.result4
}

func (fake *FakeRouter) StartParticipantSignalCallCount() int {
	fake.startParticipantSignalMutex.RLock()
	defer fake.startParticipantSignalMutex.RUnlock()
	return len(fake.startParticipantSignalArgsForCall)
}

func (fake *FakeRouter) StartParticipantSignalCalls(stub func(context.Context, livekit.RoomName, routing.ParticipantInit) (livekit.ConnectionID, routing.MessageSink, routing.MessageSource, error)) {
	fake.startParticipantSignalMutex.Lock()
	defer fake.startParticipantSignalMutex.Unlock()
	fake.StartParticipantSignalStub = stub
}

func (fake *FakeRouter) StartParticipantSignalArgsForCall(i int) (context.Context, livekit.RoomName, routing.ParticipantInit) {
	fake.startParticipantSignalMutex.RLock()
	defer fake.startParticipantSignalMutex.RUnlock()
	argsForCall := fake.startParticipantSignalArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRouter) StartParticipantSignalReturns(result1 livekit.ConnectionID, result2 routing.MessageSink, result3 routing.MessageSource, result4 error) {
	fake.startParticipantSignalMutex.Lock()
	defer fake.startParticipantSignalMutex.Unlock()
	fake.StartParticipantSignalStub = nil
	fake.startParticipantSignalReturns = struct {
		result1 livekit.ConnectionID
		result2 routing.MessageSink
		result3 routing.MessageSource
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeRouter) StartParticipantSignalReturnsOnCall(i int, result1 livekit.ConnectionID, result2 routing.MessageSink, result3 routing.MessageSource, result4 error) {
	fake.startParticipantSignalMutex.Lock()
	defer fake.startParticipantSignalMutex.Unlock()
	fake.StartParticipantSignalStub = nil
	if fake.startParticipantSignalReturnsOnCall == nil {
		fake.startParticipantSignalReturnsOnCall = make(map[int]struct {
			result1 livekit.ConnectionID
			result2 routing.MessageSink
			result3 routing.MessageSource
			result4 error
		})
	}
	fake.startParticipantSignalReturnsOnCall[i] = struct {
		result1 livekit.ConnectionID
		result2 routing.MessageSink
		result3 routing.MessageSource
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeRouter) Stop() {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
	}{})
	stub := fake.StopStub
	fake.recordInvocation("Stop", []interface{}{})
	fake.stopMutex.Unlock()
	if stub != nil {
		fake.StopStub()
	}
}

func (fake *FakeRouter) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeRouter) StopCalls(stub func()) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = stub
}

func (fake *FakeRouter) UnregisterNode() error {
	fake.unregisterNodeMutex.Lock()
	ret, specificReturn := fake.unregisterNodeReturnsOnCall[len(fake.unregisterNodeArgsForCall)]
	fake.unregisterNodeArgsForCall = append(fake.unregisterNodeArgsForCall, struct {
	}{})
	stub := fake.UnregisterNodeStub
	fakeReturns := fake.unregisterNodeReturns
	fake.recordInvocation("UnregisterNode", []interface{}{})
	fake.unregisterNodeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRouter) UnregisterNodeCallCount() int {
	fake.unregisterNodeMutex.RLock()
	defer fake.unregisterNodeMutex.RUnlock()
	return len(fake.unregisterNodeArgsForCall)
}

func (fake *FakeRouter) UnregisterNodeCalls(stub func() error) {
	fake.unregisterNodeMutex.Lock()
	defer fake.unregisterNodeMutex.Unlock()
	fake.UnregisterNodeStub = stub
}

func (fake *FakeRouter) UnregisterNodeReturns(result1 error) {
	fake.unregisterNodeMutex.Lock()
	defer fake.unregisterNodeMutex.Unlock()
	fake.UnregisterNodeStub = nil
	fake.unregisterNodeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) UnregisterNodeReturnsOnCall(i int, result1 error) {
	fake.unregisterNodeMutex.Lock()
	defer fake.unregisterNodeMutex.Unlock()
	fake.UnregisterNodeStub = nil
	if fake.unregisterNodeReturnsOnCall == nil {
		fake.unregisterNodeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unregisterNodeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) WriteParticipantRTC(arg1 context.Context, arg2 livekit.RoomName, arg3 livekit.ParticipantIdentity, arg4 *livekit.RTCNodeMessage) error {
	fake.writeParticipantRTCMutex.Lock()
	ret, specificReturn := fake.writeParticipantRTCReturnsOnCall[len(fake.writeParticipantRTCArgsForCall)]
	fake.writeParticipantRTCArgsForCall = append(fake.writeParticipantRTCArgsForCall, struct {
		arg1 context.Context
		arg2 livekit.RoomName
		arg3 livekit.ParticipantIdentity
		arg4 *livekit.RTCNodeMessage
	}{arg1, arg2, arg3, arg4})
	stub := fake.WriteParticipantRTCStub
	fakeReturns := fake.writeParticipantRTCReturns
	fake.recordInvocation("WriteParticipantRTC", []interface{}{arg1, arg2, arg3, arg4})
	fake.writeParticipantRTCMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRouter) WriteParticipantRTCCallCount() int {
	fake.writeParticipantRTCMutex.RLock()
	defer fake.writeParticipantRTCMutex.RUnlock()
	return len(fake.writeParticipantRTCArgsForCall)
}

func (fake *FakeRouter) WriteParticipantRTCCalls(stub func(context.Context, livekit.RoomName, livekit.ParticipantIdentity, *livekit.RTCNodeMessage) error) {
	fake.writeParticipantRTCMutex.Lock()
	defer fake.writeParticipantRTCMutex.Unlock()
	fake.WriteParticipantRTCStub = stub
}

func (fake *FakeRouter) WriteParticipantRTCArgsForCall(i int) (context.Context, livekit.RoomName, livekit.ParticipantIdentity, *livekit.RTCNodeMessage) {
	fake.writeParticipantRTCMutex.RLock()
	defer fake.writeParticipantRTCMutex.RUnlock()
	argsForCall := fake.writeParticipantRTCArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeRouter) WriteParticipantRTCReturns(result1 error) {
	fake.writeParticipantRTCMutex.Lock()
	defer fake.writeParticipantRTCMutex.Unlock()
	fake.WriteParticipantRTCStub = nil
	fake.writeParticipantRTCReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) WriteParticipantRTCReturnsOnCall(i int, result1 error) {
	fake.writeParticipantRTCMutex.Lock()
	defer fake.writeParticipantRTCMutex.Unlock()
	fake.WriteParticipantRTCStub = nil
	if fake.writeParticipantRTCReturnsOnCall == nil {
		fake.writeParticipantRTCReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeParticipantRTCReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) WriteRoomRTC(arg1 context.Context, arg2 livekit.RoomName, arg3 *livekit.RTCNodeMessage) error {
	fake.writeRoomRTCMutex.Lock()
	ret, specificReturn := fake.writeRoomRTCReturnsOnCall[len(fake.writeRoomRTCArgsForCall)]
	fake.writeRoomRTCArgsForCall = append(fake.writeRoomRTCArgsForCall, struct {
		arg1 context.Context
		arg2 livekit.RoomName
		arg3 *livekit.RTCNodeMessage
	}{arg1, arg2, arg3})
	stub := fake.WriteRoomRTCStub
	fakeReturns := fake.writeRoomRTCReturns
	fake.recordInvocation("WriteRoomRTC", []interface{}{arg1, arg2, arg3})
	fake.writeRoomRTCMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRouter) WriteRoomRTCCallCount() int {
	fake.writeRoomRTCMutex.RLock()
	defer fake.writeRoomRTCMutex.RUnlock()
	return len(fake.writeRoomRTCArgsForCall)
}

func (fake *FakeRouter) WriteRoomRTCCalls(stub func(context.Context, livekit.RoomName, *livekit.RTCNodeMessage) error) {
	fake.writeRoomRTCMutex.Lock()
	defer fake.writeRoomRTCMutex.Unlock()
	fake.WriteRoomRTCStub = stub
}

func (fake *FakeRouter) WriteRoomRTCArgsForCall(i int) (context.Context, livekit.RoomName, *livekit.RTCNodeMessage) {
	fake.writeRoomRTCMutex.RLock()
	defer fake.writeRoomRTCMutex.RUnlock()
	argsForCall := fake.writeRoomRTCArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRouter) WriteRoomRTCReturns(result1 error) {
	fake.writeRoomRTCMutex.Lock()
	defer fake.writeRoomRTCMutex.Unlock()
	fake.WriteRoomRTCStub = nil
	fake.writeRoomRTCReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) WriteRoomRTCReturnsOnCall(i int, result1 error) {
	fake.writeRoomRTCMutex.Lock()
	defer fake.writeRoomRTCMutex.Unlock()
	fake.WriteRoomRTCStub = nil
	if fake.writeRoomRTCReturnsOnCall == nil {
		fake.writeRoomRTCReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeRoomRTCReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.clearRoomStateMutex.RLock()
	defer fake.clearRoomStateMutex.RUnlock()
	fake.drainMutex.RLock()
	defer fake.drainMutex.RUnlock()
	fake.getNodeForRoomMutex.RLock()
	defer fake.getNodeForRoomMutex.RUnlock()
	fake.getRegionMutex.RLock()
	defer fake.getRegionMutex.RUnlock()
	fake.listNodesMutex.RLock()
	defer fake.listNodesMutex.RUnlock()
	fake.onNewParticipantRTCMutex.RLock()
	defer fake.onNewParticipantRTCMutex.RUnlock()
	fake.onRTCMessageMutex.RLock()
	defer fake.onRTCMessageMutex.RUnlock()
	fake.registerNodeMutex.RLock()
	defer fake.registerNodeMutex.RUnlock()
	fake.removeDeadNodesMutex.RLock()
	defer fake.removeDeadNodesMutex.RUnlock()
	fake.setNodeForRoomMutex.RLock()
	defer fake.setNodeForRoomMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.startParticipantSignalMutex.RLock()
	defer fake.startParticipantSignalMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	fake.unregisterNodeMutex.RLock()
	defer fake.unregisterNodeMutex.RUnlock()
	fake.writeParticipantRTCMutex.RLock()
	defer fake.writeParticipantRTCMutex.RUnlock()
	fake.writeRoomRTCMutex.RLock()
	defer fake.writeRoomRTCMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRouter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ routing.Router = new(FakeRouter)
