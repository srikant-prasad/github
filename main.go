package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func index(w http.ResponseWriter, r *http.Request)  {
	// io.WriteString(w, "This is Home pages!")
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func main()  {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}