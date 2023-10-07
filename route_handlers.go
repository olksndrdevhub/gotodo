package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// auth
func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)
	template, err := template.ParseFiles("templates/base.html", "templates/register.html")
	if err != nil {
		log.Fatal(err)
	}
	if r.Method == "POST" {
		log.Println(r.FormValue("first_name"))
		log.Println(r.FormValue("last_name"))
		log.Println(r.FormValue("email"))
		log.Println(r.FormValue("password"))

		w = AddNotification(w, r, "Registration successful", "success")
	}
	err = template.ExecuteTemplate(w, "base.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

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
