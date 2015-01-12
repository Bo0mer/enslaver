package enslaver

import (
	"sync"

	. "github.com/Bo0mer/enslaver/model"
)

type Enslaver interface {
	Register(Slave)
	Slaves() []Slave
}

type enslaver struct {
	slavesMtx sync.RWMutex
	slaves    map[string]Slave
}

func NewEnslaver() Enslaver {
	return &enslaver{
		slaves: make(map[string]Slave),
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
