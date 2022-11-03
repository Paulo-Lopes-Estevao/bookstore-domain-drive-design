package person

type Person struct {
	Name             string
	Phone            string `valid:"required"`
	Email            string `valid:"required"`
	Password         string `valid:"required"`
	VerifyEmail      bool   `valid:"-"`
	VerificationCode string `valid:"-"`
}
