package facade

import (
	"encoding/json"
	"net/http"

	. "github.com/Bo0mer/enslaver/enslaver"
	"github.com/Bo0mer/enslaver/model"

	"github.com/Bo0mer/os-agent/server"
)

type EnslaverFacade interface {
	RegisterSlave(server.Request, server.Response)
	CreateJob(server.Request, server.Response)
	Slaves(server.Request, server.Response)
}

type enslaverFacade struct {
	enslaver Enslaver
}

func NewEnslaverFacade(enslaver Enslaver) EnslaverFacade {
	return &enslaverFacade{
		enslaver: enslaver,
	}
}

func (f *enslaverFacade) RegisterSlave(req server.Request, resp server.Response) {
	slave := &model.Slave{}
	err := json.Unmarshal(req.Body(), slave)
	if err != nil {
		resp.SetStatusCode(http.StatusBadRequest)
		return
	}

	f.enslaver.Register(*slave)
	resp.SetStatusCode(http.StatusOK)
}

func (f *enslaverFacade) CreateJob(req server.Request, resp server.Response) {
	jobRequest := &model.JobRequest{}
	err := json.Unmarshal(req.Body(), jobRequest)
	if err != nil {
		resp.SetStatusCode(http.StatusBadRequest)
		return
	}

	job, err := f.enslaver.Execute(*jobRequest)
	if err != nil {
		resp.SetStatusCode(http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(job)
	resp.SetBody(responseBody)
	resp.SetStatusCode(http.StatusOK)
}

func (f *enslaverFacade) Slaves(req server.Request, resp server.Response) {
	slaves := f.enslaver.Slaves()
	slavesJSON, _ := json.Marshal(slaves)

	resp.SetBody(slavesJSON)
	resp.SetStatusCode(http.StatusOK)
}
