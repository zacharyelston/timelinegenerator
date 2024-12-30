package filtering

import (
	"fmt"
	"github.com/zacharyelston/timelinegenerator/internal/models"
	"time"
)

// Filter represents a timeline filter
type Filter struct {
	StartDate         time.Time
	EndDate           time.Time
	Tags              []string
	ExcludeTags       []string
	Categories        []string
	ExcludeCategories []string
}

// ApplyFilters filters events based on the given criteria
func ApplyFilters(timeline *models.Timeline, filter *Filter) (*models.Timeline, error) {
	// Validate filter configuration
	if err := ValidateFilter(filter); err != nil {
		return nil, fmt.Errorf("invalid filter configuration: %w", err)
	}

	filtered := &models.Timeline{
		Version:     timeline.Version,
		Title:       timeline.Title,
		Description: timeline.Description,
		Events:      make([]models.Event, 0),
	}

	for _, event := range timeline.Events {
		matches, err := eventMatchesFilters(event, filter)
		if err != nil {
			return nil, fmt.Errorf("error filtering event '%s': %w", event.Title, err)
		}

		if matches {
			filtered.Events = append(filtered.Events, event)
		}
	}

	// Validate we have at least one event after filtering
	if len(filtered.Events) == 0 {
		return nil, fmt.Errorf("no events match the specified filters")
	}

	return filtered, nil
}

func eventMatchesFilters(event models.Event, filter *Filter) (bool, error) {
	// Date filtering
	if !matchesDateFilter(event, filter) {
		return false, nil
	}

	// Tag filtering
	matches, err := matchesTagFilter(event, filter)
	if err != nil {
		return false, fmt.Errorf("tag filter error: %w", err)
	}
	if !matches {
		return false, nil
	}

	// Category filtering
	matches, err = matchesCategoryFilter(event, filter)
	if err != nil {
		return false, fmt.Errorf("category filter error: %w", err)
	}
	if !matches {
		return false, nil
	}

	return true, nil
}

func matchesDateFilter(event models.Event, filter *Filter) bool {
	if !filter.StartDate.IsZero() && event.Start.Before(filter.StartDate) {
		return false
	}

	if !filter.EndDate.IsZero() {
		if event.End != nil {
			if event.End.After(filter.EndDate) {
				return false
			}
		} else if event.Start.After(filter.EndDate) {
			return false
		}
	}

	return true
}

func matchesTagFilter(event models.Event, filter *Filter) (bool, error) {
	// Check excluded tags first
	for _, excludeTag := range filter.ExcludeTags {
		for _, tag := range event.Tags {
			if tag == excludeTag {
				return false, nil
			}
		}
	}

	// If no include tags specified, all non-excluded tags match
	if len(filter.Tags) == 0 {
		return true, nil
	}

	// Check for at least one matching tag
	for _, includeTag := range filter.Tags {
		for _, tag := range event.Tags {
			if tag == includeTag {
				return true, nil
			}
		}
	}

	return false, nil
}

func matchesCategoryFilter(event models.Event, filter *Filter) (bool, error) {
	// Check excluded categories
	for _, excludeCat := range filter.ExcludeCategories {
		if event.Category == excludeCat {
			return false, nil
		}
	}

	// If no include categories specified, all non-excluded categories match
	if len(filter.Categories) == 0 {
		return true, nil
	}

	// Check for matching category
	for _, includeCat := range filter.Categories {
		if event.Category == includeCat {
			return true, nil
		}
	}

	return false, nil
}