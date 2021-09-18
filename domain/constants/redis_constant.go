package constants

import "time"

const (
	RedisPrefixVerifyOTP = "verify-"
	RedisVerifyOTPExpiredTime = 5 * time.Minute

	RedisPrefixRecoverOTP = "recover-"
	RedisRecoverOTPExpiredTime = 1 * time.Minute
)
