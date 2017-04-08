package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"github.com/strava/go.strava"
	sql "database/sql"
)


func main() {
	var accessToken string
	var db *sql.DB
	var lathigh, latmid, latlow, lonleft, lonmid, lonright float64 = 52.421035, 52.3721965, 52.323358, 4.743425, 4.887964, 5.032503


	// Provide an access token, with write permissions.
	// You'll need to complete the oauth flow to get one.
	flag.StringVar(&accessToken, "token", "1dd663e33d658e4254f4acd5b46d4e0d2eacbda8", "Access Token")
	flag.Parse()

	if accessToken == "" {
		fmt.Println("\nPlease provide an access_token, one can be found at https://www.strava.com/settings/api")

		flag.PrintDefaults()
		os.Exit(1)
	}
	
	client := strava.NewClient(accessToken)
	
	//Quadrants of Amsterdam
	//---------------------
	//|	1 			|  2      |
	//|         |         |
	//---------------------
	//| 3				|  4      |
	//|         |         |
	//---------------------
	
	fmt.Printf("Fetching segment info...\n")
	datarunning1, err := strava.NewSegmentsService(client).Explore(latmid, lonleft, lathigh, lonmid).ActivityType("running").Do()
	datariding1, err := strava.NewSegmentsService(client).Explore(latmid, lonleft, lathigh, lonmid).ActivityType("biking").Do()
	datarunning2, err := strava.NewSegmentsService(client).Explore(latmid, lonmid, lathigh, lonright).ActivityType("running").Do()
	datariding2, err := strava.NewSegmentsService(client).Explore(latmid, lonmid, lathigh, lonright).ActivityType("biking").Do()
	datarunning3, err := strava.NewSegmentsService(client).Explore(latlow, lonleft, latmid, lonmid).ActivityType("running").Do()
	datariding3, err := strava.NewSegmentsService(client).Explore(latlow, lonleft, latmid, lonmid).ActivityType("biking").Do()
	datarunning4, err := strava.NewSegmentsService(client).Explore(latlow, lonmid, latmid, lonright).ActivityType("running").Do()
	datariding4, err := strava.NewSegmentsService(client).Explore(latlow, lonmid, latmid, lonright).ActivityType("biking").Do()
	var list[8]strava.SegmentExplorerSegment = [datarunning1, datarunning2, datarunning3, datarunning4, datariding1, datariding2, datariding3, datariding4]
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var i int = 1
	for_, element := range list {
		if i <= 4 {
			getLeaderboards(element, "running")
		}
		else {
			getLeaderboards(element, "cycling")
		}
		i+=1
	}
	
	func getLeaderboards(seg *strava.SegmentExplorerSegment, cat string)	
		db, err = sql.Open("postgres", "postgres://user:pass@86.87.235.82:8082/twitter?sslmode=disable")
		tx, err := db.Begin()
		for _, segment := range seg {
			fmt.Printf("Fetching new leaderboard...\n")
			results, err := strava.NewSegmentsService(client).GetLeaderboard(segment.Id).Do()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			var category, id = cat, segment.Id
			var start[2]float64 = segment.StartLocation
			//var end[2]float64 = segment.EndLocation
			
			for _, e := range results.Entries {
				var name string = e.AthleteName
				
				var t time.Time	
				var timeErr error
				t, timeErr = time.Parse(time.RubyDate, e.StartDateLocal)
				if timeErr != nil {
					fmt.Println(timeErr)
				}
				day := t.Weekday()
				var daypart string
				h := t.Hour()
				if h >= 0 && h < 6 {
					daypart = "Night"
				} else if h >= 6 && h < 12 {
					daypart = "Morning"
				} else if h >= 12 && h < 18 {
					daypart = "Midday"
				} else {
					daypart = "Evening"
				}
				
				var id int
				dbErr := tx.QueryRow("INSERT INTO strava (category, athlete_name, lat, long, day, daypart) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
					category, name, start[0], start[1], day.String(), daypart).Scan(&id)
				if dbErr != nil {
					fmt.Println(dbErr)
				}
				tx.Commit()
			}
		}
	}
}
