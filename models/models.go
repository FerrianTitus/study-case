package models

import "github.com/golang-jwt/jwt/v5"

// Customer model
type Customer struct {
	CustomerID int     `json:"customer_id"`
	NIK        string  `json:"nik"`
	FullName   string  `json:"full_name"`
	Salary     float64 `json:"salary"`
}

// Transaction model
type Transaction struct {
	TransactionID  int     `json:"transaction_id"`
	ContractNumber string  `json:"contract_number"`
	CustomerID     int     `json:"customer_id"`
	OTR            float64 `json:"otr"`
	AdminFee       float64 `json:"admin_fee"`
}

// Limit model
type Limit struct {
	CustomerID  int     `json:"customer_id"`
	Tenor       int     `json:"tenor"`
	LimitAmount float64 `json:"limit_amount"`
}

// JWT Claims
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
