package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateCmd(t *testing.T) {
	// Create temporary directory for test files
	tempDir := t.TempDir()
	
	// Create sample timeline file
	samplePath := filepath.Join(tempDir, "sample.yml")
	err := os.WriteFile(samplePath, []byte(sampleYAML), 0644)
	if err != nil {
		t.Fatalf("Failed to create sample file: %v", err)
	}

	tests := []struct {
		name        string
		args        []string
		wantErr     bool
		outputCheck func(t *testing.T, outputPath string)
	}{
		{
			name: "basic generation",
			args: []string{samplePath, filepath.Join(tempDir, "output")},
			outputCheck: func(t *testing.T, outputPath string) {
				if _, err := os.Stat(outputPath + ".html"); os.IsNotExist(err) {
					t.Error("Output file not created")
				}
			},
		},
		{
			name:    "missing input file",
			args:    []string{"nonexistent.yml", "output"},
			wantErr: true,
		},
		{
			name: "with date filter",
			args: []string{
				"--start", "2024-01-01",
				"--end", "2024-12-31",
				samplePath,
				filepath.Join(tempDir, "output-filtered"),
			},
			outputCheck: func(t *testing.T, outputPath string) {
				if _, err := os.Stat(outputPath + ".html"); os.IsNotExist(err) {
					t.Error("Output file not created")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := generateCmd
			cmd.SetArgs(tt.args)
			
			err := cmd.Execute()
			
			if tt.wantErr {
				if err == nil {
					t.Error("Expected error but got none")
				}
				return
			}
			
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if tt.outputCheck != nil {
				tt.outputCheck(t, tt.args[len(tt.args)-1])
			}
		})
	}
}