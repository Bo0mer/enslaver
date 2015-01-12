// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/Bo0mer/enslaver/enslaver"
	. "github.com/Bo0mer/enslaver/model"
)

type FakeEnslaver struct {
	RegisterStub        func(Slave)
	registerMutex       sync.RWMutex
	registerArgsForCall []struct {
		arg1 Slave
	}
	SlavesStub        func() []Slave
	slavesMutex       sync.RWMutex
	slavesArgsForCall []struct{}
	slavesReturns struct {
		result1 []Slave
	}
}

func (fake *FakeEnslaver) Register(arg1 Slave) {
	fake.registerMutex.Lock()
	fake.registerArgsForCall = append(fake.registerArgsForCall, struct {
		arg1 Slave
	}{arg1})
	fake.registerMutex.Unlock()
	if fake.RegisterStub != nil {
		fake.RegisterStub(arg1)
	}
}

func (fake *FakeEnslaver) RegisterCallCount() int {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	return len(fake.registerArgsForCall)
}

func (fake *FakeEnslaver) RegisterArgsForCall(i int) Slave {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	return fake.registerArgsForCall[i].arg1
}

func (fake *FakeEnslaver) Slaves() []Slave {
	fake.slavesMutex.Lock()
	fake.slavesArgsForCall = append(fake.slavesArgsForCall, struct{}{})
	fake.slavesMutex.Unlock()
	if fake.SlavesStub != nil {
		return fake.SlavesStub()
	} else {
		return fake.slavesReturns.result1
	}
}

func (fake *FakeEnslaver) SlavesCallCount() int {
	fake.slavesMutex.RLock()
	defer fake.slavesMutex.RUnlock()
	return len(fake.slavesArgsForCall)
}

func (fake *FakeEnslaver) SlavesReturns(result1 []Slave) {
	fake.SlavesStub = nil
	fake.slavesReturns = struct {
		result1 []Slave
	}{result1}
}

var _ enslaver.Enslaver = new(FakeEnslaver)