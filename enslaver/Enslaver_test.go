package enslaver_test

import (
	. "github.com/Bo0mer/enslaver/enslaver"
	"github.com/Bo0mer/enslaver/model"
	"github.com/Bo0mer/enslaver/slaveclient/fakes"

	osaModel "github.com/Bo0mer/os-agent/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Enslaver", func() {
	var enslaver Enslaver
	var slave model.Slave
	var slaveClient *fakes.FakeSlaveClient

	BeforeEach(func() {
		slaveClient = new(fakes.FakeSlaveClient)
		enslaver = NewEnslaver(slaveClient)

		slave = model.Slave{
			Id:   "slave-id",
			Host: "10.4.0.2",
			Port: 8080,
		}
		enslaver.Register(slave)
	})

	Describe("Slaves", func() {

		It("should be able to get registered slaves", func() {
			slaves := enslaver.Slaves()
			Expect(slaves).To(HaveLen(1))
			Expect(slaves[0]).To(Equal(slave))
		})

	})

	Describe("Execute", func() {
		var jobRequest model.JobRequest

		BeforeEach(func() {
			jobRequest = model.JobRequest{}
		})

		Context("when no slaves are present", func() {

			It("should return an error", func() {
				_, err := enslaver.Execute(jobRequest)
				Expect(err).To(HaveOccurred())
			})

		})

		Context("when some of the slaves are missing", func() {

			BeforeEach(func() {
				jobRequest.Slaves = []model.Slave{model.Slave{Id: "missing"}}
			})

			It("should return an error", func() {
				_, err := enslaver.Execute(jobRequest)
				Expect(err).To(HaveOccurred())
			})

		})

		Context("when all of the slaves are present", func() {

			Context("when the request is sync", func() {
				var err error
				var job model.Job
				var slaveJobResponse osaModel.Job

				BeforeEach(func() {
					jobRequest.Async = false
					jobRequest.Slaves = []model.Slave{slave}
					jobRequest.Command = osaModel.CommandRequest{
						Name: "yooo",
					}

					slaveJobResponse = osaModel.Job{
						Status: osaModel.JOB_COMPLETED,
						Result: osaModel.CommandResponse{
							Stdout: "hehe, I'm here!",
						},
					}

					slaveClient.JobReturns(slaveJobResponse, nil)

					job, err = enslaver.Execute(jobRequest)
				})

				It("should have called the enslaver client", func() {
					Expect(slaveClient.ExecuteCallCount()).To(Equal(1))
					actualCommand, actualSlave := slaveClient.ExecuteArgsForCall(0)
					Expect(actualCommand).To(Equal(jobRequest.Command))
					Expect(actualSlave).To(Equal(slave))
				})

				It("should not have returned an error", func() {
					Expect(err).ToNot(HaveOccurred())
				})

				It("should return result for every slave", func() {
					Expect(len(job.Results)).To(Equal(1))
				})

				It("should return the correct result for each slave", func() {
					slaveResult := job.Results[0]
					Expect(slaveResult.Status).To(Equal(model.JobStatus(slaveJobResponse.Status)))
					Expect(slaveResult.Result).To(Equal(slaveJobResponse.Result))
				})

				It("should return finished the job", func() {
					Expect(job.Status).To(Equal(model.JOB_COMPLETED))
				})

			})

		})
	})

})
