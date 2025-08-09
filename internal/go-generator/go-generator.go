package gogenerator

import (
	"bytes"
	"fmt"
	"github/luciancaetano/yogo/internal/yogofile"
	"regexp"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type endpointData struct {
	Name       string
	Method     string
	Path       string
	FmtPath    string
	PathParams []string
	Request    map[string]string
	Responses  map[int]map[string]string
}

type templateData struct {
	PackageName string
	Endpoints   []endpointData
}

var pathParamRegex = regexp.MustCompile(`\{(\w+)\}`)

func Generate(y *yogofile.YOGOFile) (string, error) {
	if y == nil || !y.IsValid() {
		return "", fmt.Errorf("invalid YOGO file")
	}

	goGeneratorConfig := y.GetGenerator("go")

	if goGeneratorConfig == nil {
		return "", fmt.Errorf("no Go generator defined in YOGO file")
	}

	output := goGeneratorConfig.Output
	if !strings.HasSuffix(output, ".go") {
		return "", fmt.Errorf("output file must have a .go extension")
	}

	funcs := template.FuncMap{
		"title": func(s string) string {
			return cases.Title(language.Und).String(s)
		},
	}

	mainTpl, err := templates.ReadFile("templates/main.go.tmpl")

	if err != nil {
		return "", fmt.Errorf("error reading template: %w", err)
	}

	tpl, err := template.New("client").Funcs(funcs).Parse(string(mainTpl))

	if err != nil {
		return "", fmt.Errorf("error compiling template: %w", err)
	}

	var endpoints []endpointData

	for _, ep := range y.Endpoints {
		params := extractPathParams(ep.Path)
		fmtPath := pathParamRegex.ReplaceAllString(ep.Path, "%v")

		endpoints = append(endpoints, endpointData{
			Name:       ep.Name,
			Method:     ep.Method,
			Path:       ep.Path,
			FmtPath:    fmtPath,
			PathParams: params,
			Request:    ep.Request,
			Responses:  ep.Responses,
		})
	}

	data := templateData{
		PackageName: goGeneratorConfig.Package,
		Endpoints:   endpoints,
	}

	var buf bytes.Buffer

	if err := tpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("error executing template: %w", err)
	}

	return buf.String(), nil
}

func extractPathParams(path string) []string {
	matches := pathParamRegex.FindAllStringSubmatch(path, -1)
	var params []string
	for _, m := range matches {
		if len(m) > 1 {
			params = append(params, m[1])
		}
	}
	return params
}
