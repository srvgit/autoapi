package generator

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
)

type HPAConfig struct {
	MinReplicas  int
	TemplatePath string
	TargetPath   string
}

func (config *HPAConfig) GenerateHPA() error {
	templateContent, err := os.ReadFile(config.TemplatePath)
	if err != nil {
		return fmt.Errorf("failed to read template file: %v", err)
	}
	tmpl, err := template.New("hpa").Parse(string(templateContent))
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
	fmt.Printf("HPA YAML written to %s\n", config.TargetPath)
	return nil
}
