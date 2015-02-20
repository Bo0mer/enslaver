package model

import (
	osaModel "github.com/Bo0mer/os-agent/model"
)

type JobStatus string

const (
	JOB_IN_PROCESS JobStatus = "IN_PROCESS"
	JOB_COMPLETED  JobStatus = "COMPLETED"
)

type JobRequest struct {
	Slaves  []Slave                 `json:"slaves"`
	Async   bool                    `json:"async"`
	Command osaModel.CommandRequest `json:"command"`
}

type JobResult struct {
	Slave  Slave                    `json:"slave"`
	Status JobStatus                `json:"status"`
	Result osaModel.CommandResponse `json:"result"`
}

type Job struct {
	Id      string      `json:"id"`
	Status  JobStatus   `json:"status"`
	Results []JobResult `json:"results"`
}
