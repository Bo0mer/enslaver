package facade_test

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Bo0mer/enslaver/enslaver/fakes"
	. "github.com/Bo0mer/enslaver/facade"
	"github.com/Bo0mer/enslaver/model"

	osaModel "github.com/Bo0mer/os-agent/model"
	serverfakes "github.com/Bo0mer/os-agent/server/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EnslaverFacade", func() {

	var facade EnslaverFacade
	var fakeEnslaver *fakes.FakeEnslaver
	var req *serverfakes.FakeRequest
	var resp *serverfakes.FakeResponse

	var slave model.Slave

	BeforeEach(func() {
		fakeEnslaver = new(fakes.FakeEnslaver)
		facade = NewEnslaverFacade(fakeEnslaver)
		req = new(serverfakes.FakeRequest)
		resp = new(serverfakes.FakeResponse)
	})

	var itBehavesLikeBadRequest = func() {
		It("should return status code bad request", func() {
			Expect(resp.SetStatusCodeArgsForCall(0)).To(Equal(http.StatusBadRequest))
		})
	}

	Describe("RegisterSlave", func() {

		Context("when the request is invalid", func() {

			BeforeEach(func() {
				req.BodyReturns([]byte("invalid request"))
				facade.RegisterSlave(req, resp)
			})

			It("should have not called the enslaver", func() {
				Expect(fakeEnslaver.RegisterCallCount()).To(Equal(0))
			})

			itBehavesLikeBadRequest()

		})

		Context("when the request is valid", func() {

			BeforeEach(func() {
				slave = model.Slave{
					Id:   "slave-id",
					Host: "10.0.4.2",
					Port: 8081,
				}
				requestBody, _ := json.Marshal(slave)
				req.BodyReturns(requestBody)

				facade.RegisterSlave(req, resp)
			})

			It("should have called the enslaver", func() {
				Expect(fakeEnslaver.RegisterCallCount()).To(Equal(1))

				actualSlave := fakeEnslaver.RegisterArgsForCall(0)
				Expect(actualSlave).To(Equal(slave))
			})

			It("should have returned status code 200 OK", func() {
				Expect(resp.SetStatusCodeArgsForCall(0)).To(Equal(http.StatusOK))
			})

		})
	})

	Describe("CreateJob", func() {

		Context("when the request is invalid", func() {

			BeforeEach(func() {
				req.BodyReturns([]byte("invalid request"))
				facade.CreateJob(req, resp)
			})

			It("should not have called the enslaver", func() {
				Expect(fakeEnslaver.ExecuteCallCount()).To(Equal(0))
			})

			itBehavesLikeBadRequest()

		})

		Context("when the request is valid", func() {
			var jobRequest model.JobRequest
			var job model.Job

			BeforeEach(func() {
				jobRequest = model.JobRequest{
					Slaves: []model.Slave{slave},
					Command: osaModel.CommandRequest{
						Name: "whoami?",
					},
				}

				requestBody, _ := json.Marshal(jobRequest)
				req.BodyReturns(requestBody)
			})

			Context("when the execution is successful", func() {

				BeforeEach(func() {
					job = model.Job{
						Id:     "the-one-and-only",
						Status: model.JOB_COMPLETED,
					}
					fakeEnslaver.ExecuteReturns(job, nil)

					facade.CreateJob(req, resp)
				})

				It("should have called the enslaver", func() {
					Expect(fakeEnslaver.ExecuteCallCount()).To(Equal(1))
					Expect(fakeEnslaver.ExecuteArgsForCall(0)).To(Equal(jobRequest))
				})

				It("should have returned status code 200 OK", func() {
					Expect(resp.SetStatusCodeArgsForCall(0)).To(Equal(http.StatusOK))
				})

			})

			Context("when the execution fails", func() {

				BeforeEach(func() {
					fakeEnslaver.ExecuteReturns(model.Job{}, errors.New("Division by zero!"))
					facade.CreateJob(req, resp)
				})

				It("should have returned status code 500", func() {
					Expect(resp.SetStatusCodeArgsForCall(0)).To(Equal(http.StatusInternalServerError))
				})

			})

		})

	})

})
