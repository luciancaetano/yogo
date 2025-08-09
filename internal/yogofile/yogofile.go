package yogofile

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

const currentVersion = "1.0.0"

type Endpoint struct {
	Name      string                    `yaml:"name"`
	Method    string                    `yaml:"method"`
	Path      string                    `yaml:"path"`
	Request   map[string]string         `yaml:"request,omitempty"`
	Responses map[int]map[string]string `yaml:"responses"`
}

type Generator struct {
	Name    string `yaml:"name"`
	Package string `yaml:"package"`
	Output  string `yaml:"output"`
}

type YOGOFile struct {
	Version    string      `yaml:"version"`
	Endpoints  []Endpoint  `yaml:"endpoints"`
	Generators []Generator `yaml:"generators,omitempty"`
}

func (y *YOGOFile) IsValid() bool {
	return y.Version == currentVersion && len(y.Endpoints) > 0
}

func (y *YOGOFile) GetEndpoint(name string) *Endpoint {
	for _, endpoint := range y.Endpoints {
		if endpoint.Name == name {
			return &endpoint
		}
	}
	return nil
}

func (y *YOGOFile) AddEndpoint(endpoint Endpoint) {
	y.Endpoints = append(y.Endpoints, endpoint)
}

func (y *YOGOFile) RemoveEndpoint(name string) {
	for i, endpoint := range y.Endpoints {
		if endpoint.Name == name {
			y.Endpoints = append(y.Endpoints[:i], y.Endpoints[i+1:]...)
			return
		}
	}
}

func (y *YOGOFile) UpdateEndpoint(name string, updated Endpoint) {
	for i, endpoint := range y.Endpoints {
		if endpoint.Name == name {
			y.Endpoints[i] = updated
			return
		}
	}
}

func (y *YOGOFile) Unmarshal(data []byte) error {
	return yaml.Unmarshal(data, y)
}

func (y *YOGOFile) Marshal() ([]byte, error) {
	return yaml.Marshal(y)
}

func New() *YOGOFile {
	return &YOGOFile{
		Version:   currentVersion,
		Endpoints: []Endpoint{},
	}
}

func (y *YOGOFile) Validate() error {
	if y.Version != currentVersion {
		return fmt.Errorf("invalid version: %s", y.Version)
	}
	if len(y.Endpoints) == 0 {
		return fmt.Errorf("no endpoints defined")
	}

	if len(y.Generators) == 0 {
		return fmt.Errorf("at least one generator must be defined")
	}

	return nil
}

func (y *YOGOFile) ContainsGenerator(name string) bool {
	for _, generator := range y.Generators {
		if generator.Name == name {
			return true
		}
	}
	return false
}

func (y *YOGOFile) GetGenerator(name string) *Generator {
	for _, generator := range y.Generators {
		if generator.Name == name {
			return &generator
		}
	}
	return nil
}

func (y *YOGOFile) CountGenerator() int {
	return len(y.Generators)
}

func (y *YOGOFile) DebugLog() {
	fmt.Printf("YOGO File Version: %s\n", y.Version)
	fmt.Printf("Number of Endpoints: %d\n", len(y.Endpoints))
	for _, ep := range y.Endpoints {
		fmt.Printf("  - Endpoint Name: %s, Method: %s, Path: %s\n", ep.Name, ep.Method, ep.Path)
	}
	fmt.Printf("Number of Generators: %d\n", len(y.Generators))
	for _, gen := range y.Generators {
		fmt.Printf("  - Generator Name: %s, Package: %s, Output: %s\n", gen.Name, gen.Package, gen.Output)
	}
}
