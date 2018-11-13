# simple-generator

golang default template engine (https://astaxie.gitbooks.io/build-web-application-with-golang/en/07.4.html) with template functions from http://masterminds.github.io/sprig/

Works with JSON as input data format **only**

# Install

If you have Go installed:

```
go get github.com/veith/simple-generator
```

# Usage

```

 simple-generator -j data.json -t template.xxx // renders to stdout
 simple-generator -j data.json -t template.xxx > target.file
 
 simple-generator -h  // help 
 Usage of simple-generator:
   -j string
         shortcut for json
   -json string
         Path to JSON file
   -t string
         shortcut for template
   -template string
         Path to template


```