package constants

import "time"

const (
	JWTTokenLiveTIme        = 2 * time.Hour
	JWTRefreshTokenLiveTime = 24 * time.Hour

	JWTResponseToken        = "token"
	JWTResponseRefreshToken = "refresh_token"

	JWTPayloadId            = "id"
	JWTPayloadName          = "name"
	JWTPayloadEmail         = "email"
	JWTPayloadIsVerified    = "is_verified"
	JWTPayloadIsAdmin       = "is_admin"
	JWTPayloadTokenLiveTime = "exp"
)
