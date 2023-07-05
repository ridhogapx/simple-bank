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

// Test case for getting account by id
func TestGetAccount(t *testing.T) {
	newAccount := generateRandomAccount(t)

	list, err := testQueries.GetAccount(context.Background(), newAccount.ID)

	require.Nil(t, err)
	require.NotEmpty(t, list)

	require.Equal(t, newAccount.ID, list.ID, "ID is not as expected")
	require.Equal(t, newAccount.Owner, list.Owner, "Owner is not as expected")
	require.Equal(t, newAccount.Balance, list.Balance, "Balance is not as expected")

}

func TestUpdateAccount(t *testing.T) {
	account, err := testQueries.GetAccount(context.Background(), 8)

	newChanged := UpdateAccountParams{
		ID:      8,
		Balance: 5000,
	}

	updated, errChanged := testQueries.UpdateAccount(context.Background(), newChanged)

	// Make sure that err is nil, or there is nothing error
	require.Nil(t, err)
	require.Nil(t, errChanged)
	// Make sure that record in database is not empty
	require.NotEmpty(t, account)
	require.NotEmpty(t, updated)

	// Make sure updated is as we expected
	require.Equal(t, newChanged.Balance, updated.Balance)

}

func TestDeleteAccount(t *testing.T) {
	err := testQueries.DeleteAccount(context.Background(), 2)

	// Check if there's error when deleting record
	require.Nil(t, err)

}
