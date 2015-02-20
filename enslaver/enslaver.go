package enslaver

import (
	"errors"
	"fmt"
	"sync"
	"time"

	. "github.com/Bo0mer/enslaver/model"
	"github.com/Bo0mer/enslaver/slaveclient"
)

type Enslaver interface {
	Register(Slave)
	Slaves() []Slave
	Execute(JobRequest) (Job, error)
}

type enslaver struct {
	slavesMtx   sync.RWMutex
	slaves      map[string]Slave
	slaveClient slaveclient.SlaveClient
	jobs        map[string]Job
}

func NewEnslaver(slaveClient slaveclient.SlaveClient) Enslaver {
	return &enslaver{
		slaves:      make(map[string]Slave),
		slaveClient: slaveClient,
	}
}

func (e *enslaver) Register(slave Slave) {
	e.slavesMtx.Lock()
	e.slaves[slave.Id] = slave
	e.slavesMtx.Unlock()
}

func (e *enslaver) Slaves() []Slave {
	slavesSlice := make([]Slave, 0, len(e.slaves))

	for _, slave := range e.slaves {
		slavesSlice = append(slavesSlice, slave)
	}

	return slavesSlice
}

func (e *enslaver) Execute(jobRequest JobRequest) (Job, error) {
	if len(jobRequest.Slaves) == 0 {
		return Job{}, errors.New(fmt.Sprintf("No slaves specified."))
	}
	for _, slave := range jobRequest.Slaves {
		_, found := e.slaves[slave.Id]
		if !found {
			return Job{}, errors.New(fmt.Sprintf("Slave with id %s is missing.", slave.Id))
		}
	}

	job := Job{
		Id:      fmt.Sprintf("%d", time.Now().UnixNano()),
		Status:  JOB_IN_PROCESS,
		Results: make([]JobResult, len(jobRequest.Slaves)),
	}

	for i, slave := range jobRequest.Slaves {
		// get the host and the port
		slave = e.slaves[slave.Id]
		slaveJobId, _ := e.slaveClient.Execute(jobRequest.Command, slave)
		slaveJob, _ := e.slaveClient.Job(slaveJobId, slave)

		job.Results[i] = JobResult{
			Slave:  slave,
			Status: JobStatus(slaveJob.Status),
			Result: slaveJob.Result,
		}
	}

	job.Status = JOB_COMPLETED

	return job, nil
}
