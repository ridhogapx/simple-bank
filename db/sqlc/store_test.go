package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransaction(t *testing.T) {
	// Create new object
	store := NewStore(testDB)

	// Generate random account
	account_a := generateRandomAccount(t)
	account_b := generateRandomAccount(t)

	// Balanace that we're gonna transfer
	amount := int64(5)

	// Should be run in n concurrent
	n := 5

	// Channel for handling communication data
	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account_a.ID,
				ToAccountID:   account_b.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		errResult := <-errs
		transferResult := <-results

		assert.Nil(t, errResult)
		assert.NotEmpty(t, transferResult)

	}

}
