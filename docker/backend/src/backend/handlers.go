package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"backend/users"
)

type Message struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Token struct {
  Token string `json:"token"`
}

func sendMessage(w http.ResponseWriter, message Message, status int) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		panic(err)
	}
}

type RegisterUserData struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var userInfo RegisterUserData

	err := json.NewDecoder(r.Body).Decode(&userInfo)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if userInfo.Email == "" {
		http.Error(w, "Email parameter missing.", http.StatusInternalServerError)
		return
	}

	if userInfo.Password == "" {
		http.Error(w, "Password parameter missing.", http.StatusInternalServerError)
		return
	}

	if userInfo.FirstName == "" {
		http.Error(w, "First name parameter missing.", http.StatusInternalServerError)
		return
	}

	if userInfo.LastName == "" {
		http.Error(w, "Last name parameter missing.", http.StatusInternalServerError)
		return
	}

	uec := users.RedisUserExistChecker{
		Client: client,
	}

	exists, err := users.Exists(uec, userInfo.Email)

	if err != nil {
		panic(err)
	}

  w.Header().Set("Content-Type", "application/json")
	if !exists {
		up := users.RedisUserPersister{
			Client: client,
		}

		encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(userInfo.Password), bcrypt.DefaultCost)

		if err != nil {
			panic(err)
		}

    rur := users.RedisUserRetriever{
      Client: client,
    }

    allUsers, err := users.RetrieveAll(rur)

    if err != nil {
      panic(err)
    }

		user := users.User{
			Email:             userInfo.Email,
			EncryptedPassword: string(encryptedPassword),
			IsAdmin:           len(allUsers) == 0,
			FirstName:         userInfo.FirstName,
			LastName:          userInfo.LastName,
		}

		users.Persist(up, user)

    token, err := getToken(user, w, r);

    if err != nil {
      panic(err)
    }

    if err := json.NewEncoder(w).Encode(Token{
      Token: token,
    }); err != nil {
      panic(err)
    }
	} else {
		sendMessage(w, Message{false, fmt.Sprintf("A user with the email '%s' already exists.", userInfo.LastName)}, http.StatusInternalServerError)
	}
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	var userInfo RegisterUserData

	err := json.NewDecoder(r.Body).Decode(&userInfo)

	if userInfo.Email == "" {
		http.Error(w, "Email parameter missing.", http.StatusInternalServerError)
		return
	}

	if userInfo.Password == "" {
		http.Error(w, "Password parameter missing.", http.StatusInternalServerError)
		return
	}

	uec := users.RedisUserExistChecker{
		Client: client,
	}

	exists, err := users.Exists(uec, userInfo.Email)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if exists {
		ur := users.RedisUserRetriever{
			Client: client,
		}

		user, err := users.Retrieve(ur, userInfo.Email)

		if err != nil {
			panic(err)
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(userInfo.Password))

		if err != nil {
      fmt.Printf(err.Error());
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			return
		}

    token, err := getToken(user, w, r);

    if err != nil {
      panic(err)
    }

    if err := json.NewEncoder(w).Encode(Token{
      Token: token,
    }); err != nil {
      panic(err)
    }
	} else {
		sendMessage(w, Message{false, fmt.Sprintf("A user with the email '%s' does not exists.", userInfo.Email)}, http.StatusInternalServerError)
	}
}

type GetUsersItem struct {
  Email string `json:"email"`
  IsAdmin bool `json:"isAdmin"`
  FirstName string `json:"firstName"`
  LastName string `json:"lastName"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
  result := []GetUsersItem{}

  ur := users.RedisUserRetriever{
    Client: client,
  }

  users, err := users.RetrieveAll(ur)

  if err != nil {
    panic(err)
  }

  for _, user := range users {
    result = append(result, GetUsersItem{
      Email: user.Email,
      IsAdmin: user.IsAdmin,
      FirstName: user.FirstName,
      LastName: user.LastName,
    })
  }

  w.Header().Set("Content-Type", "application/json")
  if err := json.NewEncoder(w).Encode(result); err != nil {
    panic(err)
  }
}
