package main

import (
	"encoding/json"
	"flag"
	"github.com/Masterminds/sprig"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// more on Flags https://flaviocopes.com/go-command-line-flags/
	var flagTemplate = flag.String("t", "", "Path to tpl file")
	var flagDatafile = flag.String("d", "", "Path to data file which contains YAML or JSON")
	flag.Parse()

	dataBytes, readError := ioutil.ReadFile(*flagDatafile)
	checkError(readError)

	var templateData map[string]interface{}
	parseError := json.Unmarshal([]byte(dataBytes), &templateData)
	checkError(parseError)

	basePath := filepath.Base(*flagTemplate)
	tmpl, templateError := template.New(basePath).Funcs(sprig.FuncMap()).ParseFiles(*flagTemplate)
	checkError(templateError)

	checkError(tmpl.ExecuteTemplate(os.Stdout, basePath, templateData))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
