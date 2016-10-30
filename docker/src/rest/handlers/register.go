package handlers

import (
  "golang.org/x/crypto/bcrypt"
  "encoding/json"
  "fmt"
  "net/http"
  "rest/users"
)

type RegisterUserData struct {
  Email     string `json:"email"`
  Password  string `json:"password"`
  FirstName string `json:"firstName"`
  LastName  string `json:"lastName"`
}

type Token struct {
  Token string `json:"token"`
}

func Register(w http.ResponseWriter, r *http.Request) {
  var userInfo RegisterUserData

  err := json.NewDecoder(r.Body).Decode(&userInfo)
  if err != nil {
    http.Error(w, err.Error(), 400)
    return
  }

  w.Header().Set("Content-Type", "application/json")

  if userInfo.Email == "" {
    sendMessage(w, Message{false, fmt.Sprint("Email parameter missing.")}, http.StatusInternalServerError)
    return
  }

  if userInfo.Password == "" {
    sendMessage(w, Message{false, fmt.Sprint("Password parameter missing.")}, http.StatusInternalServerError)
    return
  }

  if userInfo.FirstName == "" {
    sendMessage(w, Message{false, fmt.Sprint("First name parameter missing.")}, http.StatusInternalServerError)
    return
  }

  if userInfo.LastName == "" {
    sendMessage(w, Message{false, fmt.Sprint("Last name parameter missing.")}, http.StatusInternalServerError)
    return
  }

  uec := users.RedisUserExistChecker{
    Client: client,
  }

  exists, err := users.Exists(uec, userInfo.Email)

  if err != nil {
    panic(err)
  }

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
    sendMessage(w, Message{false, fmt.Sprintf("A user with the email '%s' already exists.", userInfo.Email)}, http.StatusInternalServerError)
  }
}
