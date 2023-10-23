package generator

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
)

type KustomizationConfig struct {
	AppLabel         string
	TemplateFilePath string
	TargetFilePath   string
}

func (config *KustomizationConfig) GenerateKustomization() error {
	templateContent, err := os.ReadFile(config.TemplateFilePath)
	if err != nil {
		return fmt.Errorf("failed to read template file: %v", err)
	}

	tmpl, err := template.New("kustomization").Parse(string(templateContent))
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, config); err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	err = os.WriteFile(config.TargetFilePath, buf.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to target file: %v", err)
	}

	fmt.Printf("Kustomization YAML written to %s\n", config.TargetFilePath)
	return nil
}
