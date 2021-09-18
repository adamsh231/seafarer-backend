package usecase

import (
	"errors"
	"seafarer-backend/api"
	"seafarer-backend/api/authentication/helpers"
	"seafarer-backend/api/authentication/interfaces"
	"seafarer-backend/api/authentication/router/requests"
	"seafarer-backend/api/user/repositories"
	"seafarer-backend/domain/constants/messages"
	"seafarer-backend/domain/models"
	"seafarer-backend/libraries"
)

type AuthenticationUseCase struct {
	*api.Contract
}

func NewAuthenticationUseCase(ucContract *api.Contract) interfaces.IAuthenticationUseCase {
	return &AuthenticationUseCase{ucContract}
}

// todo: breaking change & set contract
const (
	CompanyName    = "PT. Seafarindo"
	CompanyAddress = "Jl. Usaha No.36 Tanjung Priok Jakarta Utara, Indonesia"
	CompanyCountry = "Indonesia"
)

func (uc AuthenticationUseCase) Login(input *requests.LoginRequest) (jwt interface{}, err error) {

	// get user model
	model := models.NewUser()
	repo := repositories.NewUserRepository(uc.Postgres)
	if err = repo.ReadByEmail(input.Email, model); err != nil {
		api.NewErrorLog("Authentication.IsValidCredential", "repo.GetUserByEmail", err.Error())
		return jwt, err
	}

	// check password is valid
	if isValid := helpers.NewHashHelper().CheckHashString(input.Password, model.Password); !isValid {
		return jwt, errors.New(messages.CredentialIsNotMatchMessage)
	}

	// send jwt
	jwtLibrary := libraries.NewJWTLibrary()
	jwt, err = jwtLibrary.GenerateToken(model.ID, model.Name, model.Email, model.IsVerified)
	if err != nil {
		api.NewErrorLog("Authentication.GenerateToken", "jwtHelper.GenerateToken", err.Error())
		return jwt, err
	}

	return jwt, err
}