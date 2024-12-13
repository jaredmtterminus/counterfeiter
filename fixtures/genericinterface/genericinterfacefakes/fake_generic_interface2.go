// Code generated by counterfeiter. DO NOT EDIT.
package genericinterfacefakes

import (
	"sync"

	"github.com/maxbrunsfeld/counterfeiter/v6/fixtures/genericinterface"
)

type FakeGenericInterface2[T genericinterface.CustomTypeT] struct {
	DoSomethingStub        func()
	doSomethingMutex       sync.RWMutex
	doSomethingArgsForCall []struct {
	}
	ReturnTStub        func() T
	returnTMutex       sync.RWMutex
	returnTArgsForCall []struct {
	}
	returnTReturns struct {
		result1 T
	}
	returnTReturnsOnCall map[int]struct {
		result1 T
	}
	TakeAndReturnTStub        func(T) T
	takeAndReturnTMutex       sync.RWMutex
	takeAndReturnTArgsForCall []struct {
		arg1 T
	}
	takeAndReturnTReturns struct {
		result1 T
	}
	takeAndReturnTReturnsOnCall map[int]struct {
		result1 T
	}
	TakeTStub        func(T)
	takeTMutex       sync.RWMutex
	takeTArgsForCall []struct {
		arg1 T
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGenericInterface2[T]) DoSomething() {
	fake.doSomethingMutex.Lock()
	fake.doSomethingArgsForCall = append(fake.doSomethingArgsForCall, struct {
	}{})
	stub := fake.DoSomethingStub
	fake.recordInvocation("DoSomething", []interface{}{})
	fake.doSomethingMutex.Unlock()
	if stub != nil {
		fake.DoSomethingStub()
	}
}

func (fake *FakeGenericInterface2[T]) DoSomethingCallCount() int {
	fake.doSomethingMutex.RLock()
	defer fake.doSomethingMutex.RUnlock()
	return len(fake.doSomethingArgsForCall)
}

func (fake *FakeGenericInterface2[T]) DoSomethingCalls(stub func()) {
	fake.doSomethingMutex.Lock()
	defer fake.doSomethingMutex.Unlock()
	fake.DoSomethingStub = stub
}

func (fake *FakeGenericInterface2[T]) ReturnT() T {
	fake.returnTMutex.Lock()
	ret, specificReturn := fake.returnTReturnsOnCall[len(fake.returnTArgsForCall)]
	fake.returnTArgsForCall = append(fake.returnTArgsForCall, struct {
	}{})
	stub := fake.ReturnTStub
	fakeReturns := fake.returnTReturns
	fake.recordInvocation("ReturnT", []interface{}{})
	fake.returnTMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeGenericInterface2[T]) ReturnTCallCount() int {
	fake.returnTMutex.RLock()
	defer fake.returnTMutex.RUnlock()
	return len(fake.returnTArgsForCall)
}

func (fake *FakeGenericInterface2[T]) ReturnTCalls(stub func() T) {
	fake.returnTMutex.Lock()
	defer fake.returnTMutex.Unlock()
	fake.ReturnTStub = stub
}

func (fake *FakeGenericInterface2[T]) ReturnTReturns(result1 T) {
	fake.returnTMutex.Lock()
	defer fake.returnTMutex.Unlock()
	fake.ReturnTStub = nil
	fake.returnTReturns = struct {
		result1 T
	}{result1}
}

func (fake *FakeGenericInterface2[T]) ReturnTReturnsOnCall(i int, result1 T) {
	fake.returnTMutex.Lock()
	defer fake.returnTMutex.Unlock()
	fake.ReturnTStub = nil
	if fake.returnTReturnsOnCall == nil {
		fake.returnTReturnsOnCall = make(map[int]struct {
			result1 T
		})
	}
	fake.returnTReturnsOnCall[i] = struct {
		result1 T
	}{result1}
}

func (fake *FakeGenericInterface2[T]) TakeAndReturnT(arg1 T) T {
	fake.takeAndReturnTMutex.Lock()
	ret, specificReturn := fake.takeAndReturnTReturnsOnCall[len(fake.takeAndReturnTArgsForCall)]
	fake.takeAndReturnTArgsForCall = append(fake.takeAndReturnTArgsForCall, struct {
		arg1 T
	}{arg1})
	stub := fake.TakeAndReturnTStub
	fakeReturns := fake.takeAndReturnTReturns
	fake.recordInvocation("TakeAndReturnT", []interface{}{arg1})
	fake.takeAndReturnTMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeGenericInterface2[T]) TakeAndReturnTCallCount() int {
	fake.takeAndReturnTMutex.RLock()
	defer fake.takeAndReturnTMutex.RUnlock()
	return len(fake.takeAndReturnTArgsForCall)
}

func (fake *FakeGenericInterface2[T]) TakeAndReturnTCalls(stub func(T) T) {
	fake.takeAndReturnTMutex.Lock()
	defer fake.takeAndReturnTMutex.Unlock()
	fake.TakeAndReturnTStub = stub
}

func (fake *FakeGenericInterface2[T]) TakeAndReturnTArgsForCall(i int) T {
	fake.takeAndReturnTMutex.RLock()
	defer fake.takeAndReturnTMutex.RUnlock()
	argsForCall := fake.takeAndReturnTArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGenericInterface2[T]) TakeAndReturnTReturns(result1 T) {
	fake.takeAndReturnTMutex.Lock()
	defer fake.takeAndReturnTMutex.Unlock()
	fake.TakeAndReturnTStub = nil
	fake.takeAndReturnTReturns = struct {
		result1 T
	}{result1}
}

func (fake *FakeGenericInterface2[T]) TakeAndReturnTReturnsOnCall(i int, result1 T) {
	fake.takeAndReturnTMutex.Lock()
	defer fake.takeAndReturnTMutex.Unlock()
	fake.TakeAndReturnTStub = nil
	if fake.takeAndReturnTReturnsOnCall == nil {
		fake.takeAndReturnTReturnsOnCall = make(map[int]struct {
			result1 T
		})
	}
	fake.takeAndReturnTReturnsOnCall[i] = struct {
		result1 T
	}{result1}
}

func (fake *FakeGenericInterface2[T]) TakeT(arg1 T) {
	fake.takeTMutex.Lock()
	fake.takeTArgsForCall = append(fake.takeTArgsForCall, struct {
		arg1 T
	}{arg1})
	stub := fake.TakeTStub
	fake.recordInvocation("TakeT", []interface{}{arg1})
	fake.takeTMutex.Unlock()
	if stub != nil {
		fake.TakeTStub(arg1)
	}
}

func (fake *FakeGenericInterface2[T]) TakeTCallCount() int {
	fake.takeTMutex.RLock()
	defer fake.takeTMutex.RUnlock()
	return len(fake.takeTArgsForCall)
}

func (fake *FakeGenericInterface2[T]) TakeTCalls(stub func(T)) {
	fake.takeTMutex.Lock()
	defer fake.takeTMutex.Unlock()
	fake.TakeTStub = stub
}

func (fake *FakeGenericInterface2[T]) TakeTArgsForCall(i int) T {
	fake.takeTMutex.RLock()
	defer fake.takeTMutex.RUnlock()
	argsForCall := fake.takeTArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGenericInterface2[T]) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.doSomethingMutex.RLock()
	defer fake.doSomethingMutex.RUnlock()
	fake.returnTMutex.RLock()
	defer fake.returnTMutex.RUnlock()
	fake.takeAndReturnTMutex.RLock()
	defer fake.takeAndReturnTMutex.RUnlock()
	fake.takeTMutex.RLock()
	defer fake.takeTMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGenericInterface2[T]) recordInvocation(key string, args []interface{}) {
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

var _ genericinterface.GenericInterface2[genericinterface.CustomTypeT] = new(FakeGenericInterface2[genericinterface.CustomTypeT])