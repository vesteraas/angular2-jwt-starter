package users

import (
	"gopkg.in/redis.v4"
	"strconv"
  "fmt"
  "strings"
)

type UserRetriever interface {
	Retrieve(email string) (User, error)
  RetrieveAll() ([]User, error)
}

type RedisUserRetriever struct {
	Client *redis.Client
}

func Retrieve(ur UserRetriever, email string) (User, error) {
	return ur.Retrieve(email)
}

func RetrieveAll(ur UserRetriever) ([]User, error) {
  return ur.RetrieveAll()
}

func (rur RedisUserRetriever) Retrieve(email string) (User, error) {
	user, err := rur.Client.HMGet(fmt.Sprintf("user:%s", email), "encryptedPassword", "isAdmin", "firstName", "lastName").Result()

	if err != nil {
		return User{}, err
	}

	if user[0] != nil {
		password := user[0].(string)
		isAdmin, err := strconv.ParseBool(user[1].(string))

		if err != nil {
			panic(err)
		}

		firstName := user[2].(string)
		lastName := user[3].(string)

		return User{
			email,
			password,
			isAdmin,
			firstName,
			lastName,
		}, nil
	} else {
		return User{}, err
	}
}

type UserData struct {
  Email string `json:"email"`
  IsAdmin bool `json:"isAdmin"`
  FirstName string `json:"firstName"`
  LastName string `json:"lastName"`
}

func (rur RedisUserRetriever) RetrieveAll() ([]User, error) {
  users := []User{}

  keys, err := rur.Client.Keys("user:*").Result();

  if err != nil {
    return []User{}, err
  }

  pipeline := rur.Client.Pipeline();

  var results = make([]*redis.SliceCmd, 0)

  for _, key := range keys {
    results = append(results, pipeline.HMGet(key, "isAdmin", "firstName", "lastName"))
  }

  _, err = pipeline.Exec()

  if err != nil {
    return []User{}, err
  }

  for i, result := range results {
    current := result.Val();
    isAdmin, err := strconv.ParseBool(current[0].(string))

    if err != nil {
      return []User{}, err
    }

    users = append(users, User{
      Email: strings.Split(keys[i], ":")[1],
      IsAdmin: isAdmin,
      FirstName: current[1].(string),
      LastName: current[2].(string),
    })
  }

  return users, nil
}
