package db

import (
	"context"
	"simple-bank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func generateRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	// Test
	require.Nil(t, err)
	require.NotEmpty(t, account)

	// Should be same as we inserted before
	require.Equal(t, arg.Owner, account.Owner, "Owner is not as expected")
	require.Equal(t, arg.Balance, account.Balance, "Balance is not as expected")
	require.Equal(t, arg.Currency, account.Currency, "Currency is not as expected")

	return account
}

func TestCreateAccount(t *testing.T) {
	generateRandomAccount(t)
}
