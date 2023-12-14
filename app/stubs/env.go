package stubs

import (
	"fmt"
	"os"
)

// GetEnvVar retrieves the value of an environmental variable.
// If the variable is not set, it returns the default value.
func GetEnvVar(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// GetEnvVarOrPanic retrieves the value of an environmental variable.
// If the variable is not set, it panics with an error message.
func GetEnvVarOrPanic(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	panic(fmt.Sprintf("Environmental variable %s is not set", key))
}
