package models

// User model contains data of the user being me
type User struct {
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Stack     string    `json:"stack"`
}