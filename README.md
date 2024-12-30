# Timeline Generator CLI

Generate beautiful timelines from YAML files with multiple export formats and filtering capabilities.

![CI Status](https://github.com/zacharyelston/timelinegenerator/workflows/CI/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/zacharyelston/timelinegenerator)](https://goreportcard.com/report/github.com/zacharyelston/timelinegenerator)
[![GoDoc](https://godoc.org/github.com/zacharyelston/timelinegenerator?status.svg)](https://godoc.org/github.com/zacharyelston/timelinegenerator)

## Features

- Multiple export formats:
  - TimelineJS - Interactive timeline with media support
  - Vis.js - Dynamic timeline with grouping and zooming
  - Mermaid - Simple markdown-based timeline diagrams
  - Bootstrap - Responsive card-based layout
- Filter events by:
  - Date range
  - Tags
  - Categories
- Markdown support in descriptions
- YAML-based data format
- Cross-platform support (Windows, macOS, Linux)

## Installation

### Using Go

```bash
go install github.com/zacharyelston/timelinegenerator@latest
```

### From Releases

Download the latest release for your platform from the [releases page](https://github.com/zacharyelston/timelinegenerator/releases).

```bash
# Linux/macOS
curl -LO https://github.com/zacharyelston/timelinegenerator/releases/latest/download/timelinegenerator_$(uname -s)_$(uname -m).tar.gz
tar xzf timelinegenerator_*.tar.gz
sudo mv timelinegenerator /usr/local/bin/

# Windows (PowerShell)
Invoke-WebRequest -Uri "https://github.com/zacharyelston/timelinegenerator/releases/latest/download/timelinegenerator_Windows_x86_64.zip" -OutFile "timelinegenerator.zip"
Expand-Archive timelinegenerator.zip -DestinationPath .
```

## Quick Start

1. Generate a sample timeline:
```bash
timelinegenerator sample timeline.yml
```

2. Generate HTML output:
```bash
timelinegenerator generate timeline.yml output
```

## Usage

### Basic Usage

```bash
# Generate with default exporter (TimelineJS)
timelinegenerator generate input.yml output

# Specify different exporter
timelinegenerator generate -e visjs input.yml output
timelinegenerator generate -e mermaid input.yml output
timelinegenerator generate -e bootstrap input.yml output
```

### Filtering Examples

```bash
# Filter by date range
timelinegenerator generate --start 2024-01-01 --end 2024-12-31 input.yml output

# Filter by tags
timelinegenerator generate --tag work --tag important input.yml output
timelinegenerator generate --exclude-tag personal input.yml output

# Filter by categories
timelinegenerator generate --category "Project A" input.yml output
timelinegenerator generate --exclude-category "Archive" input.yml output

# Combine filters
timelinegenerator generate \
  --start 2024-01-01 \
  --tag important \
  --category "Project A" \
  input.yml output
```

## Configuration

### Configuration File

Create `timelinegenerator.yml` in your project directory or `~/.config/timelinegenerator/`:

```yaml
# Default settings
exporter: timelinejs
input: timeline.yml
output: timeline

# Default filters
filters:
  start_date: "2024-01-01"
  end_date: "2024-12-31"
  tags: 
    - important
    - milestone
  exclude_tags:
    - draft
  categories:
    - Development
    - Release
  exclude_categories:
    - Archive
```

### Environment Variables

All settings can be configured using environment variables:

```bash
# Set defaults
export TIMELINE_EXPORTER=mermaid
export TIMELINE_INPUT=custom.yml
export TIMELINE_OUTPUT=output

# Set filters
export TIMELINE_FILTERS_START_DATE=2024-01-01
export TIMELINE_FILTERS_TAGS=work,important
```

## YAML Format

```yaml
version: 1.0
title: "Project Timeline"
description: "Timeline of project milestones"

events:
  - title: "Project Start"
    description: |
      # Project Kickoff
      Initial planning and team setup.
    start: "2024-01-01"
    tags: ["milestone", "planning"]
    category: "Planning"
    
  - title: "Phase 1 Complete"
    description: "First milestone achieved"
    start: "2024-03-15"
    end: "2024-03-20"
    tags: ["milestone"]
    category: "Development"
    location: "HQ"
    image: "phase1.jpg"
    link: "https://docs.example.com/phase1"
```

## Export Formats

### TimelineJS
- Interactive timeline with media support
- Supports date ranges and categories
- Mobile-friendly and responsive
- [Example](https://timeline.knightlab.com/)

### Vis.js
- Dynamic timeline with zoom controls
- Supports grouping by category
- Interactive events
- [Example](https://visjs.github.io/vis-timeline/examples/timeline/)

### Mermaid
- Simple markdown-based timeline
- Great for documentation
- Supports sections by category
- [Example](https://mermaid.js.org/syntax/timeline.html)

### Bootstrap
- Card-based layout
- Fully responsive
- Markdown support
- [Example](https://getbootstrap.com/docs/5.3/components/card/)

## Development

### Prerequisites

- Go 1.22 or later
- Make (optional, for using Makefile commands)

### Building from Source

```bash
# Clone repository
git clone https://github.com/zacharyelston/timelinegenerator.git
cd timelinegenerator

# Build
make build

# Run tests
make test

# Run linter
make lint
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License - See [LICENSE](LICENSE) for details.