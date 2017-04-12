package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/unrolled/render"
)

var rOutput *render.Render
var constring string = os.Getenv("CONSTRING")

func main() {
	fmt.Println(os.Getenv("CONSTRING"))
	// setupDummyIssues()
	fmt.Println(constring)
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
	mux.HandleFunc("/getpoints", GetPoints).Methods("GET")
	mux.HandleFunc("/getpolygons", GetPolygons).Methods("GET")

	mux.HandleFunc("/voronoi", getFile).Methods("GET")

	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(mux)
	n.Run(":8080")
	// log.Fatal(http.ListenAndServe(":8080", router))
}

func getFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "shapefile.zip")
}
