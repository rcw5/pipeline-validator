package resources_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/rcw5/pipeline-validator/resources"
)

var _ = Describe("Vars", func() {
	var vars Vars
	BeforeEach(func() {
		varsYAML := `key1: value1
key2: value2`
		var err error
		vars, err = NewVars(varsYAML)
		Expect(err).NotTo(HaveOccurred())
	})
	Context("New", func() {
		It("Returns the total number of vars", func() {
			Expect(len(vars)).To(Equal(2))
		})
	})
	Context("Keys", func() {
		It("Returns the names of the secret keys", func() {
			Expect(vars).To(ConsistOf("key1", "key2"))
		})
	})
})
