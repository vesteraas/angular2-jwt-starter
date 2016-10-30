package handlers

import (
  "net/http"
  "encoding/json"
  "rest/users"
)

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
