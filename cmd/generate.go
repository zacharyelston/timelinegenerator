package cmd

import (
	"fmt"
	"github.com/zacharyelston/timelinegenerator/internal/exporters"
	"github.com/zacharyelston/timelinegenerator/internal/filtering"
	"github.com/zacharyelston/timelinegenerator/pkg/yaml"
	"github.com/spf13/cobra"
	"time"
)

var generateCmd = &cobra.Command{
	Use:   "generate [inputPath] [outputPath]",
	Short: "Generate a timeline using an exporter",
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get input/output paths
		inputPath := args[0]
		outputPath := "timeline"
		if len(args) > 1 {
			outputPath = args[1]
		}

		// Load timeline
		timeline, err := yaml.Import(inputPath)
		if err != nil {
			return fmt.Errorf("loading timeline: %w", err)
		}

		// Create filter
		filter := &filtering.Filter{}

		// Parse date filters
		if startStr, _ := cmd.Flags().GetString("start"); startStr != "" {
			start, err := time.Parse("2006-01-02", startStr)
			if err != nil {
				return fmt.Errorf("parsing start date: %w", err)
			}
			filter.StartDate = start
		}

		if endStr, _ := cmd.Flags().GetString("end"); endStr != "" {
			end, err := time.Parse("2006-01-02", endStr)
			if err != nil {
				return fmt.Errorf("parsing end date: %w", err)
			}
			filter.EndDate = end
		}

		// Get other filters
		filter.Tags, _ = cmd.Flags().GetStringSlice("tag")
		filter.ExcludeTags, _ = cmd.Flags().GetStringSlice("exclude-tag")
		filter.Categories, _ = cmd.Flags().GetStringSlice("category")
		filter.ExcludeCategories, _ = cmd.Flags().GetStringSlice("exclude-category")

		// Apply filters
		filtered := filtering.ApplyFilters(timeline, filter)

		// Create exporter
		exporterName, _ := cmd.Flags().GetString("exporter")
		exporter, err := exporters.NewExporter(exporterName)
		if err != nil {
			return fmt.Errorf("creating exporter: %w", err)
		}

		// Export timeline
		outputFile, err := exporter.Export(filtered, outputPath)
		if err != nil {
			return fmt.Errorf("exporting timeline: %w", err)
		}

		fmt.Printf("Timeline generated at: %s\n", outputFile)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	
	generateCmd.Flags().StringP("exporter", "e", "timelinejs", "Exporter to use (timelinejs, visjs, mermaid, bootstrap)")
	generateCmd.Flags().String("start", "", "Filter events starting on or after this date (YYYY-MM-DD)")
	generateCmd.Flags().String("end", "", "Filter events ending on or before this date (YYYY-MM-DD)")
	generateCmd.Flags().StringSlice("tag", nil, "Filter events with these tags")
	generateCmd.Flags().StringSlice("exclude-tag", nil, "Exclude events with these tags")
	generateCmd.Flags().StringSlice("category", nil, "Filter events with these categories")
	generateCmd.Flags().StringSlice("exclude-category", nil, "Exclude events with these categories")
}