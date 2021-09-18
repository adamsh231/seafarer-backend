package usecase

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"seafarer-backend/api"
	"seafarer-backend/api/authentication/helpers"
	"seafarer-backend/api/authentication/interfaces"
	"seafarer-backend/api/authentication/router/requests"
	"seafarer-backend/api/user/repositories"
	adminRepo "seafarer-backend/api/admin/repositories"

	"seafarer-backend/domain/constants"
	"seafarer-backend/domain/constants/messages"
	"seafarer-backend/domain/models"
	"seafarer-backend/libraries"
	"seafarer-backend/utils"
	userUsecase "seafarer-backend/api/user/usecase"
	"time"
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

func (uc AuthenticationUseCase) Register(input *requests.RegisterRequest) (jwt map[string]string, err error) {

	// init
	now := time.Now()
	password, err := helpers.NewHashHelper().HashAndSalt(input.Password)
	if err != nil {
		api.NewErrorLog("AuthenticationUseCase.Register", "helpers.HashAndSalt", err.Error())
		return jwt, err
	}
	id := uuid.New().String()
	model := &models.User{
		ID:         id,
		Name:       input.Name,
		Email:      input.Email,
		Password:   password,
		IsVerified: false,
		CreatedAt:  now,
		UpdatedAt:  now,
		CompanyID:  input.CompanyID,
	}

	// save not verified user
	repo := repositories.NewUserRepository(uc.Postgres)
	if err = repo.Add(model, uc.PostgresTX); err != nil {
		api.NewErrorLog("AuthenticationUseCase.Register", "repo.Add", err.Error())
		return jwt, err
	}

	// send otp
	uc.UserID = model.ID
	uc.UserName = model.Name
	uc.UserEmail = model.Email
	if err = uc.SendEmailOTPVerify(); err != nil {
		return jwt, err
	}

	// generate jwt
	jwtLibrary := libraries.NewJWTLibrary()
	jwt, err = jwtLibrary.GenerateToken(model.ID, model.Name, model.Email, model.IsVerified)
	if err != nil {
		api.NewErrorLog("AuthenticationUseCase.Register", "jwtHelper.GenerateToken", err.Error())
		return jwt, err
	}

	return jwt, err
}

func (uc AuthenticationUseCase) SendEmailOTPVerify() (err error) {

	// get otp from redis
	redisLibrary := libraries.NewRedisLibrary(uc.Redis)
	otp, err := redisLibrary.GetKeyFromRedis(fmt.Sprintf("%s%s", constants.RedisPrefixVerifyOTP, uc.UserEmail))
	if err == redis.Nil {

		// create otp
		otp, err = helpers.NewHashHelper().GenerateOTP(constants.OTPCountNumber)
		if err != nil {
			api.NewErrorLog("AuthenticationUseCase.SendEmailOTPVerify", "helpers.NewHashHelper.GenerateOTP", err.Error())
			return err
		}

		// send to redis
		if err = redisLibrary.SendToRedis(fmt.Sprintf("%s%s", constants.RedisPrefixVerifyOTP, uc.UserEmail), otp, constants.RedisVerifyOTPExpiredTime); err != nil {
			api.NewErrorLog("AuthenticationUseCase.SendEmailOTPVerify", "helpers.NewRedisHelper.SendToRedis", err.Error())
		}

	} else if err != nil {
		api.NewErrorLog("AuthenticationUseCase.SendEmailVerifyOTP", "redisHelper.GetKeyFromRedis", err.Error())
		return err
	}

	// parsing email template
	templateHelper := utils.NewTemplateUtil()
	dataTemplate := constants.MailDataTemplateOTP{
		Name:           uc.UserName,
		OTP:            otp,
		CompanyName:    CompanyName,
		CompanyAddress: CompanyAddress,
		CompanyCountry: CompanyCountry,
	}
	tpl, err := templateHelper.ParseTemplateToBuffer(constants.MailVerifyTemplate, dataTemplate)
	if err != nil {
		api.NewErrorLog("AuthenticationUseCase.SendEmailVerifyOTP", "templateHelper.ParseTemplateToBuffer", err.Error())
		return err
	}

	// send to email
	subject := constants.MailVerifyOTP
	to := []string{uc.UserEmail}
	go uc.Mail.SendMail(subject, CompanyName, tpl.String(), to) // TODO: change to svc-notification with kafka

	return err
}

