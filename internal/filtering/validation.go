package filtering

import (
	"fmt"
	"time"
)

// ValidateFilter checks if the filter configuration is valid
func ValidateFilter(filter *Filter) error {
	if err := validateDates(filter); err != nil {
		return fmt.Errorf("invalid date filter: %w", err)
	}

	if err := validateTags(filter); err != nil {
		return fmt.Errorf("invalid tag filter: %w", err)
	}

	if err := validateCategories(filter); err != nil {
		return fmt.Errorf("invalid category filter: %w", err)
	}

	return nil
}

func validateDates(filter *Filter) error {
	// Check if end date is before start date
	if !filter.StartDate.IsZero() && !filter.EndDate.IsZero() {
		if filter.EndDate.Before(filter.StartDate) {
			return fmt.Errorf("end date (%s) cannot be before start date (%s)", 
				filter.EndDate.Format("2006-01-02"),
				filter.StartDate.Format("2006-01-02"))
		}
	}

	// Validate dates are not in the future
	now := time.Now()
	if !filter.StartDate.IsZero() && filter.StartDate.After(now) {
		return fmt.Errorf("start date (%s) cannot be in the future", 
			filter.StartDate.Format("2006-01-02"))
	}
	if !filter.EndDate.IsZero() && filter.EndDate.After(now) {
		return fmt.Errorf("end date (%s) cannot be in the future", 
			filter.EndDate.Format("2006-01-02"))
	}

	return nil
}

func validateTags(filter *Filter) error {
	// Check for empty tags
	for _, tag := range filter.Tags {
		if tag == "" {
			return fmt.Errorf("empty tag not allowed")
		}
	}
	for _, tag := range filter.ExcludeTags {
		if tag == "" {
			return fmt.Errorf("empty exclude tag not allowed")
		}
	}

	// Check for duplicate tags
	tagMap := make(map[string]bool)
	for _, tag := range filter.Tags {
		if tagMap[tag] {
			return fmt.Errorf("duplicate tag: %s", tag)
		}
		tagMap[tag] = true
	}

	// Check for conflicts between include and exclude tags
	for _, includeTag := range filter.Tags {
		for _, excludeTag := range filter.ExcludeTags {
			if includeTag == excludeTag {
				return fmt.Errorf("tag '%s' cannot be both included and excluded", includeTag)
			}
		}
	}

	return nil
}

func validateCategories(filter *Filter) error {
	// Check for empty categories
	for _, cat := range filter.Categories {
		if cat == "" {
			return fmt.Errorf("empty category not allowed")
		}
	}
	for _, cat := range filter.ExcludeCategories {
		if cat == "" {
			return fmt.Errorf("empty exclude category not allowed")
		}
	}

	// Check for duplicate categories
	catMap := make(map[string]bool)
	for _, cat := range filter.Categories {
		if catMap[cat] {
			return fmt.Errorf("duplicate category: %s", cat)
		}
		catMap[cat] = true
	}

	// Check for conflicts between include and exclude categories
	for _, includeCat := range filter.Categories {
		for _, excludeCat := range filter.ExcludeCategories {
			if includeCat == excludeCat {
				return fmt.Errorf("category '%s' cannot be both included and excluded", includeCat)
			}
		}
	}

	return nil
}