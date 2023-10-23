package util

import (
	"autoapi/generator"
	"autoapi/graph/model"
	"fmt"
)

func CreateKustomizations(config *model.ServerConfig) error {
	directory := "./local-repo"
	err := CleanUp(directory)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	repoURL := "https://github.com/srvgit/autoapi-k8"

	err = CloneRepo(repoURL, directory)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	srcDir := "templates/argo/graphlet"
	dstDir := directory + "/overlays/" + config.ApiserverName
	extPattern := ".yml"
	err = CopyDir(srcDir, dstDir, extPattern)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	kustomizeConfig := generator.KustomizationConfig{
		AppLabel:         config.ApiserverName,
		TemplateFilePath: "templates/argo/graphlet/base/kustomization.tmpl",
		TargetFilePath:   directory + "/overlays/" + config.ApiserverName + "/base" + "/kustomization.yml",
	}
	kustomizeConfig.GenerateKustomization()

	derivedConfig := generator.MapExpectationsToResources(config.PerformanceRequirements)
	hpa := generator.HPAConfig{
		MinReplicas:  derivedConfig.MinReplicas,
		TemplatePath: "templates/argo/graphlet/dev/environment-patch-hpa.tmpl",
		TargetPath:   directory + "/overlays/" + config.ApiserverName + "/dev" + "/environment-patch-hpa.yml",
	}
	if err = hpa.GenerateHPA(); err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}
	deploymentConfig := generator.DeploymentConfig{
		Features:     config.Features,
		TemplatePath: "templates/argo/graphlet/dev/environment-patch-deployment.tmpl",
		TargetPath:   directory + "/overlays/" + config.ApiserverName + "/dev" + "/environment-patch-deployment.yml",
	}

	if err = deploymentConfig.GenerateDeployment(); err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}
	vpa := generator.VPAConfig{
		MaxCPU:       derivedConfig.MaxCPU,
		MaxMemory:    derivedConfig.MaxMemory,
		MinCPU:       derivedConfig.MinCPU,
		MinMemory:    derivedConfig.MinMemory,
		TemplatePath: "templates/argo/graphlet/dev/environment-patch-vpa.tmpl",
		TargetPath:   directory + "/overlays/" + config.ApiserverName + "/dev" + "/environment-patch-vpa.yml",
	}
	if err = vpa.GenerateVPA(); err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	appConfig := generator.ArgoCDAppConfig{
		AppName:      config.ApiserverName,
		TemplatePath: "templates/argo/graphlet/app.tmpl",
		TargetPath:   directory + "/apps/" + config.ApiserverName + "-app.yml",
	}

	appConfig.GenerateArgoCDApp()

	if err = CommitAndPush(directory, "AutoAPI: Add new graphlet:"+config.ApiserverName); err != nil {
		return err
	}

	return nil
}