func (uc AuthenticationUseCase) OTPVerify(input *requests.OTPVerify) (jwt map[string]string, err error) {

	// get otp from redis
	redisLibrary := libraries.NewRedisLibrary(uc.Redis)
	otp, err := redisLibrary.GetKeyFromRedis(fmt.Sprintf("%s%s", constants.RedisPrefixVerifyOTP, uc.UserEmail))
	if err == redis.Nil {
		return jwt, errors.New(messages.OTPIsExpiredMessage)
	}
	if err != nil {
		return jwt, err
	}
	if input.OTP != otp {
		return jwt, errors.New(messages.OTPIsNotMatchMessage)
	}

	// update user to verified
	repo := repositories.NewUserRepository(uc.Postgres)
	if err = repo.UpdateVerifiedByEmail(uc.UserEmail, uc.PostgresTX); err != nil {
		return jwt, err
	}

	// use for jwt payload
	model := &models.User{
		ID:         uc.UserID,
		Name:       uc.UserName,
		Email:      uc.UserEmail,
		IsVerified: true,
	}

	// generate jwt
	jwtLibrary := libraries.NewJWTLibrary()
	jwt, err = jwtLibrary.GenerateToken(model.ID, model.Name, model.Email, model.IsVerified)
	if err != nil {
		api.NewErrorLog("AuthenticationUseCase.OTPVerify", "jwtHelper.GenerateToken", err.Error())
		return jwt, err
	}

	// remove key from redis
	_, err = redisLibrary.RemoveKeyFromRedis(fmt.Sprintf("%s%s", constants.RedisPrefixVerifyOTP, model.Email))
	if err != nil {
		err = nil // because its okay if error
	}

	return jwt, err
}

func (uc AuthenticationUseCase) SendEmailOTPRecover(input *requests.OTPEmailRecoverRequest) (err error) {
	// get user model
	model := models.NewUser()
	repo := repositories.NewUserRepository(uc.Postgres)
	if err = repo.ReadByEmail(input.Email, model); err != nil {
		api.NewErrorLog("AuthenticationUseCase.SendEmailOTPRecover", "repo.GetUserByEmail", err.Error())
		return err
	}

	// get otp from redis
	redisLibrary := libraries.NewRedisLibrary(uc.Redis)
	otp, err := redisLibrary.GetKeyFromRedis(fmt.Sprintf("%s%s", constants.RedisPrefixRecoverOTP, model.Email))
	if err == redis.Nil {

		// create otp
		otp, err = helpers.NewHashHelper().GenerateOTP(constants.OTPCountNumber)
		if err != nil {
			api.NewErrorLog("AuthenticationUseCase.SendEmailOTPRecover", "helpers.NewHashHelper.GenerateOTP", err.Error())
			return err
		}

		// send to redis
		if err = redisLibrary.SendToRedis(fmt.Sprintf("%s%s", constants.RedisPrefixRecoverOTP, model.Email), otp, constants.RedisRecoverOTPExpiredTime); err != nil {
			api.NewErrorLog("AuthenticationUseCase.SendEmailOTPRecover", "helpers.NewRedisHelper.SendToRedis", err.Error())
		}

	} else if err != nil {
		api.NewErrorLog("AuthenticationUseCase.SendEmailOTPRecover", "redisHelper.GetKeyFromRedis", err.Error())
		return err
	}

	// parsing email template
	templateHelper := utils.NewTemplateUtil()
	dataTemplate := constants.MailDataTemplateOTP{
		Email:          model.Email,
		OTP:            otp,
		CompanyName:    CompanyName,
		CompanyAddress: CompanyAddress,
		CompanyCountry: CompanyCountry,
	}
	tpl, err := templateHelper.ParseTemplateToBuffer(constants.MailRecoverTemplate, dataTemplate)
	if err != nil {
		api.NewErrorLog("AuthenticationUseCase.SendEmailOTPRecover", "templateHelper.ParseTemplateToBuffer", err.Error())
		return err
	}

	// send to email
	subject := constants.MailRecoverOTP
	to := []string{model.Email}
	go uc.Mail.SendMail(subject, CompanyName, tpl.String(), to) // TODO: change to svc-notification with kafka

	return err
}

