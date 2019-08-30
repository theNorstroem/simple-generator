# simple-generator

Very simple golang template engine for cli usage, using the golang template package and the template functions from [sprig](http://masterminds.github.io/sprig/). 

Package template implements data-driven templates for generating textual output.

Simple-generator works with JSON or YMAL as input data format **only** 

## Dokumentation 
 - [Build web application with Golang, 7.4 Templates](https://astaxie.gitbooks.io/build-web-application-with-golang/en/07.4.html)
 - [**sprig**, Useful template functions for Go templates](http://masterminds.github.io/sprig/)




# Install

If you have Go 1.12.x installed:

```
GO111MODULE=on go get -u github.com/veith/simple-generator@v0.1.3
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

