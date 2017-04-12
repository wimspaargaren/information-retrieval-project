package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/unrolled/render"
)

var db *sql.DB

//GetPoints retrieves all points.
func GetPoints(w http.ResponseWriter, r *http.Request) {
	var err error
	db, err = sql.Open("postgres", constring)
	if err != nil {
		fmt.Printf("errorrr: %v\n", err)
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
			fmt.Println("this did not work", err)
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

type polyFullResponse struct {
	Clusters []Cluster `json:"clusters"`
}

type Cluster struct {
	IDs      []int  `json:"ids"`
	Category string `json:"category"`
}

//GetPolygons retrieves all polygons.
func GetPolygons(w http.ResponseWriter, r *http.Request) {
	var err error
	fmt.Println("CONSTRING: ", constring)
	db, err = sql.Open("postgres", constring)
	if err != nil {
		fmt.Printf("connection error: %v\n", err)
	}
	rows, errQ := db.Query("SELECT id, unnest(data_ids), category FROM clusters WHERE cluster_set_id = 15")
	if errQ != nil {
		fmt.Printf("error: %v\n", err)
	}

	var result polyFullResponse
	idTemp := -1
	// var intArray []int
	// var test string
	for rows.Next() {
		var id int
		var tweetID int
		var category string
		err = rows.Scan(&id, &tweetID, &category)
		if err != nil {
			fmt.Println("this did not work", err)
		}

		if id != idTemp {
			var c Cluster
			c.IDs = append(c.IDs, tweetID)
			c.Category = category
			result.Clusters = append(result.Clusters, c)
		} else {
			last := len(result.Clusters) - 1
			c := result.Clusters[last]
			c.IDs = append(c.IDs, tweetID)
			result.Clusters[last] = c
		}
		idTemp = id
	}
	render := render.New()
	render.JSON(w, http.StatusOK, result)
}

//GetPolygonsStrava retrieves all polygons.
func GetPolygonsStrava(w http.ResponseWriter, r *http.Request) {
	var err error
	fmt.Println("CONSTRING: ", constring)
	db, err = sql.Open("postgres", constring)
	if err != nil {
		fmt.Printf("connection error: %v\n", err)
	}
	rows, errQ := db.Query("SELECT id, unnest(data_ids), category FROM clusters WHERE cluster_set_id = 16")
	if errQ != nil {
		fmt.Printf("error strava: %v\n", err)
	}

	var result polyFullResponse
	idTemp := -1
	// var intArray []int
	// var test string
	for rows.Next() {
		var id int
		var tweetID int
		var category string
		err = rows.Scan(&id, &tweetID, &category)
		if err != nil {
			fmt.Println("this did not work", err)
		}

		if id != idTemp {
			var c Cluster
			c.IDs = append(c.IDs, tweetID)
			c.Category = category
			result.Clusters = append(result.Clusters, c)
		} else {
			last := len(result.Clusters) - 1
			c := result.Clusters[last]
			c.IDs = append(c.IDs, tweetID)
			result.Clusters[last] = c
		}
		idTemp = id
	}
	render := render.New()
	render.JSON(w, http.StatusOK, result)
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

type PolyResponse struct {
	Type     string           `json:"type"`
	Features []PolygonFeature `json:"features"`
}

type PolygonFeature struct {
	Type       string                 `json:"type"`
	Geometry   PolygonCoordinatesList `json:"geometry"`
	Properties PolyProp               `json:"properties"`
}

type PolyProp struct {
	Category string `json:"sport-category"`
	ID       int    `json:"id"`
}

type PolygonCoordinatesList struct {
	Type        string         `json:"type"`
	Coordinates [][][2]float64 `json:"coordinates"`
}
