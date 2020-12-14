package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDevelopment(t *testing.T) {
	logger := Init("development")
	assert.NotNil(t, logger)
}

func TestInitProduction(t *testing.T) {
	logger := Init("production")
	assert.NotNil(t, logger)
}
