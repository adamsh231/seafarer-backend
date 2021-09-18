package interfaces

import "seafarer-backend/api/authentication/router/requests"

type IAuthenticationUseCase interface {
	Register(input *requests.RegisterRequest) (jwt map[string]string, err error)

	Login(input *requests.LoginRequest) (jwt interface{}, err error)

	SendEmailOTPVerify() (err error)

	OTPVerify(input *requests.OTPVerify) (jwt map[string]string, err error)

	SendEmailOTPRecover(input *requests.OTPEmailRecoverRequest) (err error)

	OTPRecover(input *requests.OTPRecoverRequest) (jwt interface{}, err error)

	ChangePasswordRecover(input *requests.RecoverPasswordRequest) (err error)

	LoginAdmin(input *requests.LoginRequest) (jwt interface{}, err error)
}