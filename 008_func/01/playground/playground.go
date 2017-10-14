package main

import (
	"html/template"
	"os"
)

func main() {
	tpl := template.Must(template.New("somthing").Parse("here us the text in the template"))
	tpl.ExecuteTemplate(os.Stdout, "something", nil)
}
