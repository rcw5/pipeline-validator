package utils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rcw5/pipeline-validator/utils"
)

var _ = Describe("Utils", func() {
	Context("Contains", func() {
		It("Returns true when the array contains a value", func() {
			result := utils.Contains([]string{"val1", "val2"}, "val2")
			Expect(result).To(BeTrue())
		})
		It("Returns false when the array does not contain the value", func() {
			result := utils.Contains([]string{"val1", "val2"}, "val3")
			Expect(result).To(BeFalse())
		})
	})
	Context("CompareArrays", func() {
		var arr1, arr2 []string
		BeforeEach(func() {
			arr1 = []string{"val1", "val2"}
			arr2 = []string{"val2", "val3"}
		})
		It("Returns elements present in both arrays", func() {
			_, _, inBoth := utils.CompareArrays(arr1, arr2)
			Expect(inBoth).To(ConsistOf("val2"))
		})
		It("Returns elements present only in the left hand arrays", func() {
			inLeft, _, _ := utils.CompareArrays(arr1, arr2)
			Expect(inLeft).To(ConsistOf("val1"))
		})
		It("Returns elements present only in the right hand arrays", func() {
			_, inRight, _ := utils.CompareArrays(arr1, arr2)
			Expect(inRight).To(ConsistOf("val3"))
		})
		It("Returns unique arrays", func() {
			arr1 = append(arr1, arr1...)
			arr2 = append(arr2, arr2...)
			inLeft, inRight, inBoth := utils.CompareArrays(arr1, arr2)
			Expect(inLeft).To(ConsistOf("val1"))
			Expect(inRight).To(ConsistOf("val3"))
			Expect(inBoth).To(ConsistOf("val2"))
		})

	})
})
