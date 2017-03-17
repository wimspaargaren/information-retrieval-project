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
	jsonstring := `"type": "FeatureCollection", "features": [`
	for rows.Next() {

		var lat float64
		var long float64
		var category string
		err = rows.Scan(&lat, &long, &category)
		if err != nil {
			fmt.Println("this did not work")
		}
		fmt.Println(lat)
		fmt.Println(long)
		fmt.Println(category)
		jsonstring += `{ "type": "Feature", "geometry": { "type": "Point", "coordinates": [` + strconv.FormatFloat(lat, 'f', 6, 64) + `,` + strconv.FormatFloat(long, 'f', 6, 64) + `]},`
		jsonstring += `"properties": { "sport-category": "` + category + `" }`
	}
	jsonstring += `)]`

	//bla := json.Unmarshal(jsonstring)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	render := render.New()
	render.JSON(w, http.StatusOK, jsonstring)
}
