package models

import (
	"github.com/google/uuid"
	"time"
)

type Person struct {
	ID          uuid.UUID `json:"id"`
	FullName    string    `json:"full_name"`
	Age         *int      `json:"age,omitempty"`
	Gender      *string   `json:"gender,omitempty"`
	Nationality *string   `json:"nationality,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PeopleFilter struct {
	Page        int
	Limit       int
	Gender      *string
	Nationality *string
	Age         *int
}
