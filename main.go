package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)
	template, err := template.ParseFiles("templates/base.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)
	template, err := template.ParseFiles("templates/base.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	// serve static files
	router.ServeFiles("/static/*filepath", http.Dir("static"))

	router.GET(("/"), Index)
	router.GET("/hello/:name", Hello)

	log.Print("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
