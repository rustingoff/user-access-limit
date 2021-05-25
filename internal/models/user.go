package models

import "github.com/trucktrace/internal/privileges"

type User struct {
	ID         uint                        `json:"id"`
	Name       string                      `json:"name"`
	Email      string                      `json:"email"`
	Password   string                      `json:"password"`
	Privileges privileges.PrivelegesStruct `json:"privileges"`
}
