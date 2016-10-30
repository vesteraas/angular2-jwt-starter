package users

type User struct {
	Email             string
	EncryptedPassword string
	IsAdmin           bool
	FirstName         string
	LastName          string
}
