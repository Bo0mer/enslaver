package slaveclient_test

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Bo0mer/enslaver/model"
	. "github.com/Bo0mer/enslaver/slaveclient"

	osaModel "github.com/Bo0mer/os-agent/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Slaveclient", func() {

	var fakeServer *ghttp.Server
	var slaveClient SlaveClient
	var serverHost string
	var serverPort int

	BeforeEach(func() {
		slaveClient = NewSlaveClient()

		fakeServer = ghttp.NewServer()

		serverUrl, _ := url.Parse(fakeServer.URL())
		hostPort := strings.Split(serverUrl.Host, ":")
		serverHost = hostPort[0]
		serverPort, _ = strconv.Atoi(hostPort[1])
	})

	Describe("Execute", func() {

		var err error
		var command osaModel.CommandRequest
		var slave model.Slave

		BeforeEach(func() {
			slave = model.Slave{
				Host: serverHost,
				Port: serverPort,
			}
		})

		Context("when the salve returns non-2XX code", func() {

			BeforeEach(func() {
				fakeServer.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/jobs"),
					ghttp.RespondWithJSONEncoded(http.StatusInternalServerError, []byte{})),
				)

				_, err = slaveClient.Execute(command, slave)
			})

			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})

		})

		Context("when the slave returns 200", func() {
			var jobId string

			BeforeEach(func() {
				jobId = "mega-secret"

				fakeServer.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/jobs"),
					ghttp.RespondWithJSONEncoded(http.StatusOK, osaModel.Job{Id: jobId})),
				)

			})

			It("should return the job id", func() {
				actualJobId, err := slaveClient.Execute(command, slave)
				Expect(err).ToNot(HaveOccurred())
				Expect(actualJobId).To(Equal(jobId))
			})

		})

	})

	Describe("ExecuteAsync", func() {
		var err error
		var command osaModel.CommandRequest
		var slave model.Slave

		BeforeEach(func() {
			slave = model.Slave{
				Host: serverHost,
				Port: serverPort,
			}
		})

		Context("when the salve returns non-2XX code", func() {

			BeforeEach(func() {
				fakeServer.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/jobs"),
					ghttp.RespondWithJSONEncoded(http.StatusInternalServerError, []byte{})),
				)

				_, err = slaveClient.ExecuteAsync(command, slave)
			})

			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})

		})

		Context("when the slave returns 200", func() {
			var jobId string

			BeforeEach(func() {
				jobId = "mega-secret"

				fakeServer.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/jobs"),
					ghttp.RespondWithJSONEncoded(http.StatusOK, osaModel.Job{Id: jobId})),
				)

			})

			It("should return the job id", func() {
				actualJobId, err := slaveClient.ExecuteAsync(command, slave)
				Expect(err).ToNot(HaveOccurred())
				Expect(actualJobId).To(Equal(jobId))
			})

		})
	})

	Describe("Job", func() {
		var slave model.Slave

		BeforeEach(func() {
			slave = model.Slave{
				Host: serverHost,
				Port: serverPort,
			}
		})

		Context("when the slave returns non-2XX status", func() {

			BeforeEach(func() {
				fakeServer.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/jobs", "id=10"),
					ghttp.RespondWithJSONEncoded(http.StatusInternalServerError, []byte{})),
				)
			})

			It("should return an error", func() {
				_, err := slaveClient.Job("10", slave)
				Expect(err).To(HaveOccurred())
			})

		})

		Context("when the slave returns actual job", func() {
			var job osaModel.Job

			BeforeEach(func() {
				job = osaModel.Job{
					Id: "10",
				}
				fakeServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/jobs", "id=10"),
						ghttp.RespondWithJSONEncoded(http.StatusOK, job)),
				)
			})

			It("should return the job", func() {
				actualJob, err := slaveClient.Job("10", slave)
				Expect(err).ToNot(HaveOccurred())
				Expect(actualJob).To(Equal(job))
			})

		})

	})

})
