package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	env := "FOO"
	initVal := os.Getenv(env)
	err := os.Setenv(env, "FOO")
	assert.Nil(t, err)
	defer func() {
		os.Setenv(env, initVal)
	}()

	result := GetEnv(env, "")
	assert.Equal(t, "FOO", result)

	fallback := "FALLBACK"
	result = GetEnv("BAR", fallback)
	assert.Equal(t, fallback, result)
}
