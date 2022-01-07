package model

import "time"

type Customer struct {
	Id             int32  `json:"id" db:"id"`
	CustomerId     string `json:"customer_id" db:"customer_id"`
	FirstName      string `json:"first_name" db:"first_name"`
	LastName       string `json:"last_name" db:"last_name"`
	Avatar         string
	Description    string    `json:"description"`
	SecretKey      string    `json:"secret_key"`
	AquilaDb       string    `json:"aquila_db"`
	SharableHash   string    `json:"sharable_hash"`
	IsSharable     bool      `json:"is_sharable"`
	DocumentNumber string    `json:"document_number"`
	IsActive       bool      `json:"is_active" db:"is_active"`
	CreatedAt      time.Time `json:"created_at"`
	Token          string
}
