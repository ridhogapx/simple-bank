package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	arg := CreateEntryParams{
		AccountID: 3,
		Amount:    6000,
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	// Test Case
	require.Nil(t, err)
	require.NotEmpty(t, entry)

	// Should be equal as we input it
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

}
