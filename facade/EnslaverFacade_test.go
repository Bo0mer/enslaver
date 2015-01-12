package facade_test

import (
	"encoding/json"
	"net/http"

	"github.com/Bo0mer/enslaver/enslaver/fakes"
	. "github.com/Bo0mer/enslaver/facade"
	"github.com/Bo0mer/enslaver/model"

	serverfakes "github.com/Bo0mer/os-agent/server/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EnslaverFacade", func() {

	var facade EnslaverFacade
	var fakeEnslaver *fakes.FakeEnslaver
	var req *serverfakes.FakeRequest
	var resp *serverfakes.FakeResponse

	Describe("RegisterSlave", func() {

		BeforeEach(func() {
			fakeEnslaver = new(fakes.FakeEnslaver)
			facade = NewEnslaverFacade(fakeEnslaver)
			req = new(serverfakes.FakeRequest)
			resp = new(serverfakes.FakeResponse)
		})

		Context("when the request is invalid", func() {

			BeforeEach(func() {
				req.BodyReturns([]byte("invalid request"))
				facade.RegisterSlave(req, resp)
			})

			It("should have not called the enslaver", func() {
				Expect(fakeEnslaver.RegisterCallCount()).To(Equal(0))
			})

			It("should have returned status Bad Request", func() {
				Expect(resp.SetStatusCodeArgsForCall(0)).To(Equal(http.StatusBadRequest))
			})
		})

		Context("when the request is valid", func() {

			var slave model.Slave

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
})
