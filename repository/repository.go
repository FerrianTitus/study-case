// repository.go
package repository

import (
	"database/sql"
	"log"
	"study-case/models"
)

// Database instance
var db *sql.DB

// Initialize DB
func InitDB() {
	var err error
	dsn := "user:password@tcp(localhost:3306)/pt_xyz"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected successfully")
}

// Fetch all customers from DB
func GetAllCustomersFromDB() ([]models.Customer, error) {
	rows, err := db.Query("SELECT customer_id, nik, full_name, salary FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var c models.Customer
		if err := rows.Scan(&c.CustomerID, &c.NIK, &c.FullName, &c.Salary); err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

// Get customer limit
func GetCustomerLimit(customerID int, tenor int) (float64, error) {
	var limitAmount float64
	err := db.QueryRow("SELECT limit_amount FROM customer_limits WHERE customer_id = ? AND tenor = ?", customerID, tenor).Scan(&limitAmount)
	if err != nil {
		return 0, err
	}
	return limitAmount, nil
}

// Update customer limit
func UpdateCustomerLimit(customerID int, tenor int, newLimit float64) error {
	_, err := db.Exec("UPDATE customer_limits SET limit_amount = ? WHERE customer_id = ? AND tenor = ?", newLimit, customerID, tenor)
	return err
}

// Insert transaction to DB
func InsertTransactionToDB(t models.Transaction) error {
	stmt, err := db.Prepare("INSERT INTO transactions (contract_number, customer_id, otr, admin_fee) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.ContractNumber, t.CustomerID, t.OTR, t.AdminFee)
	return err
}
