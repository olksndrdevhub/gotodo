package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)
	template, err := template.ParseFiles("templates/base.html", "templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	err = template.ExecuteTemplate(w, "base.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)
	template, err := template.ParseFiles("templates/base.html", "templates/hi.html")
	if err != nil {
		log.Fatal(err)
	}
	err = template.ExecuteTemplate(w, "base.html", ps.ByName("name"))
	if err != nil {
		log.Fatal(err)
	}

}
