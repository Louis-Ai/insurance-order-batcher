package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	OutputDirectory string
	BatchSize       int
}

func LoadConfig() (*Config, error) {

	outputDirectory := os.Getenv("OUTPUT_DIRECTORY")
	if outputDirectory == "" {
		return nil, fmt.Errorf("OUTPUT_DIRECTORY variable is not set")
	}

	batchSizeString := os.Getenv("BATCH_SIZE")
	if outputDirectory == "" {
		return nil, fmt.Errorf("BATCH_SIZE variable is not set")
	}

	batchSize, err := strconv.Atoi(batchSizeString)
	if err != nil || batchSize <= 0 {
		return nil, fmt.Errorf("invalid BATCH_SIZE: %s", batchSizeString)
	}

	return &Config{
		OutputDirectory: outputDirectory,
		BatchSize:       batchSize,
	}, nil
}
