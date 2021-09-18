package libraries

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisLibrary struct {
	Client *redis.Client
}

func NewRedisLibrary(client *redis.Client) RedisLibrary {
	return RedisLibrary{Client: client}
}

func (lib RedisLibrary) SendToRedis(key string, value interface{}, exp time.Duration) (err error) {
	if err = lib.Client.Set(context.Background(), key, value, exp).Err(); err != nil {
		return err
	}

	return err
}

func (lib RedisLibrary) GetKeyFromRedis(key string) (result string, err error){
	result, err = lib.Client.Get(context.Background(), key).Result()
	if err != nil {
		return result, err
	}

	return result, err
}

func (lib RedisLibrary) RemoveKeyFromRedis(key string) (count int64, err error){
	count, err = lib.Client.Del(context.Background(), key).Result()
	if err != nil{
		return count, err
	}

	return count, err
}

