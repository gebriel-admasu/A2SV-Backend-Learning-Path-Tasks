package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`    // hashed password
	Role     string `json:"role"` // "admin" or "user"
}
