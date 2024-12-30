package exporters

import (
	"fmt"
	"html/template"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/zacharyelston/timelinegenerator/internal/models"
)

type Bootstrap struct{}

func NewBootstrap() *Bootstrap {
	return &Bootstrap{}
}

func (e *Bootstrap) Export(timeline *models.Timeline, outputPath string) (string, error) {
	outputFile := outputPath + ".html"

	tmpl, err := template.New("bootstrap").Parse(bootstrapTemplate)
	if err != nil {
		return "", fmt.Errorf("parsing template: %w", err)
	}

	// Convert markdown to HTML for descriptions
	for i := range timeline.Events {
		if timeline.Events[i].Description != "" {
			html := markdown.ToHTML([]byte(timeline.Events[i].Description), nil, nil)
			timeline.Events[i].Description = string(html)
		}
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

const bootstrapTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>{{.Title}}</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet">
    <style>
        .timeline-card {
            margin-bottom: 2rem;
            border-left: 4px solid #0d6efd;
        }
        .timeline-date {
            color: #6c757d;
            font-size: 0.9rem;
        }
        .timeline-tag {
            font-size: 0.8rem;
            margin-right: 0.5rem;
        }
        .timeline-location {
            font-size: 0.9rem;
            color: #6c757d;
        }
    </style>
</head>
<body>
    <div class="container py-5">
        <header class="text-center mb-5">
            <h1>{{.Title}}</h1>
            <p class="lead">{{.Description}}</p>
        </header>

        <div class="timeline">
            {{range .Events}}
            <div class="card timeline-card">
                <div class="card-body">
                    <h5 class="card-title">{{.Title}}</h5>
                    <p class="timeline-date mb-2">
                        {{.Start.Format "January 2, 2006"}}
                        {{if .End}} - {{.End.Format "January 2, 2006"}}{{end}}
                    </p>
                    {{if .Location}}
                    <p class="timeline-location mb-2">
                        <i class="bi bi-geo-alt"></i> {{.Location}}
                    </p>
                    {{end}}
                    <div class="card-text mb-3">{{.Description}}</div>
                    {{if .Tags}}
                    <div class="mb-2">
                        {{range .Tags}}
                        <span class="badge bg-primary timeline-tag">#{{.}}</span>
                        {{end}}
                    </div>
                    {{end}}
                    {{if .Category}}
                    <span class="badge bg-secondary">{{.Category}}</span>
                    {{end}}
                    {{if .Link}}
                    <div class="mt-2">
                        <a href="{{.Link}}" class="btn btn-outline-primary btn-sm" target="_blank">Learn More</a>
                    </div>
                    {{end}}
                </div>
            </div>
            {{end}}
        </div>
    </div>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js"></script>
</body>
</html>`