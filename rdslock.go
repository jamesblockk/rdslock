package rdslock

import (
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
func InitAssigned(in *redis.Client) (err error) {
	_client = in
	ping()
}

func ping() {
	_, err = _client.Ping().Result()
	if err != nil {
		panic(err)
	}

	return nil
}

// Lock ...
func Lock(key string, exp time.Second) (err error) {
	r := _client.SetNX(key, 1, exp)
	ok, err := r.Result()
	if err != nil {
		return
	}
	if !ok {
		err = error.Errors("lock failed")
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
		err = error.Errors("unlock failed")
	}

	return
}
