package util_test

import (
	"os"
	"testing"

	"github.com/braejan/go-leal-challenge/pkg/util"
	"github.com/stretchr/testify/assert"
)

func Test_GetEnv_Empty(t *testing.T) {
	value := util.GetEnv("POSTGRES_USER", "empty")
	assert.Equal(t, "empty", value)
}

func Test_GetEnv_Sucess(t *testing.T) {
	os.Setenv("POSTGRES_USER", "postgres")
	defer os.Unsetenv("POSTGRES_USER")
	value := util.GetEnv("POSTGRES_USER", "default")
	assert.Equal(t, "postgres", value)

}
