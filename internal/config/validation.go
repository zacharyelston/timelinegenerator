package config

import (
	"fmt"
	"time"
)

func (c *Config) Validate() error {
	if err := validateExporter(c.Exporter); err != nil {
		return fmt.Errorf("invalid exporter: %w", err)
	}

	if err := validateFilters(c.Filters); err != nil {
		return fmt.Errorf("invalid filters: %w", err)
	}

	return nil
}

func validateExporter(exporter string) error {
	valid := map[string]bool{
		"timelinejs": true,
		"visjs":      true,
		"mermaid":    true,
		"bootstrap":  true,
	}

	if !valid[exporter] {
		return fmt.Errorf("unsupported exporter: %s", exporter)
	}
	return nil
}

func validateFilters(f Filters) error {
	if f.StartDate != "" && f.EndDate != "" {
		start, err := time.Parse("2006-01-02", f.StartDate)
		if err != nil {
			return fmt.Errorf("invalid start date: %w", err)
		}

		end, err := time.Parse("2006-01-02", f.EndDate)
		if err != nil {
			return fmt.Errorf("invalid end date: %w", err)
		}

		if end.Before(start) {
			return fmt.Errorf("end date cannot be before start date")
		}
	}

	return nil
}