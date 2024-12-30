package exporters

import (
	"fmt"
	"html/template"
	"os"

	"github.com/zacharyelston/timelinegenerator/internal/models"
)

type VisJS struct{}

func NewVisJS() *VisJS {
	return &VisJS{}
}

func (e *VisJS) Export(timeline *models.Timeline, outputPath string) (string, error) {
	outputFile := outputPath + ".html"

	tmpl, err := template.New("visjs").Parse(visJSTemplate)
	if err != nil {
		return "", fmt.Errorf("parsing template: %w", err)
	}

	f, err := os.Create(outputFile)
	if err != nil {
		return "", fmt.Errorf("creating output file: %w", err)
	}
	defer f.Close()

	if err := tmpl.Execute(f, timeline); err != nil {
		return "", fmt.Errorf("executing template: %w", err)
	}

	return outputFile, nil
}

const visJSTemplate = `<!DOCTYPE html>
<html lang="en">
  <head>
    <title>{{.Title}}</title>
    <script type="text/javascript" src="https://unpkg.com/vis-timeline@latest/standalone/umd/vis-timeline-graph2d.min.js"></script>
    <link href="https://unpkg.com/vis-timeline@latest/styles/vis-timeline-graph2d.min.css" rel="stylesheet" type="text/css" />
    <style type="text/css">
      #visualization {
        border: 1px solid lightgray;
        padding: 1em;
        margin: 1em;
      }
      .vis-item {
        border-color: #3498db;
        background-color: #ffffff;
        padding: 0.5em;
      }
      .vis-item.vis-selected {
        border-color: #2980b9;
        background-color: #ecf0f1;
      }
      .vis-time-axis .vis-text {
        color: #34495e;
      }
    </style>
  </head>
  <body>
    <h1>{{.Title}}</h1>
    <p>{{.Description}}</p>
    <div id="visualization"></div>
    <script type="text/javascript">
      // Create groups for categories
      var groups = [
        {{$first := true}}
        {{range $category := .Events | uniqueCategories}}
          {{if $first}}{{$first = false}}{{else}},{{end}}
          {id: {{if $category}}'{{$category}}'{{else}}'default'{{end}}, 
           content: {{if $category}}'{{$category}}'{{else}}'Default'{{end}}}
        {{end}}
      ];

      // Create timeline items
      var items = [
        {{$first := true}}
        {{range .Events}}
          {{if $first}}{{$first = false}}{{else}},{{end}}
          {
            id: '{{.Title}}',
            content: '{{.Title}}',
            start: '{{.Start.Format "2006-01-02"}}',
            {{if .End}}
            end: '{{.End.Format "2006-01-02"}}',
            {{end}}
            group: {{if .Category}}'{{.Category}}'{{else}}'default'{{end}},
            title: '{{.Description}}'
          }
        {{end}}
      ];

      // Create timeline
      var container = document.getElementById('visualization');
      var options = {
        height: '600px',
        stack: true,
        showMajorLabels: true,
        showCurrentTime: false,
        zoomMin: 1000 * 60 * 60 * 24 * 7, // One week
        zoomMax: 1000 * 60 * 60 * 24 * 365 * 10 // Ten years
      };

      var timeline = new vis.Timeline(container, items, groups, options);
    </script>
  </body>
</html>`