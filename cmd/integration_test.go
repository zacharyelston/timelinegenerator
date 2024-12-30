package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCLIIntegration(t *testing.T) {
	// Setup test environment
	tempDir := t.TempDir()
	samplePath := filepath.Join(tempDir, "sample.yml")
	err := os.WriteFile(samplePath, []byte(sampleYAML), 0644)
	if err != nil {
		t.Fatalf("Failed to create sample file: %v", err)
	}

	tests := []struct {
		name           string
		args           []string
		env           map[string]string
		wantFiles     []string
		wantContent   []string
		wantErr       bool
		errorContains string
	}{
		{
			name: "sample command",
			args: []string{"sample", filepath.Join(tempDir, "output.yml")},
			wantFiles: []string{"output.yml"},
			wantContent: []string{"version: 1.0", "title:", "events:"},
		},
		{
			name: "generate with timelinejs",
			args: []string{"generate", "-e", "timelinejs", samplePath, filepath.Join(tempDir, "timeline-js")},
			wantFiles: []string{"timeline-js.html"},
			wantContent: []string{"TimelineJS", "timeline-embed"},
		},
		{
			name: "generate with visjs",
			args: []string{"generate", "-e", "visjs", samplePath, filepath.Join(tempDir, "timeline-vis")},
			wantFiles: []string{"timeline-vis.html"},
			wantContent: []string{"vis-timeline", "visualization"},
		},
		{
			name: "generate with mermaid",
			args: []string{"generate", "-e", "mermaid", samplePath, filepath.Join(tempDir, "timeline-mermaid")},
			wantFiles: []string{"timeline-mermaid.md"},
			wantContent: []string{"```mermaid", "timeline"},
		},
		{
			name: "generate with bootstrap",
			args: []string{"generate", "-e", "bootstrap", samplePath, filepath.Join(tempDir, "timeline-bootstrap")},
			wantFiles: []string{"timeline-bootstrap.html"},
			wantContent: []string{"bootstrap", "card"},
		},
		{
			name: "filter by date range",
			args: []string{
				"generate",
				"--start", "2024-01-01",
				"--end", "2024-12-31",
				samplePath,
				filepath.Join(tempDir, "filtered"),
			},
			wantFiles: []string{"filtered.html"},
		},
		{
			name: "filter by tags",
			args: []string{
				"generate",
				"--tag", "tag1",
				"--exclude-tag", "tag2",
				samplePath,
				filepath.Join(tempDir, "tagged"),
			},
			wantFiles: []string{"tagged.html"},
		},
		{
			name: "invalid exporter",
			args: []string{"generate", "-e", "invalid", samplePath, "output"},
			wantErr: true,
			errorContains: "unknown exporter",
		},
		{
			name: "invalid date range",
			args: []string{
				"generate",
				"--start", "2024-12-31",
				"--end", "2024-01-01",
				samplePath,
				"output",
			},
			wantErr: true,
			errorContains: "end date cannot be before start date",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variables
			for k, v := range tt.env {
				os.Setenv(k, v)
				defer os.Unsetenv(k)
			}

			// Execute command
			rootCmd.SetArgs(tt.args)
			err := rootCmd.Execute()

			// Check error cases
			if tt.wantErr {
				if err == nil {
					t.Error("Expected error but got none")
				} else if !strings.Contains(err.Error(), tt.errorContains) {
					t.Errorf("Error = %v, want error containing %v", err, tt.errorContains)
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			// Check output files
			for _, file := range tt.wantFiles {
				path := filepath.Join(tempDir, file)
				if _, err := os.Stat(path); os.IsNotExist(err) {
					t.Errorf("Expected file not created: %s", file)
					continue
				}

				// Check file contents if specified
				if len(tt.wantContent) > 0 {
					content, err := os.ReadFile(path)
					if err != nil {
						t.Errorf("Failed to read file %s: %v", file, err)
						continue
					}

					for _, want := range tt.wantContent {
						if !strings.Contains(string(content), want) {
							t.Errorf("File %s missing content: %s", file, want)
						}
					}
				}
			}
		})
	}
}