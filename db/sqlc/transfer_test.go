package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {
	arg := CreateTransferParams{
		FromAccountID: 5,
		ToAccountID:   3,
		Amount:        10000,
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.Nil(t, err)
	require.NotEmpty(t, transfer)
}

func TestListTransfer(t *testing.T) {
	transfer, err := testQueries.ListTransfer(context.Background(), 3)

	// Case: Make sure there's no error and data is not empty
	require.Nil(t, err)
	require.NotEmpty(t, transfer)
}
