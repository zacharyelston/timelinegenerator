package filtering

import (
	"testing"
	"time"

	"github.com/zacharyelston/timelinegenerator/internal/models"
)

func TestApplyFilters(t *testing.T) {
	timeline := &models.Timeline{
		Events: []models.Event{
			{
				Title:    "Event 1",
				Start:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				Tags:     []string{"tag1", "tag2"},
				Category: "cat1",
			},
			{
				Title:    "Event 2",
				Start:    time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
				Tags:     []string{"tag2", "tag3"},
				Category: "cat2",
			},
		},
	}

	tests := []struct {
		name          string
		filter        *Filter
		wantEventLen  int
		wantErr       bool
		errorContains string
	}{
		{
			name: "no filters returns all events",
			filter: &Filter{},
			wantEventLen: 2,
		},
		{
			name: "filter by tag",
			filter: &Filter{
				Tags: []string{"tag1"},
			},
			wantEventLen: 1,
		},
		{
			name: "filter by category",
			filter: &Filter{
				Categories: []string{"cat1"},
			},
			wantEventLen: 1,
		},
		{
			name: "filter by date range",
			filter: &Filter{
				StartDate: time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2024, 2, 15, 0, 0, 0, 0, time.UTC),
			},
			wantEventLen: 1,
		},
		{
			name: "invalid date range",
			filter: &Filter{
				StartDate: time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr:       true,
			errorContains: "end date cannot be before start date",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filtered, err := ApplyFilters(timeline, tt.filter)
			
			if tt.wantErr {
				if err == nil {
					t.Error("expected error but got none")
				} else if !contains(err.Error(), tt.errorContains) {
					t.Errorf("error = %v, want error containing %v", err, tt.errorContains)
				}
				return
			}
			
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if len(filtered.Events) != tt.wantEventLen {
				t.Errorf("got %d events, want %d", len(filtered.Events), tt.wantEventLen)
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[0:len(substr)] == substr
}