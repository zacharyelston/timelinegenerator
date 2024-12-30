package yaml

import (
	"fmt"
	"os"

	"github.com/zacharyelston/timelinegenerator/internal/models"
	"github.com/zacharyelston/timelinegenerator/internal/validation"
	"gopkg.in/yaml.v3"
)

func Import(path string) (*models.Timeline, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	var timeline models.Timeline
	if err := yaml.Unmarshal(data, &timeline); err != nil {
		return nil, fmt.Errorf("parsing YAML: %w", err)
	}

	if err := validation.ValidateTimeline(&timeline); err != nil {
		return nil, fmt.Errorf("validating timeline: %w", err)
	}

	return &timeline, nil
}