package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvironment(t *testing.T) {
	envName := "TEST_ENV"
	expectedValue := "test"
	os.Setenv(envName, expectedValue)

	actualValue := GetEnvironment(envName)

	assert.Equal(t, expectedValue, actualValue)
	assert.NotEqual(t, "not test", actualValue)
}

func TestGetEnvironmentNotFound(t *testing.T) {
	envName := "TEST_ENV"
	expectedValue := ""
	os.Setenv(envName, expectedValue)

	actualValue := GetEnvironment(envName)

	assert.Equal(t, expectedValue, actualValue)
	assert.NotEqual(t, "not test", actualValue)
}
