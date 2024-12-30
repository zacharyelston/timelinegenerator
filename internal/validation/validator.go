package validation

import (
	"fmt"
	"github.com/zacharyelston/timelinegenerator/internal/models"
)

// ValidateTimeline performs validation on the timeline and its events
func ValidateTimeline(timeline *models.Timeline) error {
	if timeline == nil {
		return fmt.Errorf("timeline is nil")
	}

	if timeline.Title == "" {
		return fmt.Errorf("timeline title is required")
	}

	if len(timeline.Events) == 0 {
		return fmt.Errorf("timeline must contain at least one event")
	}

	for i, event := range timeline.Events {
		if err := ValidateEvent(event); err != nil {
			return fmt.Errorf("event %d (%s): %w", i+1, event.Title, err)
		}
	}

	return nil
}

// ValidateEvent performs validation on a single event
func ValidateEvent(event models.Event) error {
	if event.Title == "" {
		return fmt.Errorf("event title is required")
	}

	if event.Start.IsZero() {
		return fmt.Errorf("event start date is required")
	}

	if event.End != nil && event.End.Before(event.Start) {
		return fmt.Errorf("event end date cannot be before start date")
	}

	return nil
}