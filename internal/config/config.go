package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config holds application configuration
type Config struct {
	DatabaseType string `yaml:"database_type"`
	DatabaseURL  string `yaml:"database_url"`
	ServerPort   int    `yaml:"server_port"`
	Environment  string `yaml:"environment"`
	
	// Optional configurations
	MaxUploadSize   int64  `yaml:"max_upload_size"`
	EnableLiveness  bool   `yaml:"enable_liveness"`
	WorkerCount     int    `yaml:"worker_count"`
}

// LoadConfig reads and parses the configuration file
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	// Set defaults
	if cfg.ServerPort == 0 {
		cfg.ServerPort = 8080
	}
	if cfg.Environment == "" {
		cfg.Environment = "development"
	}
	if cfg.MaxUploadSize == 0 {
		cfg.MaxUploadSize = 100 * 1024 * 1024 // 100MB
	}
	if cfg.WorkerCount == 0 {
		cfg.WorkerCount = 5
	}

	return &cfg, nil
}

