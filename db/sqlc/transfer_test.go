package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {
	arg := CreateTransferParams{
		FromAccountID: 3,
		ToAccountID:   5,
		Amount:        3000,
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.Nil(t, err)
	require.NotEmpty(t, transfer)
}
