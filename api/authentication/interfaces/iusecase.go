package interfaces

import "seafarer-backend/api/authentication/router/requests"

type IAuthenticationUseCase interface {
	Login(input *requests.LoginRequest) (jwt interface{}, err error)
}