package enslaver_test

import (
	. "github.com/Bo0mer/enslaver/enslaver"
	"github.com/Bo0mer/enslaver/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Enslaver", func() {
	var enslaver Enslaver
	var slave model.Slave

	BeforeEach(func() {
		enslaver = NewEnslaver()
		slave = model.Slave{
			Id:   "slave-id",
			Host: "10.4.0.2",
			Port: 8080,
		}

	})

	It("should be able to register and get registered slaves", func() {
		enslaver.Register(slave)
		slaves := enslaver.Slaves()
		Expect(slaves).To(HaveLen(1))
		Expect(slaves[0]).To(Equal(slave))
	})

})
