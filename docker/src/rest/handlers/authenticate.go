package handlers

import (
  "golang.org/x/crypto/bcrypt"
  "fmt"
  "net/http"
  "encoding/json"
  "rest/users"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
  var userInfo RegisterUserData

  err := json.NewDecoder(r.Body).Decode(&userInfo)

  w.Header().Set("Content-Type", "application/json")

  if userInfo.Email == "" {
    sendMessage(w, Message{false, fmt.Sprint("Email parameter missing.")}, http.StatusInternalServerError)
    return
  }

  if userInfo.Password == "" {
    sendMessage(w, Message{false, fmt.Sprint("Password parameter missing.")}, http.StatusInternalServerError)
    return
  }

  uec := users.RedisUserExistChecker{
    Client: client,
  }

  exists, err := users.Exists(uec, userInfo.Email)

  if err != nil {
    panic(err)
  }

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
      sendMessage(w, Message{false, fmt.Sprint("You've entered the wrong password!")}, http.StatusInternalServerError)
      return
    }

    token, err := getToken(user, w, r)

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
