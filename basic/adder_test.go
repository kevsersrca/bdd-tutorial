package adder_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	adder "bdd-tutorial/basic"
)

var _ = Describe("Adder", func() {
	Context("when summands are positive", func() {

		It("adds two numbers", func() {
			sum, err := adder.Add(2, 3)
			Expect(err).NotTo(HaveOccurred())
			Expect(sum).To(Equal(5))
		})

	})

	Context("when summands are zero", func() {

		It("returns the zero error message", func() {
			sum, err := adder.Add(0, 0)
			Expect(err).To(HaveOccurred())
			Expect(sum).To(Equal(0))
		})

	})

	Context("when summands are negative", func() {

		It("returns the invalid error message", func() {
			sum, err := adder.Add(-2, -3)
			Expect(err).To(HaveOccurred())
			Expect(sum).To(Equal(0))
		})

	})
})
