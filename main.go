package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Masterminds/sprig"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
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

	jsonFile, openFileError := os.Open(*flagJSON)
	if openFileError != nil {
		fmt.Println(openFileError)
	}

	defer jsonFile.Close()
	jsonStringArray, readError := ioutil.ReadAll(jsonFile)
	if readError != nil {
		fmt.Println(readError)
	}

	var jsonData map[string]interface{} // Objekt aus dem json file

	if parseError := json.Unmarshal([]byte(jsonStringArray), &jsonData); parseError != nil {
		log.Fatal(parseError)
	}

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
