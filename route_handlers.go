package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)
	template, err := template.ParseFiles("templates/base.html", "templates/partials/_header.html")
	if err != nil {
		log.Fatal(err)
	}
	err = template.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)
	template, err := template.ParseFiles("templates/base.html")
	if err != nil {
		log.Fatal(err)
	}
	err = template.Execute(w, ps.ByName("name"))
	if err != nil {
		log.Fatal(err)
	}

}
