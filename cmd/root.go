package cmd

import (
	"fmt"
	"github.com/zacharyelston/timelinegenerator/internal/config"
	"github.com/spf13/cobra"
	"os"
)

var cfg *config.Config

var rootCmd = &cobra.Command{
	Use:   "timelinegenerator",
	Short: "Generate timelines from YAML files",
	Long: `Timeline Generator CLI generates timelines from YAML files.
Multiple export formats are supported (TimelineJS, Vis.js, Mermaid, Bootstrap).
Data can be filtered by date, tags, and categories.`,
}

func Execute() {
	var err error
	cfg, err = config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}