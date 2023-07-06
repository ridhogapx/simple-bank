package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	arg := CreateEntryParams{
		AccountID: 5,
		Amount:    7000,
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	// Test Case
	require.Nil(t, err)
	require.NotEmpty(t, entry)

	// Should be equal as we input it
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

}

func TestListEntry(t *testing.T) {
	expected := []Entry{
		{
			ID:        1,
			AccountID: 5,
			Amount:    2300,
		},
		{
			ID:        3,
			AccountID: 5,
			Amount:    7000,
		}, {
			ID:        4,
			AccountID: 5,
			Amount:    7000,
		},
	}
	list, err := testQueries.ListEntries(context.Background(), 5)

	// Case: Should not error and data is not empty!
	require.Nil(t, err)
	require.NotEmpty(t, list)

	// Case: Should as we expected
	require.Equal(t, expected, list)

}

func TestUpdateEntry(t *testing.T) {
	arg := UpdateEntryParams{
		ID:     2,
		Amount: 9000,
	}

	updated, err := testQueries.UpdateEntry(context.Background(), arg)

	// Case: Should not error and not empty
	require.Nil(t, err)
	require.NotEmpty(t, updated)
}
