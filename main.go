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

	// Flags https://flaviocopes.com/go-command-line-flags/
	var (
		flagTemplate      = flag.String("template", "", "Path to template")
		flagTemplateShort = flag.String("t", "", "shortcut for template")
		flagJSON          = flag.String("json", "", "Path to JSON file")
		flagJSONShort     = flag.String("j", "", "shortcut for json")
	)

	flag.Parse()

	if len(*flagJSONShort) > 0 {
		flagJSON = flagJSONShort
	}
	if len(*flagTemplateShort) > 0 {
		flagTemplate = flagTemplateShort
	}

	// open the file from flag flagJSON for reading
	jsonFile, openFileError := os.Open(*flagJSON)
	logError(openFileError)

	defer jsonFile.Close()

	jsonStringArray, readError := ioutil.ReadAll(jsonFile)
	logError(readError)

	// Objekt aus dem json file
	var jsonData map[string]interface{}
	parseError := json.Unmarshal([]byte(jsonStringArray), &jsonData)
	logError(parseError)

	var (
		templateError error
		tmpl          *template.Template
		name          string
	)

	name = filepath.Base(*flagTemplate)
	tmpl, templateError = template.New(name).Funcs(sprig.FuncMap()).ParseFiles(*flagTemplate)

	if templateError != nil {
		log.Fatal(templateError)
	}

	if err := tmpl.ExecuteTemplate(os.Stdout, name, jsonData); err != nil {
		log.Fatal(err)
	}
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
