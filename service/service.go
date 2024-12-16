package service

import (
	"fmt"
	"study-case/models"
	"study-case/repository"
	"study-case/utils"
)

// CreateTransaction melakukan logika transaksi
func CreateTransaction(t models.Transaction, tenor int) error {
	// Validasi tenor
	if !utils.IsValidTenor(tenor) {
		return fmt.Errorf("tenor tidak valid")
	}

	// Fetch current customer limit
	currentLimit, err := repository.GetCustomerLimit(t.CustomerID, tenor)
	if err != nil {
		return err
	}

	// Check if limit is sufficient
	if currentLimit < t.OTR {
		return fmt.Errorf("insufficient limit")
	}

	// Update limit after transaction
	newLimit := currentLimit - t.OTR
	if err := repository.UpdateCustomerLimit(t.CustomerID, tenor, newLimit); err != nil {
		return err
	}

	// Insert transaction into database
	if err := repository.InsertTransactionToDB(t); err != nil {
		return err
	}

	return nil
}
