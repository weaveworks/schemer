package definition

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSchemaDefinition(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Definition suite")
}
