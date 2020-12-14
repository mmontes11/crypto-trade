package config

import "os"

// GetEnv gets a env variable and, if not found, returns the default value
func GetEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
