package db

import (
	"testing"
)

func TestTransaction(t *testing.T) {
	// Create new object
	store := NewStore(testDB)

	// Generate random account
	account_a := generateRandomAccount(t)
	account_b := generateRandomAccount(t)

}
