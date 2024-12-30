package models

import "time"

type Event struct {
	Title       string     `yaml:"title"`
	Description string     `yaml:"description"`
	Start       time.Time  `yaml:"start"`
	End         *time.Time `yaml:"end,omitempty"`
	Tags        []string   `yaml:"tags,omitempty"`
	Category    string     `yaml:"category,omitempty"`
	Location    string     `yaml:"location,omitempty"`
	Image       string     `yaml:"image,omitempty"`
	Link        string     `yaml:"link,omitempty"`
	Color       string     `yaml:"color,omitempty"`
	Icon        string     `yaml:"icon,omitempty"`
}

type Timeline struct {
	Version     string  `yaml:"version"`
	Title       string  `yaml:"title"`
	Description string  `yaml:"description"`
	Events      []Event `yaml:"events"`
}
