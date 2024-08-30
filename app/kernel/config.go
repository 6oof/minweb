package kernel

import (
	"fmt"

	"github.com/6oof/minweb/app/configs"
	"github.com/spf13/viper"
)

// InitConfig initializes Viper for framework configuration management.
func InitConfig() *Config {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil && viper.ConfigFileUsed() != "" {
		// Handle missing .env file if necessary
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
