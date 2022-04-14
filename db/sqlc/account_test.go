package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ardaatahan/simplebank/util"

	"github.com/stretchr/testify/require"
)

func createTestedRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
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
	createTestedRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	firstAccount := createTestedRandomAccount(t)
	secondAccount, err := testQueries.GetAccount(context.Background(), firstAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, secondAccount)
	require.Equal(t, firstAccount.ID, secondAccount.ID)
	require.Equal(t, firstAccount.Owner, secondAccount.Owner)
	require.Equal(t, firstAccount.Balance, secondAccount.Balance)
	require.Equal(t, firstAccount.Currency, secondAccount.Currency)
	require.WithinDuration(t, firstAccount.CreatedAt, secondAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	firstAccount := createTestedRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      firstAccount.ID,
		Balance: util.RandomBalance(),
	}
	secondAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, secondAccount)
	require.Equal(t, firstAccount.ID, secondAccount.ID)
	require.Equal(t, firstAccount.Owner, secondAccount.Owner)
	require.Equal(t, arg.Balance, secondAccount.Balance)
	require.Equal(t, firstAccount.Currency, secondAccount.Currency)
	require.WithinDuration(t, firstAccount.CreatedAt, secondAccount.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	firstAccount := createTestedRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), firstAccount.ID)
	require.NoError(t, err)
	secondAccount, err := testQueries.GetAccount(context.Background(), firstAccount.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, secondAccount)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createTestedRandomAccount(t)
	}
	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	require.Len(t, accounts, 5)
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
