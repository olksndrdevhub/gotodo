package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()
	// serve static files
	router.ServeFiles("/static/*filepath", http.Dir("static"))

	router.GET(("/"), Index)

	// auth
	router.GET("/register", Register)
	router.POST("/register", Register)

	log.Print("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
