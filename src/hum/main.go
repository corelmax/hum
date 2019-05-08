package main

import (
	"flag"
	"os"
	"text/template"
)

type DeploymentTemplate struct {
	Registry   string
	AppVersion string
}

func main() {
	var deploymentFile string
	var appVersion string
	var registry string
	flag.StringVar(&appVersion, "appVersion", "latest", "specific app's version to deploy")
	flag.StringVar(&deploymentFile, "f", "", "specific deployment template")
	flag.StringVar(&registry, "registry", "github.com", "registry url")

	flag.Parse()

	if deploymentFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	deployment := DeploymentTemplate{
		AppVersion: appVersion,
		Registry:   registry,
	}

	var tpl *template.Template
	var err error

	if tpl, err = template.ParseFiles(deploymentFile); err != nil {
		panic(err)
	}

	if err := tpl.Execute(os.Stdout, deployment); err != nil {
		panic(err)
	}
}
