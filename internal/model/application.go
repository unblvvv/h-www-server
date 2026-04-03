package model

import "time"

type Application struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	AnimalID   string    `json:"animal_id"`
	AnimalName string    `json:"animal_name,omitempty"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Message    string    `json:"message"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}
