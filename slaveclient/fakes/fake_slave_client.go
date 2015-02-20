// This file was generated by counterfeiter
package fakes

import (
	"sync"

	. "github.com/Bo0mer/enslaver/model"
	"github.com/Bo0mer/enslaver/slaveclient"
	osaModel "github.com/Bo0mer/os-agent/model"
)

type FakeSlaveClient struct {
	ExecuteStub        func(osaModel.CommandRequest, Slave) (string, error)
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		arg1 osaModel.CommandRequest
		arg2 Slave
	}
	executeReturns struct {
		result1 string
		result2 error
	}
	ExecuteAsyncStub        func(osaModel.CommandRequest, Slave) (string, error)
	executeAsyncMutex       sync.RWMutex
	executeAsyncArgsForCall []struct {
		arg1 osaModel.CommandRequest
		arg2 Slave
	}
	executeAsyncReturns struct {
		result1 string
		result2 error
	}
	JobStub        func(string, Slave) (osaModel.Job, error)
	jobMutex       sync.RWMutex
	jobArgsForCall []struct {
		arg1 string
		arg2 Slave
	}
	jobReturns struct {
		result1 osaModel.Job
		result2 error
	}
}

func (fake *FakeSlaveClient) Execute(arg1 osaModel.CommandRequest, arg2 Slave) (string, error) {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		arg1 osaModel.CommandRequest
		arg2 Slave
	}{arg1, arg2})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub(arg1, arg2)
	} else {
		return fake.executeReturns.result1, fake.executeReturns.result2
	}
}

func (fake *FakeSlaveClient) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeSlaveClient) ExecuteArgsForCall(i int) (osaModel.CommandRequest, Slave) {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.executeArgsForCall[i].arg1, fake.executeArgsForCall[i].arg2
}

func (fake *FakeSlaveClient) ExecuteReturns(result1 string, result2 error) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeSlaveClient) ExecuteAsync(arg1 osaModel.CommandRequest, arg2 Slave) (string, error) {
	fake.executeAsyncMutex.Lock()
	fake.executeAsyncArgsForCall = append(fake.executeAsyncArgsForCall, struct {
		arg1 osaModel.CommandRequest
		arg2 Slave
	}{arg1, arg2})
	fake.executeAsyncMutex.Unlock()
	if fake.ExecuteAsyncStub != nil {
		return fake.ExecuteAsyncStub(arg1, arg2)
	} else {
		return fake.executeAsyncReturns.result1, fake.executeAsyncReturns.result2
	}
}

func (fake *FakeSlaveClient) ExecuteAsyncCallCount() int {
	fake.executeAsyncMutex.RLock()
	defer fake.executeAsyncMutex.RUnlock()
	return len(fake.executeAsyncArgsForCall)
}

func (fake *FakeSlaveClient) ExecuteAsyncArgsForCall(i int) (osaModel.CommandRequest, Slave) {
	fake.executeAsyncMutex.RLock()
	defer fake.executeAsyncMutex.RUnlock()
	return fake.executeAsyncArgsForCall[i].arg1, fake.executeAsyncArgsForCall[i].arg2
}

func (fake *FakeSlaveClient) ExecuteAsyncReturns(result1 string, result2 error) {
	fake.ExecuteAsyncStub = nil
	fake.executeAsyncReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeSlaveClient) Job(arg1 string, arg2 Slave) (osaModel.Job, error) {
	fake.jobMutex.Lock()
	fake.jobArgsForCall = append(fake.jobArgsForCall, struct {
		arg1 string
		arg2 Slave
	}{arg1, arg2})
	fake.jobMutex.Unlock()
	if fake.JobStub != nil {
		return fake.JobStub(arg1, arg2)
	} else {
		return fake.jobReturns.result1, fake.jobReturns.result2
	}
}

func (fake *FakeSlaveClient) JobCallCount() int {
	fake.jobMutex.RLock()
	defer fake.jobMutex.RUnlock()
	return len(fake.jobArgsForCall)
}

func (fake *FakeSlaveClient) JobArgsForCall(i int) (string, Slave) {
	fake.jobMutex.RLock()
	defer fake.jobMutex.RUnlock()
	return fake.jobArgsForCall[i].arg1, fake.jobArgsForCall[i].arg2
}

func (fake *FakeSlaveClient) JobReturns(result1 osaModel.Job, result2 error) {
	fake.JobStub = nil
	fake.jobReturns = struct {
		result1 osaModel.Job
		result2 error
	}{result1, result2}
}

var _ slaveclient.SlaveClient = new(FakeSlaveClient)
