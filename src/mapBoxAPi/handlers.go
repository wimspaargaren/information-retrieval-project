package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/unrolled/render"
)

var db *sql.DB

//Index gives some info about the API.
func Index(w http.ResponseWriter, r *http.Request) {

}

//GetAllIssues retrieves all issues.
func GetPoints(w http.ResponseWriter, r *http.Request) {
	var err error
	db, err = "withoutpass"
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	rows, errQ := db.Query("SELECT lat,long,category FROM data Where category = 'hardlopen' limit 10")
	if errQ != nil {
		fmt.Printf("error: %v\n", err)
	}
	// jsonstring := `"type": "FeatureCollection", "features": [`
	var result Response
	result.Type = "FeatureCollection"
	for rows.Next() {
		var resp Features
		var lat float64
		var long float64
		var category string
		err = rows.Scan(&lat, &long, &category)
		if err != nil {
			fmt.Println("this did not work")
		}
		resp.Type = "Feature"
		resp.Geometry.Type = "Point"
		resp.Geometry.Coordinates[0] = lat
		resp.Geometry.Coordinates[1] = long
		resp.Properties.Category = category
		result.Features = append(result.Features, resp)
		fmt.Println(lat)
		fmt.Println(long)
		fmt.Println(category)
		// jsonstring += `{ "type": "Feature", "geometry": { "type": "Point", "coordinates": [` + strconv.FormatFloat(lat, 'f', 6, 64) + `,` + strconv.FormatFloat(long, 'f', 6, 64) + `]},`
		// jsonstring += `"properties": { "sport-category": "` + category + `" }`
	}
	// jsonstring += `)]`

	//bla := json.Unmarshal(jsonstring)
	res, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	render := render.New()
	render.JSON(w, http.StatusOK, string(res))
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
}
