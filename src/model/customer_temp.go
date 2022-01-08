package model

import (
	"time"
)

type CustomerTemp struct {
	Id             int32     `json:"id" db:"id"`
	CustomerId     string    `json:"customer_id" db:"customer_id"`
	FirstName      string    `json:"first_name" db:"first_name"`
	LastName       string    `json:"last_name" db:"last_name"`
	Avatar         string    `json:"avatar"`
	SecretKey      string    `json:"secret_key" db:"secret_key"`
	AquilaDb       string    `json:"aquila_db" db:"aquila_db_database_name"`
	SharableHash   string    `json:"sharable_hash" db:"shared_hash"`
	IsSharable     bool      `json:"is_sharable" db:"is_sharable"`
	DocumentNumber string    `json:"document_number" db:"document_number"`
	IsActive       bool      `json:"is_active" db:"is_active"`
	IsPermanent    bool      `default:"false" db:"is_permanent"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	Token          string
}
