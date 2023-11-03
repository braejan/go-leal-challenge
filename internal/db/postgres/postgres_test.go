package postgres_test

import (
	"os"
	"testing"

	"github.com/braejan/go-leal-challenge/internal/db/postgres"
	"github.com/stretchr/testify/assert"
)

func Test_OpenDefautDatabase_ErrorOpenning(t *testing.T) {
	os.Setenv("POSTGRES_USER", "another_user")
	defer os.Unsetenv("POSTGRES_USER")
	_, err := postgres.OpenDefautDatabase()
	assert.NotNil(t, err)
}
