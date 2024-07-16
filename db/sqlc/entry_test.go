package db

import (
	"context"
	"database/sql"
	"go-backend-master/painless/models/generator"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    generator.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry

}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry_one := createRandomEntry(t, account)
	entry_two, err := testQueries.GetEntry(context.Background(), entry_one.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry_two)

	require.Equal(t, entry_one.ID, entry_two.ID)
	require.Equal(t, entry_one.AccountID, entry_two.AccountID)
	require.Equal(t, entry_one.Amount, entry_two.Amount)

	require.WithinDuration(t, entry_one.CreatedAt, entry_two.CreatedAt, time.Second)

}

func TestDeleteEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry_one := createRandomEntry(t, account)

	err := testQueries.DeleteEntry(context.Background(), entry_one.ID)
	require.NoError(t, err)

	entry_two, err := testQueries.GetEntry(context.Background(), entry_one.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry_two)

}

func TestListEntry(t *testing.T) {

	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}

}
