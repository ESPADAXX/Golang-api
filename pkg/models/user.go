package models

import "errors"

// Role defines an enum-like type for user roles
type Role string

const (
	RoleClient  Role = "client"
	RoleAdmin   Role = "admin"
	RoleSupAdmin Role = "supadmin"
)

// Validate checks if the role is valid
func (r Role) Validate() error {
	switch r {
	case RoleClient, RoleAdmin, RoleSupAdmin:
		return nil
	default:
		return errors.New("invalid role")
	}
}

// User defines a user with an ID, name, email, and role
type User struct {
	ID    string `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" validate:"required,email"`
	Role  Role   `json:"role" bson:"role"`
}
