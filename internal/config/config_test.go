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

func TestGetIntEnv(t *testing.T) {
	env := "FOO"
	initVal := os.Getenv(env)
	err := os.Setenv(env, "1")
	assert.Nil(t, err)
	defer func() {
		os.Setenv(env, initVal)
	}()

	result := GetIntEnv(env, 0)
	assert.Equal(t, 1, result)

	fallback := 10
	result = GetIntEnv("BAR", fallback)
	assert.Equal(t, fallback, result)
}
