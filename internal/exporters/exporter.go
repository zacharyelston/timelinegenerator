package exporters

type Exporter interface {
	Export(timeline *models.Timeline, outputPath string) (string, error)
}