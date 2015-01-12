package facade_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFacade(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Facade Suite")
}
