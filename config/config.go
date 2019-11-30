package config

import (
	"github.com/joeshaw/envdecode"
	_ "github.com/joho/godotenv/autoload" // Read .env file
)

// Parameters contains all the configuration parameters.
type Parameters struct {
	APIPort      int    `env:"API_PORT"`
	DatabasePath string `env:"DATABASE_PATH"`
	Debug        bool   `env:"DEBUG_MODE"`
}

// Get checks and load all the necessary environment variables in a struct.
func Get() (*Parameters, error) {
	params := &Parameters{}

	err := envdecode.Decode(params)
	if err != nil {
		return nil, err
	}

	return params, nil
}
