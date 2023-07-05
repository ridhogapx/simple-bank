package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    "Bob",
		Balance:  2000,
		Currency: "USD",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	// Test
	require.Nil(t, err)
	require.NotEmpty(t, account)
}
