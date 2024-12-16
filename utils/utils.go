package utils

import (
	"fmt"
	"time"
)

// FormatDate mengubah format tanggal menjadi "yyyy-mm-dd"
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// ValidateEmail memeriksa apakah email valid
func ValidateEmail(email string) bool {
	// Bisa menggunakan regex untuk validasi email, di sini hanya contoh sederhana
	if len(email) < 5 {
		return false
	}
	return true
}

// ConvertCurrency mengonversi nilai dari float64 ke format string dengan tanda mata uang
func ConvertCurrency(amount float64) string {
	return fmt.Sprintf("$%.2f", amount)
}

// IsValidTenor memvalidasi apakah tenor yang diberikan valid (misalnya antara 1 dan 60 bulan)
func IsValidTenor(tenor int) bool {
	return tenor >= 1 && tenor <= 60
}
