package users

import (
	"gopkg.in/redis.v4"
	"strconv"
  "fmt"
)

type UserPersister interface {
	Persist(user User) error
}

type RedisUserPersister struct {
	Client *redis.Client
}

func Persist(up UserPersister, user User) error {
	return up.Persist(user)
}

func (rup RedisUserPersister) Persist(user User) error {
	return rup.Client.HMSet(fmt.Sprintf("user:%s", user.Email), map[string]string{
		"encryptedPassword": user.EncryptedPassword,
		"isAdmin":           strconv.FormatBool(user.IsAdmin),
		"firstName":         user.FirstName,
		"lastName":          user.LastName,
	}).Err()
}
