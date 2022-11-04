package services

type IConfirmAccount interface {
	ConfirmEmail(VerifyEmail bool, VerificationCode string)
}
