package main

import (
	"flag"
	"github.com/Masterminds/sprig"
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// more on Flags https://flaviocopes.com/go-command-line-flags/
	var flagDatafilePath = flag.String("d", "", "Path to data file which contains YAML or JSON")
	var flagTemplatePath = flag.String("t", "", "Path to tpl file")
	flag.Parse()

	dataBytes, readError := ioutil.ReadFile(*flagDatafilePath)
	checkError(readError)

	var templateData map[string]interface{}
	parseError := yaml.Unmarshal([]byte(dataBytes), &templateData) //reads yaml and json because json is just a subset of yaml
	checkError(parseError)

	basePath := filepath.Base(*flagTemplatePath)
	tmpl, templateError := template.New(basePath).Funcs(sprig.FuncMap()).ParseFiles(*flagTemplatePath)
	checkError(templateError)

	checkError(tmpl.ExecuteTemplate(os.Stdout, basePath, templateData))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
