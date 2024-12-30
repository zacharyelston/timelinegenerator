package exporters

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/zacharyelston/timelinegenerator/internal/models"
)

type TimelineJS struct{}

func NewTimelineJS() *TimelineJS {
	return &TimelineJS{}
}

func (e *TimelineJS) Export(timeline *models.Timeline, outputPath string) (string, error) {
	outputFile := outputPath + ".html"
	
	tmpl, err := template.New("timelinejs").Parse(timelineJSTemplate)
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

const timelineJSTemplate = `<!DOCTYPE html>
<html lang="en">
  <head>
    <title>{{.Title}}</title>
    <meta name="description" content="{{.Description}}">
    <meta charset="utf-8">
    <style>
      html, body {
        height:100%;
        padding: 0px;
        margin: 0px;
      }
    </style>
  </head>
  <body>
    <div id="timeline-embed">
      <div id="timeline"></div>
    </div>

    <link title="timeline-styles" rel="stylesheet" href="https://cdn.knightlab.com/libs/timeline3/latest/css/timeline.css">
    <script src="https://cdn.knightlab.com/libs/timeline3/latest/js/timeline-min.js"></script>

    <script>
    var timelineJson = {
      title: { 
        text: { 
          headline: "{{.Title}}", 
          text: "{{.Description}}" 
        } 
      },
      events: [
        {{range .Events}}
        {
          start_date: { 
            year: {{.Start.Year}}, 
            month: {{.Start.Month}}, 
            day: {{.Start.Day}} 
          },
          {{if .End}}
          end_date: { 
            year: {{.End.Year}}, 
            month: {{.End.Month}}, 
            day: {{.End.Day}} 
          },
          {{end}}
          text: { 
            headline: "{{.Title}}", 
            text: "{{.Description}}" 
          },
          {{if .Image}}
          media: { 
            url: "{{.Image}}", 
            caption: "{{.Title}}" 
          },
          {{else if .Link}}
          media: { 
            url: "{{.Link}}", 
            text: "Read more" 
          },
          {{end}}
          {{if .Category}}
          group: "{{.Category}}",
          {{end}}
        },
        {{end}}
      ]
    };

    window.addEventListener('DOMContentLoaded', function() {
      var embed = document.getElementById('timeline-embed');
      embed.style.height = getComputedStyle(document.body).height;
      window.timeline = new TL.Timeline('timeline-embed', timelineJson, {
        hash_bookmark: false
      });
      window.addEventListener('resize', function() {
        embed.style.height = getComputedStyle(document.body).height;
        timeline.updateDisplay();
      });
    });
    </script>
  </body>
</html>`