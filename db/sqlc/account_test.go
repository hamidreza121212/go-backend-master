package db

import (
	"context"
	"testing"
	"github.com/stretchr/testify/require"
	generator "go-backend-master/painless/models/generator"
)


func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner: generator.RandomOwner(),
		Balance: generator.RandomMony(),
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
}