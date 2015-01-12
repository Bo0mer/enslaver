package model

type Slave struct {
	Id   string            `json:"id"`
	Host string            `json:"host"`
	Port int               `json:"port"`
	Tags map[string]string `json:"tags"`
}
