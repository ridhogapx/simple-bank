package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	arg := CreateEntryParams{
		AccountID: 5,
		Amount:    2300,
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	// Test Case
	require.Nil(t, err)
	require.NotEmpty(t, entry)
}
