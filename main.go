package main

import (
	"bufio"
	"flag"
	"io"
	"io/fs"
	"os"
	"strings"
	"text/template"
)

type DeploymentTemplate struct {
	Registry   string
	AppVersion string
	AppName    string
}

func main() {
	var deploymentFile string
	// var appName string
	// var appVersion string
	// var registry string

	args := make(map[string]string)
	for _, s := range os.Environ() {
		split := strings.Split(s, "=")
		args[split[0]] = split[1]
	}

	// fmt.Println(args["USER"])

	// flag.StringVar(&appName, "appName", "", "specific app's name to deploy")
	// flag.StringVar(&appVersion, "appVersion", "latest", "specific app's version to deploy")
	flag.StringVar(&deploymentFile, "f", "", "specific deployment template")
	// flag.StringVar(&registry, "registry", "github.com", "registry url")

	flag.Parse()

	var tpl *template.Template
	var err error
	var info fs.FileInfo

	if info, err = os.Stdin.Stat(); err != nil {
		panic(err)
	}

	if deploymentFile == "" || deploymentFile == "-" {
		if info.Size() == 0 {
			flag.Usage()
			os.Exit(1)
		} else {
			reader := bufio.NewReader(os.Stdin)
			buffer := new(strings.Builder)
			if _, err := io.Copy(buffer, reader); err != nil {
				panic(err)
			}
			if tpl, err = template.New("stdin").Parse(buffer.String()); err != nil {
				panic(err)
			}
		}
	} else {
		if tpl, err = template.ParseFiles(deploymentFile); err != nil {
			panic(err)
		}
	}

	if err := tpl.Execute(os.Stdout, args); err != nil {
		panic(err)
	}
}
