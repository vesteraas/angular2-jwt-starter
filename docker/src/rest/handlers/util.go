package handlers

import (
  "net/http"
  "encoding/json"
  "gopkg.in/redis.v4"
  "fmt"
  "github.com/dgrijalva/jwt-go"
  "strings"
  "time"
  "rest/users"
  "context"
)

var client = getClient()

func getClient() *redis.Client {
  client := redis.NewClient(&redis.Options{
    Addr:     "redis:6379",
    Password: "", // no password set
    DB:       0,  // use default DB
  })

  fmt.Println("Connected to server")

  return client
}

type Message struct {
  Success bool   `json:"success"`
  Message string `json:"message"`
}

func sendMessage(w http.ResponseWriter, message Message, status int) {
  w.WriteHeader(status)
  if err := json.NewEncoder(w).Encode(message); err != nil {
    panic(err)
  }
}

type Key int

const MyKey Key = 0

type Claims struct {
  Email   string `json:"email"`
  IsAdmin bool   `json:"isAdmin"`
  jwt.StandardClaims
}

func getToken(user users.User, w http.ResponseWriter, r *http.Request) (signedToken string, err error) {
  claims := Claims{
    user.Email,
    user.IsAdmin,
    jwt.StandardClaims{
      ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
      Issuer:    "localhost:8080",
    },
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

  signedToken, err = token.SignedString([]byte("secret"))

  return
}

func Validate(page http.HandlerFunc) http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    authorizationHeader := r.Header.Get("Authorization")

    if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer") {
      http.Error(w, "Not authorized", 401)
      return
    }

    tokenValue := strings.Split(authorizationHeader, " ")[1]

    token, err := jwt.ParseWithClaims(tokenValue, &Claims{}, func(token *jwt.Token) (interface{}, error) {
      if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("Unexpected signing method")
      }
      return []byte("secret"), nil
    })

    if err != nil {
      http.Error(w, "Not authorized", 401)
      return
    }

    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
      ctx := context.WithValue(r.Context(), MyKey, *claims)
      page(w, r.WithContext(ctx))
    } else {
      http.Error(w, "Not authorized", 401)
      return
    }
  })
}
