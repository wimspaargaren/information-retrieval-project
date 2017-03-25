package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/unrolled/render"
)

var db *sql.DB
var constring string = os.Getenv("CONSTRING")

//GetAllIssues retrieves all issues.
func GetPoints(w http.ResponseWriter, r *http.Request) {
	var err error
	db, err = sql.Open("postgres", constring)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	rows, errQ := db.Query("SELECT lat,long,category, day, daypart FROM data Where category IS NOT NULL")
	if errQ != nil {
		fmt.Printf("error: %v\n", err)
	}
	var result Response
	result.Type = "FeatureCollection"
	for rows.Next() {
		var resp Features
		var lat float64
		var long float64
		var category string
		var day string
		var daypart string
		err = rows.Scan(&lat, &long, &category, &day, &daypart)
		if err != nil {
			fmt.Println("this did not work")
		}
		resp.Type = "Feature"
		resp.Geometry.Type = "Point"
		resp.Geometry.Coordinates[0] = long
		resp.Geometry.Coordinates[1] = lat
		resp.Properties.Category = category
		resp.Properties.Day = day
		resp.Properties.Daypart = daypart
		result.Features = append(result.Features, resp)
	}

	render := render.New()
	render.JSON(w, http.StatusOK, result)
}

func GetPolyGons(w http.ResponseWriter, r *http.Request) {
	var err error
	fmt.Println("CONSTRING: ", constring)
	db, err = sql.Open("postgres", constring)
	if err != nil {
		fmt.Printf("connection error: %v\n", err)
	}
	rows, errQ := db.Query("SELECT id,cluster_set_id,data_ids FROM clusters")
	if errQ != nil {
		fmt.Printf("error: %v\n", err)
	}
	var result Response
	result.Type = "FeatureCollection"
	for rows.Next() {
		// var id int
		// var clust_set_id int
		// var data_ids []int
		// err = rows.Scan(&id, &clust_set_id, &data_ids)
		// if err != nil {
		// 	fmt.Println("this did not work")
		// }
		// fmt.Println(data_ids)
	}
}

type Response struct {
	Type     string     `json:"type"`
	Features []Features `json:"features"`
}

type Features struct {
	Type       string   `json:"type"`
	Geometry   Geometry `json:"geometry"`
	Properties Props    `json:"properties"`
}

type Geometry struct {
	Type        string     `json:"type"`
	Coordinates [2]float64 `json:"coordinates"`
}

type Props struct {
	Category string `json:"sport-category"`
	Day      string `json:"day"`
	Daypart  string `json:"daypart"`
}
