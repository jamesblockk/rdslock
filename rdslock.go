package rdslock

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

var _client *redis.Client

// InitURL : url ( redis://:qwerty@localhost:6379/1 )
func InitURL(url string) {
	opt, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}

	_client = redis.NewClient(opt)
	ping()
}

// InitAssigned ...
func InitAssigned(in *redis.Client) {
	_client = in
	ping()
}

func ping() {
	if _, err := _client.Ping().Result(); err != nil {
		panic(err)
	}
}

// Lock ...
func Lock(key string, exp time.Duration) (err error) {
	r := _client.SetNX(key, 1, exp)
	ok, err := r.Result()
	if err != nil {
		return
	}
	if !ok {
		err = errors.New("lock failed")
	}

	return
}

// UnLock ...
func UnLock(key string) (err error) {
	r := _client.Del(key)
	num, err := r.Result()
	if err != nil {
		return
	}
	if num == 0 {
		err = errors.New("unlock failed")
	}

	return
}
