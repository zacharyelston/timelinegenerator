package exporters

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/zacharyelston/timelinegenerator/internal/models"
)

func TestVisJSExporter(t *testing.T) {
	timeline := &models.Timeline{
		Title:       "Test Timeline",
		Description: "Test Description",
		Events: []models.Event{
			{
				Title:       "Event 1",
				Description: "Event Description",
				Start:       time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				Category:    "Category 1",
				Tags:        []string{"tag1"},
			},
			{
				Title:       "Event 2",
				Description: "Event Description",
				Start:       time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
				End:         ptr(time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)),
				Category:    "Category 2",
			},
		},
	}

	tests := []struct {
		name            string
		timeline        *models.Timeline
		wantElements    []string
		wantErr         bool
		errorContains   string
	}{
		{
			name:     "basic timeline",
			timeline: timeline,
			wantElements: []string{
				// Required HTML structure
				"<!DOCTYPE html>",
				"<html lang=\"en\">",
				"<div id=\"visualization\">",
				
				// Timeline data
				"Test Timeline",
				"Test Description",
				"Event 1",
				"Event 2",
				"Category 1",
				"Category 2",
				
				// VisJS specific elements
				"vis-timeline",
				"new vis.Timeline",
				"2024-01-01",
				"2024-02-01",
			},
		},
		{
			name: "empty timeline",
			timeline: &models.Timeline{
				Events: []models.Event{},
			},
			wantErr:       true,
			errorContains: "timeline must contain events",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exporter := NewVisJS()
			tempDir := t.TempDir()
			outputPath := filepath.Join(tempDir, "test")

			outputFile, err := exporter.Export(tt.timeline, outputPath)

			if tt.wantErr {
				if err == nil {
					t.Error("expected error but got none")
				} else if !strings.Contains(err.Error(), tt.errorContains) {
					t.Errorf("error = %v, want error containing %v", err, tt.errorContains)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			content, err := os.ReadFile(outputFile)
			if err != nil {
				t.Fatalf("failed to read output file: %v", err)
			}

			htmlContent := string(content)
			for _, element := range tt.wantElements {
				if !strings.Contains(htmlContent, element) {
					t.Errorf("output missing required element: %s", element)
				}
			}
		})
	}
}

// Helper function to create pointer to time.Time
func ptr(t time.Time) *time.Time {
	return &t
}