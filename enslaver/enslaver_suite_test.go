package enslaver_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEnslaver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Enslaver Suite")
}
