package constants

const (
	MailVerifyTemplate  = "../../domain/assets/verify_email_otp.html"
	MailVerifyOTP       = "Please verify your email address"
	MailRecoverTemplate = "../../domain/assets/recover_email_otp.html"
	MailRecoverOTP      = "Recover password procedures"
)

type MailDataTemplateOTP struct {
	Name           string
	Email          string
	OTP            string
	CompanyName    string
	CompanyAddress string
	CompanyCountry string
}
