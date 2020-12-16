package config

import (
	"os"
	"strconv"
)

// GetEnv gets a env variable and, if not found, returns the default value
func GetEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// GetIntEnv same as GetEnv but also parses env variable as an integer
func GetIntEnv(key string, defaultValue int) int {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return defaultValue
	}
	intVal, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intVal
}
