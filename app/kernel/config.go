package kernel

import (
	"fmt"
	"os"

	"github.com/6oof/minweb/app/configs"
	"github.com/spf13/viper"
)

// InitConfig initializes Viper for framework configuration management.
func InitConfig() *Config {
	// Define the filenames
	envFile := ".env"
	exampleFile := ".env.example"

	// Check if .env file exists
	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		// If .env does not exist, check for .env.example
		if _, err := os.Stat(exampleFile); !os.IsNotExist(err) {
			// Copy .env.example to .env
			input, err := os.ReadFile(exampleFile)
			if err != nil {
				panic(fmt.Sprintf("Failed to read %s: %v", exampleFile, err))
			}
			err = os.WriteFile(envFile, input, 0644)
			if err != nil {
				panic(fmt.Sprintf("Failed to write %s: %v", envFile, err))
			}
			fmt.Printf(".env file created from %s\n", exampleFile)
		} else {
			// Neither .env nor .env.example exists
			panic(fmt.Sprintf("No .env or %s file found in the root directory", exampleFile))
		}
	}

	// Set up viper with the .env file
	viper.SetConfigFile(envFile)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Failed to read the configuration file: %v", err))
	}

	// Set default values for configuration
	configs.Defaults()

	return &Config{}
}

// Config holds configuration methods.
type Config struct{}

// Get retrieves a configuration value by key.
func (c *Config) Get(key string) string {
	return viper.GetString(key)
}

// GetOrPanic retrieves a configuration value by key and panics if not set.
func (c *Config) GetOrPanic(key string) string {
	value := viper.GetString(key)
	if value == "" {
		panic(fmt.Sprintf("Environmental variable %s is not set", key))
	}
	return value
}