func (uc AuthenticationUseCase) OTPRecover(input *requests.OTPRecoverRequest) (jwt interface{}, err error) {

	// get otp from redis
	redisLibrary := libraries.NewRedisLibrary(uc.Redis)
	otp, err := redisLibrary.GetKeyFromRedis(fmt.Sprintf("%s%s", constants.RedisPrefixRecoverOTP, input.Email))
	if err == redis.Nil {
		return jwt, errors.New(messages.OTPIsExpiredMessage)
	}
	if err != nil {
		return jwt, err
	}
	if input.OTP != otp {
		return jwt, errors.New(messages.OTPIsNotMatchMessage)
	}

	// get user model
	model := models.NewUser()
	repo := repositories.NewUserRepository(uc.Postgres)
	if err = repo.ReadByEmail(input.Email, model); err != nil {
		api.NewErrorLog("AuthenticationUseCase.OTPRecover", "repo.GetUserByEmail", err.Error())
		return jwt, err
	}

	// generate jwt
	jwtLibrary := libraries.NewJWTLibrary()
	jwt, err = jwtLibrary.GenerateToken(model.ID, model.Name, model.Email, model.IsVerified)
	if err != nil {
		api.NewErrorLog("AuthenticationUseCase.OTPVerify", "jwtHelper.GenerateToken", err.Error())
		return jwt, err
	}

	// remove key from redis
	_, err = redisLibrary.RemoveKeyFromRedis(fmt.Sprintf("%s%s", constants.RedisPrefixRecoverOTP, model.Email))
	if err != nil {
		err = nil // because its okay if error
	}

	return jwt, err
}

func (uc AuthenticationUseCase) ChangePasswordRecover(input *requests.RecoverPasswordRequest) (err error) {

	// hashing password
	password, err := helpers.NewHashHelper().HashAndSalt(input.Password)
	if err != nil {
		api.NewErrorLog("AuthenticationUseCase.ChangePasswordRecover", "helpers.HashAndSalt", err.Error())
		return err
	}

	// change password
	userUc := userUsecase.NewUserUseCase(uc.Contract)
	if err = userUc.ChangePassword(password); err != nil {
		return err
	}

	return err
}

func (uc AuthenticationUseCase) LoginAdmin(input *requests.LoginRequest) (jwt interface{}, err error) {

	// get user model
	model := models.NewAdmin()
	repo := adminRepo.NewAdminRepository(uc.Postgres)
	if err = repo.ReadByEmail(input.Email, model); err != nil {
		api.NewErrorLog("AuthenticationUseCase.IsValidCredential", "repo.ReadByEmail", err.Error())
		return jwt, err
	}

	// check password is valid
	if isValid := helpers.NewHashHelper().CheckHashString(input.Password, model.Password); !isValid {
		return jwt, errors.New(messages.CredentialIsNotMatchMessage)
	}

	// send jwt
	jwtLibrary := libraries.NewJWTLibrary()
	jwt, err = jwtLibrary.GenerateToken(model.ID, model.Name, model.Email, false, true)
	if err != nil {
		api.NewErrorLog("AuthenticationUseCase.GenerateToken", "jwtHelper.GenerateToken", err.Error())
		return jwt, err
	}

	return jwt, err
}
