package configure

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DynamoEndpoint  string
	DynamoRegion    string
	DynamoAccessKey string
	DynamoSecretKey string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	cfg := &Config{
		DynamoEndpoint:  os.Getenv("DYNAMO_ENDPOINT"),
		DynamoRegion:    os.Getenv("DYNAMO_REGION"),
		DynamoAccessKey: os.Getenv("DYNAMO_ACCESS_KEY"),
		DynamoSecretKey: os.Getenv("DYNAMO_SECRET_KEY"),
	}

	if cfg.DynamoEndpoint == "" || cfg.DynamoRegion == "" || cfg.DynamoAccessKey == "" || cfg.DynamoSecretKey == "" {
		return nil, fmt.Errorf("one or more environment variables are not set")
	}

	return cfg, nil
}
