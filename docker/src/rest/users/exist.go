package users

import (
	"gopkg.in/redis.v4"
  "fmt"
)

type UserExistChecker interface {
	Exists(email string) (bool, error)
}

type RedisUserExistChecker struct {
	Client *redis.Client
}

func Exists(uec UserExistChecker, email string) (bool, error) {
	return uec.Exists(email)
}

func (uec RedisUserExistChecker) Exists(email string) (bool, error) {
	exist, err := uec.Client.Exists(fmt.Sprintf("user:%s", email)).Result()

	return exist, err
}
