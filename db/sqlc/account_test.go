package db

import (
	"context"
	"database/sql"
	generator "go-backend-master/painless/models/generator"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    generator.RandomOwner(),
		Balance:  generator.RandomMoney(),
		Currency: generator.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)

}

func TestGetAccount(t *testing.T) {
	account_one := createRandomAccount(t)
	account_two, err := testQueries.GetAccount(context.Background(), account_one.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account_two)

	require.Equal(t, account_one.ID, account_two.ID)
	require.Equal(t, account_one.Owner, account_two.Owner)
	require.Equal(t, account_one.Balance, account_two.Balance)
	require.Equal(t, account_one.Currency, account_two.Currency)

	require.WithinDuration(t, account_one.CreatedAt, account_two.CreatedAt, time.Second)

}

func TestUpdateAccount(t *testing.T) {
	account_one := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account_one.ID,
		Balance: generator.RandomMoney(),
	}

	account_two, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account_two)

	require.Equal(t, account_one.ID, account_two.ID)
	require.Equal(t, account_one.Owner, account_two.Owner)
	require.Equal(t, arg.Balance, account_two.Balance)
	require.Equal(t, account_one.Currency, account_two.Currency)

	require.WithinDuration(t, account_one.CreatedAt, account_two.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	account_one := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account_one.ID)
	require.NoError(t, err)

	account_two, err := testQueries.GetAccount(context.Background(), account_one.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account_two)

}
