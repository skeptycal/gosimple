package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/skeptycal/gosimple/cli/errorlogger"
)

var log = errorlogger.New()

// local path: Users/skeptycal/go/src/github.com/skeptycal/gosimple/templates/repofiles
// local path: $GOPATH/src/github.com/skeptycal/gosimple/templates/repofiles

func main() {

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("pwd: ", pwd)

	template, err := parseFiles(".gitignore")
	// gi := template

	if err != nil {
		log.Fatal(err)
	}
	template.New(githubTemplate)

	// Print out the template to std
	err = template.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// https://www.toptal.com/developers/gitignore/api/macos,linux,windows,ssh,vscode,go,zsh,node,vue,nuxt,python,django,yarn,git

func parseFiles(filenames ...string) (*template.Template, error) {
	// for _, filename := range filenames {
	// 	filename = filepath.Join(repoFilesPath, filename)
	// 	fmt.Println(filename)
	// }
	return template.ParseFiles(filenames...)
}

// func template.ParseFiles(filenames ...string) (*template.Template, error)

const (
	githubTemplate = `{{header}}

{{personal}}

{{repoSpecific}}

{{security}}

{{divider}}

{{giFile}}
`

	repoFilesPath = "../../../repofiles"

	header = `# Copyright (c) 2021 Michael Treanor
# https://github.com/skeptycal
# MIT License

# Template Testm`

	personalItems = `### Personal ###
**/[Bb]ak/
**/*.bak
**/*temp
**/*tmp
**/.waka*
.vscode`

	repoSpecificItems = `### Repo Specific ###
**/idea.md`

	securityItems = `### Security ###
**/*[Tt]oken*
**/*[Pp]rivate*
**/*[Ss]ecret*
*history*
*hst*`

	divider = `############################################`
)
