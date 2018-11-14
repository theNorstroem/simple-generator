# simple-generator

Very simple golang template engine (https://astaxie.gitbooks.io/build-web-application-with-golang/en/07.4.html) 
with template functions from http://masterminds.github.io/sprig/

Works with JSON or YMAL as input data format **only**

# Install

If you have Go installed:

```
go get github.com/veith/simple-generator
```

# Usage

```

 simple-generator -d data.json -t template.xxx // renders to stdout
 simple-generator -d data.json -t template.xxx > target.file
 
 simple-generator -h  // help 
 Usage of simple-generator:
   -d string
          Path to data file which contains YAML or JSON
   -t string
          Path to tpl file



```