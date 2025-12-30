package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

// Config holds application configuration
type Config struct {
	DatabaseType string `yaml:"database_type" mapstructure:"database_type"`
	DatabaseURL  string `yaml:"database_url" mapstructure:"database_url"`
	Environment  string `yaml:"environment" mapstructure:"environment"`
	
	// CLI-specific configurations
	DefaultUserID string `yaml:"default_user_id" mapstructure:"default_user_id"`
	OutputFormat  string `yaml:"output_format" mapstructure:"output_format"`
	ColorOutput   bool   `yaml:"color_output" mapstructure:"color_output"`
	
	// Optional configurations
	MaxUploadSize  int64 `yaml:"max_upload_size" mapstructure:"max_upload_size"`
	EnableLiveness bool  `yaml:"enable_liveness" mapstructure:"enable_liveness"`
	WorkerCount    int   `yaml:"worker_count" mapstructure:"worker_count"`
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
	if cfg.Environment == "" {
		cfg.Environment = "development"
	}
	if cfg.MaxUploadSize == 0 {
		cfg.MaxUploadSize = 100 * 1024 * 1024 // 100MB
	}
	if cfg.WorkerCount == 0 {
		cfg.WorkerCount = 5
	}
	if cfg.DefaultUserID == "" {
		cfg.DefaultUserID = "default"
	}
	if cfg.OutputFormat == "" {
		cfg.OutputFormat = "table"
	}
	// ColorOutput defaults to true
	if !cfg.ColorOutput {
		cfg.ColorOutput = true
	}

	return &cfg, nil
}

// LoadConfigWithViper reads configuration using Viper (supports env vars and multiple formats)
func LoadConfigWithViper(configPath string) (*Config, error) {
	v := viper.New()
	
	// Set config file
	if configPath != "" {
		v.SetConfigFile(configPath)
	} else {
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath("./configs")
		v.AddConfigPath(".")
		v.AddConfigPath("$HOME/.sitemapper")
	}
	
	// Set environment variable support
	v.SetEnvPrefix("SITEMAPPER")
	v.AutomaticEnv()
	
	// Set defaults
	v.SetDefault("environment", "development")
	v.SetDefault("max_upload_size", 100*1024*1024)
	v.SetDefault("worker_count", 5)
	v.SetDefault("default_user_id", "default")
	v.SetDefault("output_format", "table")
	v.SetDefault("color_output", true)
	
	// Read config
	if err := v.ReadInConfig(); err != nil {
		// It's okay if config file doesn't exist, we'll use defaults
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config: %w", err)
		}
	}
	
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	
	return &cfg, nil
}

