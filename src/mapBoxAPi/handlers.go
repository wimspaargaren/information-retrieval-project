package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/unrolled/render"
)

var db *sql.DB

//GetPoints retrieves all points.
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

//GetPolygons retrieves all polygons.
func GetPolygons(w http.ResponseWriter, r *http.Request) {
	var err error
	fmt.Println("CONSTRING: ", constring)
	db, err = sql.Open("postgres", constring)
	if err != nil {
		fmt.Printf("connection error: %v\n", err)
	}
	rows, errQ := db.Query("SELECT id, UNNEST(data_ids),category FROM clusters WHERE cluster_set_id = 13")
	if errQ != nil {
		fmt.Printf("error: %v\n", err)
	}

	var result PolyResponse
	result.Type = "FeatureCollection"
	idTemp := -1
	var intArray []int
	var test string
	for rows.Next() {
		var id int
		var tweetID int
		var category string
		err = rows.Scan(&id, &tweetID, &category)
		if err != nil {
			fmt.Println("this did not work", err)
		}
		if idTemp != id {
			if test != "" {
				//Polygon feature
				var resp PolygonFeature
				resp.Type = "Feature"
				resp.Geometry.Type = "Polygon"
				resp.Properties.Category = category
				resp.Properties.ID = id
				rows2, errQ := db.Query("SELECT lat, long FROM data WHERE " + test)
				tempSlice := [][2]float64{}
				if errQ != nil {
					fmt.Printf("error: %v\n", err)
				}
				for rows2.Next() {
					var long float64
					var lat float64
					err = rows2.Scan(&lat, &long)
					var test2 [2]float64
					test2[0] = long
					test2[1] = lat
					tempSlice = append(tempSlice, test2)

				}
				resp.Geometry.Coordinates = append(resp.Geometry.Coordinates, tempSlice)
				result.Features = append(result.Features, resp)

			}

			idTemp = id
			intArray = []int{}
			test = "id = " + strconv.Itoa(tweetID)
			intArray = append(intArray, tweetID)
		} else {
			test += " OR id = " + strconv.Itoa(tweetID)

			intArray = append(intArray, tweetID)
		}
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
