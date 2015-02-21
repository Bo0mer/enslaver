package model

type Slave struct {
	Id   string            `json:"id"`
	Host string            `json:"host,omitempty"`
	Port int               `json:"port,omitempty"`
	Tags map[string]string `json:"tags"`
}
