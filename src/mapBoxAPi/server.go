package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/unrolled/render"
)

var rOutput *render.Render

func main() {
	// setupDummyIssues()
	// router := NewRouter()
	var origins []string
	origins = append(origins, "*")

	c := cors.New(cors.Options{
		AllowedOrigins:   origins,
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "DELETE", "PUT"},
	})

	mux := mux.NewRouter()

	// Routes
	mux.HandleFunc("/index", Index).Methods("GET")
	mux.HandleFunc("/getpoints", GetPoints).Methods("GET")

	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(mux)
	n.Run(":8080")
	// log.Fatal(http.ListenAndServe(":8080", router))
}
