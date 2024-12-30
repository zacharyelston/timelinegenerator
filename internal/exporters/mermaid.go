package exporters

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/zacharyelston/timelinegenerator/internal/models"
)

type Mermaid struct{}

func NewMermaid() *Mermaid {
	return &Mermaid{}
}

func (e *Mermaid) Export(timeline *models.Timeline, outputPath string) (string, error) {
	outputFile := outputPath + ".md"

	tmpl, err := template.New("mermaid").Parse(mermaidTemplate)
	if err != nil {
		return "", fmt.Errorf("parsing template: %w", err)
	}

	// Group events by category
	categories := make(map[string][]models.Event)
	for _, event := range timeline.Events {
		cat := event.Category
		if cat == "" {
			cat = "Default"
		}
		categories[cat] = append(categories[cat], event)
	}

	data := struct {
		*models.Timeline
		Categories map[string][]models.Event
	}{
		Timeline:   timeline,
		Categories: categories,
	}

	f, err := os.Create(outputFile)
	if err != nil {
		return "", fmt.Errorf("creating output file: %w", err)
	}
	defer f.Close()

	if err := tmpl.Execute(f, data); err != nil {
		return "", fmt.Errorf("executing template: %w", err)
	}

	return outputFile, nil
}

const mermaidTemplate = `# {{.Title}}

{{.Description}}

` + "```mermaid" + `
timeline
    title {{.Title}}
    {{range $category, $events := .Categories}}
    section {{$category}}
        {{range $events}}
        {{.Start.Format "2006-01-02"}} : {{.Title}}
        {{end}}
    {{end}}
` + "```" + `

Generated with TimelineGenerator
`