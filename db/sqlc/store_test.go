package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)
	firstAccount := createTestedRandomAccount(t)
	secondAccount := createTestedRandomAccount(t)
	numOfConcurrentTx := 5
	transferAmount := int64(10)
	errs := make(chan error)
	results := make(chan TransferTxResult)
	for i := 0; i < numOfConcurrentTx; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: firstAccount.ID,
				ToAccountID:   secondAccount.ID,
				Amount:        transferAmount,
			})
			results <- result
			errs <- err
		}()
	}
	for i := 0; i < numOfConcurrentTx; i++ {
		result := <-results
		err := <-errs

		require.NoError(t, err)
		require.NotEmpty(t, result)

		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, firstAccount.ID, transfer.FromAccountID)
		require.Equal(t, secondAccount.ID, transfer.ToAccountID)
		require.Equal(t, transferAmount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)
		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, firstAccount.ID, fromEntry.AccountID)
		require.Equal(t, -transferAmount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)
		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, secondAccount.ID, toEntry.AccountID)
		require.Equal(t, transferAmount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)
		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)
	}
}
