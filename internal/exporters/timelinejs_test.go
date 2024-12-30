package exporters

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/zacharyelston/timelinegenerator/internal/models"
)

func TestTimelineJSExporter(t *testing.T) {
	timeline := &models.Timeline{
		Title:       "Test Timeline",
		Description: "Test Description",
		Events: []models.Event{
			{
				Title:       "Test Event",
				Description: "Test Event Description",
				Start:       time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				Category:    "Test Category",
				Tags:        []string{"test"},
			},
		},
	}

	exporter := NewTimelineJS()
	tempDir := t.TempDir()
	outputPath := filepath.Join(tempDir, "test")

	outputFile, err := exporter.Export(timeline, outputPath)
	if err != nil {
		t.Fatalf("Export failed: %v", err)
	}

	content, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	// Check for required elements
	requiredElements := []string{
		timeline.Title,
		timeline.Description,
		timeline.Events[0].Title,
		"2024",
		"timeline-embed",
	}

	for _, element := range requiredElements {
		if !strings.Contains(string(content), element) {
			t.Errorf("Output missing required element: %s", element)
		}
	}
}