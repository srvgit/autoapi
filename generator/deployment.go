package generator

import (
	"autoapi/graph/model"
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strings"
)

type DeploymentConfig struct {
	Features     []model.Feature
	TemplatePath string
	TargetPath   string
}

func (config *DeploymentConfig) GenerateDeployment() error {
	templateContent, err := os.ReadFile(config.TemplatePath)
	if err != nil {
		return fmt.Errorf("failed to read template file: %v", err)
	}

	tmpl, err := template.New("deployment").Parse(string(templateContent))
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, config); err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	err = os.WriteFile(config.TargetPath, buf.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to target file: %v", err)
	}

	fmt.Printf("Deployment YAML written to %s\n", config.TargetPath)
	return nil
}

func (d *DeploymentConfig) FeaturesString() string {
	var features []string
	for _, feature := range d.Features {
		features = append(features, string(feature))
	}
	fmt.Println(features)
	return strings.Join(features, ",")
}
