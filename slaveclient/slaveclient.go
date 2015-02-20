package slaveclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/Bo0mer/enslaver/model"

	osaModel "github.com/Bo0mer/os-agent/model"
)

type SlaveClient interface {
	Execute(osaModel.CommandRequest, Slave) (string, error)
	ExecuteAsync(osaModel.CommandRequest, Slave) (string, error)
	Job(string, Slave) (osaModel.Job, error)
}

type slaveClient struct {
	httpClient *http.Client
}

// Return new Slave client
func NewSlaveClient() SlaveClient {
	return &slaveClient{
		httpClient: new(http.Client),
	}
}

// Force slave to execute the specified command
func (c *slaveClient) Execute(command osaModel.CommandRequest, slave Slave) (string, error) {
	jobRequest := osaModel.JobRequest{
		Async:   false,
		Command: command,
	}
	return c.execute(jobRequest, c.buildSlaveUrl(slave))
}

// Force slave to execute the specified command async
func (c *slaveClient) ExecuteAsync(command osaModel.CommandRequest, slave Slave) (string, error) {
	jobRequest := osaModel.JobRequest{
		Async:   true,
		Command: command,
	}
	return c.execute(jobRequest, c.buildSlaveUrl(slave))
}

// Get the job with id jobId from the specified slave
func (c *slaveClient) Job(jobId string, slave Slave) (osaModel.Job, error) {
	getJobUrl := fmt.Sprintf("%s?id=%s", c.buildSlaveUrl(slave), jobId)
	response, err := c.doGet(getJobUrl)
	if err != nil {
		return osaModel.Job{}, err
	}

	job := &osaModel.Job{}
	err = json.Unmarshal(response, job)
	if err != nil {
		return osaModel.Job{}, err
	}

	return *job, nil
}

func (c *slaveClient) execute(jobRequest osaModel.JobRequest, slaveUrl string) (string, error) {
	data, _ := json.Marshal(jobRequest)
	response, err := c.doPost(slaveUrl, data)
	if err != nil {
		return "", err
	}

	job := &osaModel.Job{}
	err = json.Unmarshal(response, job)
	if err != nil {
		return "", err
	}

	return job.Id, nil
}

func (c *slaveClient) buildSlaveUrl(slave Slave) string {
	return fmt.Sprintf("http://%s:%d/jobs", slave.Host, slave.Port)
}

func (c *slaveClient) doGet(url string) ([]byte, error) {
	request, _ := http.NewRequest("GET", url, nil)
	return c.doRequest(request)
}

func (c *slaveClient) doPost(url string, data []byte) ([]byte, error) {
	body := bytes.NewReader(data)

	request, _ := http.NewRequest("POST", url, body)
	request.Header.Add("Content-Type", "application/json")
	return c.doRequest(request)
}

func (c *slaveClient) doRequest(request *http.Request) ([]byte, error) {
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if 200 > response.StatusCode || response.StatusCode >= 300 {
		return nil, errors.New(fmt.Sprintf("Response status code: %d", response.StatusCode))
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
