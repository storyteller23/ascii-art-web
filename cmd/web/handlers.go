package main

import (
	asciiart "asciiartweb/src"
	"html/template"
	"net/http"
	"strings"
)

type data struct {
	Result string
	Color  string `default:"black"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	result := data{}
	tmpl, err := template.ParseFiles("ui/html/index.html")
	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	}

	tmpl.Execute(w, result)
}

func PostAsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	result := data{}
	tmpl, err := template.ParseFiles("ui/html/index.html")
	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	}

	r.ParseForm()
	str, ok := r.Form["text"]
	if !ok {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	tmplAscii, ok := r.Form["template"]
	if !ok {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	color := r.Form.Get("color")
	result.Color = color
	template, err := asciiart.ParseTemplateToMap("./ui/static/asciiart-templates/" + strings.Join(tmplAscii, "") + ".txt")
	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	} else {
		art, err := asciiart.StringToAsciiArt(strings.Join(str, ""), template)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		if result.Result == "" {
			result.Result = art
		}
	}

	tmpl.Execute(w, result)
}
