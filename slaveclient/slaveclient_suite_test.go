package slaveclient_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSlaveclient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Slaveclient Suite")
}
