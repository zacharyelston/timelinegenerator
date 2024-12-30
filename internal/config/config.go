package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Filters struct {
	StartDate         string   `mapstructure:"start_date"`
	EndDate           string   `mapstructure:"end_date"`
	Tags              []string `mapstructure:"tags"`
	ExcludeTags       []string `mapstructure:"exclude_tags"`
	Categories        []string `mapstructure:"categories"`
	ExcludeCategories []string `mapstructure:"exclude_categories"`
}

type Config struct {
	Exporter string  `mapstructure:"exporter"`
	Input    string  `mapstructure:"input"`
	Output   string  `mapstructure:"output"`
	Filters  Filters `mapstructure:"filters"`
}

func LoadConfig() (*Config, error) {
	v := viper.New()

	// Set defaults
	v.SetDefault("exporter", DefaultExporter)
	v.SetDefault("input", DefaultInput)
	v.SetDefault("output", DefaultOutput)

	// Environment variables
	v.SetEnvPrefix(EnvPrefix)
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Config file
	v.SetConfigName("timelinegenerator")
	v.SetConfigType("yml")

	// Search paths
	v.AddConfigPath(".")
	if configDir, err := os.UserConfigDir(); err == nil {
		v.AddConfigPath(filepath.Join(configDir, "timelinegenerator"))
	}

	// Read config
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("reading config: %w", err)
		}
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("parsing config: %w", err)
	}

	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("validating config: %w", err)
	}

	return &config, nil
}