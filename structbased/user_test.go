package user_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	u "bdd-tutorial/structbased"
)

var _ = Describe("User Struct", Label("integration", "user"), func() {
	var user u.User

	BeforeEach(func() {
		user = u.NewUser()
	})

	// Null case
	Context("When the user name, email and age not set", func() {
		It("can get the user name", Label("user", "get"), func() {
			Expect(user.GetName()).To(Equal(""))
		})

		It("can get the user email", Label("user", "get"), func() {
			Expect(user.GetEmail()).To(Equal(""))
		})

		It("can get the user age", Label("user", "get"), func() {
			Expect(user.GetAge()).To(Equal(0))
		})
	})

	// Success case
	Context("When the user name, email and age succeed", func() {
		It("can set the user name and validate", func() {
			err := user.SetName("hugo")
			Expect(err).To(BeNil())
			Expect(user.GetName()).To(Equal("HUGO")) // because variable toUppering when saved
		})

		It("can set the user email and validate", func() {
			err := user.SetEmail("hugo@hepsiburada.com")
			Expect(err).To(BeNil())
			Expect(user.GetEmail()).To(Equal("hugo@hepsiburada.com"))
		})

		It("can set the user age and validate", func() {
			err := user.SetAge("15")
			Expect(err).To(BeNil())
			Expect(user.GetAge()).To(Equal(15))
		})
	})

	// Error case
	Context("When the user email and age failed", func() {

		// occurs error and return null string
		It("if email not valid", func() {
			err := user.SetEmail("hugo")
			Expect(err).To(MatchError(u.ErrInvalid))
			Expect(user.GetEmail()).To(Equal(""))
		})

		// occurs error and return 0
		It("if age is negative", func() {
			err := user.SetAge("-15")
			Expect(err).To(MatchError(u.ErrInvalid))
			Expect(user.GetAge()).To(Equal(0))
		})

		// occurs error and return 0
		It("if age is zero", func() {
			err := user.SetAge("0")
			Expect(err).To(MatchError(u.ErrInvalid))
			Expect(user.GetAge()).To(Equal(0))
		})

		// occurs error and return 0
		It("if age is string", func() {
			err := user.SetAge("string")
			Expect(err).To(MatchError(u.ErrConvert))
			Expect(user.GetAge()).To(Equal(0))
		})
	})
})
