package db

import (
	"context"
	"testing"
	"time"

	"github.com/ardaatahan/simplebank/util"

	"github.com/stretchr/testify/require"
)

func createTestedRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       util.RandomName(),
		HashedPassword: "secret",
		FullName:       util.RandomName(),
		Email:          util.RandomEmail(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)
	return user
}

func TestCreateUser(t *testing.T) {
	createTestedRandomUser(t)
}

func TestGetUser(t *testing.T) {
	firstUser := createTestedRandomUser(t)
	secondUser, err := testQueries.GetUser(context.Background(), firstUser.Username)
	require.NoError(t, err)
	require.NotEmpty(t, secondUser)
	require.Equal(t, firstUser.Username, secondUser.Username)
	require.Equal(t, firstUser.HashedPassword, secondUser.HashedPassword)
	require.Equal(t, firstUser.FullName, secondUser.FullName)
	require.Equal(t, firstUser.Email, secondUser.Email)
	require.WithinDuration(t, firstUser.PasswordChangedAt, secondUser.PasswordChangedAt, time.Second)
	require.WithinDuration(t, firstUser.CreatedAt, secondUser.CreatedAt, time.Second)
}
