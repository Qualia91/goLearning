package main

import (
	"fmt"
	"os"
	"text/template"
)

type BlogPost struct {
	Title   string
	Content string
}

func main() {
	// simple template object substitution
	post := BlogPost{"First blog post", "This is the post"}
	createAndParseTemplate(post, "blog-tmpl", `<h1>{{.Title}}</h1><div><p>{{.Content}}</p></div>`)

	println()

	// loops
	items := []string{"one", "two", "three"}
	createAndParseTemplate(items, "loops", `{{range .}} number {{.}} and then {{else}} List was empty {{end}}`)

	println()

	itemsEmpty := []string{}
	createAndParseTemplate(itemsEmpty, "loops", `{{range .}} number {{.}} and then {{else}}List was empty {{end}}`)

	println()

	// functions
	tmplString := "{{range $index, $element := .}}{{if mod $index 2}}{{.}}{{end}}{{end}}"
	fm := template.FuncMap{"mod": func(i, j int) bool { return i%j == 0 }}
	tmpl, _ := template.New("tmplt").Funcs(fm).Parse(tmplString)
	err := tmpl.Execute(os.Stdout, items)
	if err != nil {
		fmt.Printf("Error %s", err)
	}

	println()
}

func createAndParseTemplate(obj interface{}, name, templateString string) {
	tmpl, err := template.New(name).Parse(templateString)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, obj)
	if err != nil {
		panic(err)
	}
}
