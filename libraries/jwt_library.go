package libraries

import (
	"errors"
	"seafarer-backend/domain/constants"
	"seafarer-backend/domain/constants/messages"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTLibrary struct {
}

func NewJWTLibrary() JWTLibrary {
	return JWTLibrary{}
}

var secretKey = []byte("Allahumma sholli 'ala sayyidina Muhammad!!!")

func (lib JWTLibrary) GenerateToken(id, name, email string, isVerified bool, isAdminDefault ...bool) (jwt map[string]string, err error) {
	var token string
	var refreshToken string
	var isAdmin bool
	if len(isAdminDefault) > 0 {
		isAdmin = isAdminDefault[0]
	}

	if isAdmin {
		//token
		token, err = lib.generateTokenWithTTLAdmin(id, name, email, constants.JWTTokenLiveTIme)
		if err != nil {
			return jwt, err
		}

		// refresh token
		refreshToken, err = lib.generateTokenWithTTLAdmin(id, name, email, constants.JWTRefreshTokenLiveTime)
		if err != nil {
			return jwt, err
		}
	} else {
		//token
		token, err = lib.generateTokenWithTTL(id, name, email, isVerified, constants.JWTTokenLiveTIme)
		if err != nil {
			return jwt, err
		}

		// refresh token
		refreshToken, err = lib.generateTokenWithTTL(id, name, email, isVerified, constants.JWTRefreshTokenLiveTime)
		if err != nil {
			return jwt, err
		}
	}

	credential := map[string]string{
		constants.JWTResponseToken:        token,
		constants.JWTResponseRefreshToken: refreshToken,
	}

	return credential, err
}

func (lib JWTLibrary) ValidateToken(encodedToken string) (jwt.MapClaims, bool) {

	// parse token
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return encodedToken, errors.New(messages.TokenIsNotValidMessage)
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, false
	}

	// get payload
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	}

	return nil, false
}

func (lib JWTLibrary) generateTokenWithTTL(id, name, email string, isVerified bool, duration time.Duration) (signedToken string, err error) {

	// payload
	payload := jwt.MapClaims{}
	payload[constants.JWTPayloadId] = id
	payload[constants.JWTPayloadName] = name
	payload[constants.JWTPayloadEmail] = email
	payload[constants.JWTPayloadIsVerified] = isVerified
	payload[constants.JWTPayloadIsAdmin] = false
	payload[constants.JWTPayloadTokenLiveTime] = time.Now().Add(duration).Unix()

	// token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload) // Encryption Algorithm

	// signature
	signedToken, err = token.SignedString(secretKey)
	if err != nil {
		return signedToken, err
	}

	return signedToken, err
}

func (lib JWTLibrary) generateTokenWithTTLAdmin(id, name, email string, duration time.Duration) (signedToken string, err error) {

	// payload
	payload := jwt.MapClaims{}
	payload[constants.JWTPayloadId] = id
	payload[constants.JWTPayloadName] = name
	payload[constants.JWTPayloadEmail] = email
	payload[constants.JWTPayloadIsVerified] = true
	payload[constants.JWTPayloadIsAdmin] = true
	payload[constants.JWTPayloadTokenLiveTime] = time.Now().Add(duration).Unix()

	// token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload) // Encryption Algorithm

	// signature
	signedToken, err = token.SignedString(secretKey)
	if err != nil {
		return signedToken, err
	}

	return signedToken, err
}
