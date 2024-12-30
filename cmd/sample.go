package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var sampleCmd = &cobra.Command{
	Use:   "sample [outputPath]",
	Short: "Generate a sample timeline file",
	Long: `Generate a sample timeline YAML file that demonstrates all available features.
The file includes examples of events with different properties like dates, tags,
categories, locations, images, and links.`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Default output path is sample.yml in current directory
		outputPath := "sample.yml"
		if len(args) > 0 {
			outputPath = args[0]
		}

		// Ensure directory exists
		dir := filepath.Dir(outputPath)
		if dir != "." {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return fmt.Errorf("creating directory: %w", err)
			}
		}

		// Write sample file
		if err := os.WriteFile(outputPath, []byte(sampleYAML), 0644); err != nil {
			return fmt.Errorf("writing sample file: %w", err)
		}

		fmt.Printf("Sample timeline created at: %s\n", outputPath)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(sampleCmd)
}

const sampleYAML = `version: 1.0
title: "Sample Timeline"
description: "This timeline demonstrates all available features of the timeline generator."

events:
  - title: "Project Kickoff"
    description: |
      # Project Launch
      Initial planning and team setup completed.
      
      ## Key Achievements
      - Team assembled
      - Project scope defined
      - Initial timeline created
    start: "2024-01-01"
    tags: ["milestone", "planning"]
    category: "Planning"
    location: "Main Office"
    
  - title: "Design Phase"
    description: "UI/UX design and architecture planning"
    start: "2024-01-15"
    end: "2024-02-15"
    tags: ["design", "architecture"]
    category: "Development"
    image: "https://example.com/design.jpg"
    
  - title: "First Sprint Complete"
    description: "Core features implemented"
    start: "2024-02-16"
    end: "2024-03-01"
    tags: ["sprint", "development"]
    category: "Development"
    link: "https://example.com/sprint1"
    
  - title: "Beta Release"
    description: |
      Beta version released for testing.
      - Feature A
      - Feature B
      - Feature C
    start: "2024-03-15"
    tags: ["release", "milestone"]
    category: "Release"
    location: "Virtual Event"
    color: "#4CAF50"
    icon: "🚀"
    
  - title: "User Testing"
    description: "Gathering user feedback"
    start: "2024-03-16"
    end: "2024-04-15"
    tags: ["testing", "feedback"]
    category: "Testing"
    
  - title: "Version 1.0 Launch"
    description: "Official product launch"
    start: "2024-05-01"
    tags: ["release", "milestone"]
    category: "Release"
    location: "Company HQ"
    link: "https://example.com/v1.0"
    color: "#2196F3"
    icon: "🎉"
`