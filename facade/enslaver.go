package facade

import (
	"encoding/json"
	"net/http"

	. "github.com/Bo0mer/enslaver/enslaver"
	"github.com/Bo0mer/enslaver/model"

	"github.com/Bo0mer/os-agent/server"

	l4g "code.google.com/p/log4go"
)

type EnslaverFacade interface {
	RegisterSlave(server.Request, server.Response)
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
		l4g.Error("Unable to parse register slave request", err)
		resp.SetStatusCode(http.StatusBadRequest)
		return
	}

	f.enslaver.Register(*slave)
	l4g.Debug("Registered slave: %v", *slave)
	resp.SetStatusCode(http.StatusOK)
}
