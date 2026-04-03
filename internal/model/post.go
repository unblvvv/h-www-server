package model

import "time"

type ASex string

const (
	ASexMale    ASex = "male"
	ASexFemale  ASex = "female"
	ASexUnknown ASex = "unknown"
)

type AStatus string

const (
	AStatusAvailable AStatus = "available"
	AStatusAdopted   AStatus = "adopted"
	AStatusTreatment AStatus = "treatment"
)

type APost struct {
	ID             string
	OrganizationID string
	Name           string
	Age            string
	Sex            ASex
	Description    string
	PhotoURL       *string
	Status         AStatus
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
