package generator

import (
	"os"
	"text/template"
)

type CodeGenerator interface {
	GenerateCode(data interface{}, templatePath string, outputPath string) error
}

type GoCodeGenerator struct{}

func (g *GoCodeGenerator) GenerateCode(data interface{}, templatePath string, outputPath string) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	return tmpl.Execute(outputFile, data)
}
