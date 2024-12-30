package exporters

import "fmt"

func NewExporter(name string) (Exporter, error) {
	switch name {
	case "timelinejs":
		return NewTimelineJS(), nil
	case "visjs":
		return NewVisJS(), nil
	case "mermaid":
		return NewMermaid(), nil
	case "bootstrap":
		return NewBootstrap(), nil
	default:
		return nil, fmt.Errorf("unknown exporter: %s", name)
	}
}