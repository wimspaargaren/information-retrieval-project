package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"github.com/strava/go.strava"
	sql "database/sql"
	_ "github.com/lib/pq"
)


func main() {
	var accessToken string
	var db *sql.DB
	var lathigh, lathighmid, latmid, latmidlow, latlow, lonleft, lonleftmid, lonmid, lonmidright, lonright float64 =52.421035, 52.39661575, 52.3721965, 52.34777725, 52.323358, 4.743425, 4.8156945, 4.887964, 4.9602335, 5.032503


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
	//lonleft, lonleftmid, lonmid, lonmidright, lonright
	//---------------------lathigh
	//|	1 |	 2	|  3 |  4 |
	//---------------------lathighmid
	//| 5 |  6  |  7 |  8 |
	//---------------------latmid
	//| 9	|	10  | 11 | 12 |
	//---------------------latmidlow
	//|13 | 14  | 15 | 16 |
	//---------------------latlow
	fmt.Printf("Fetching segment info...\n")
	datarunning1, err := strava.NewSegmentsService(client).Explore(lathighmid, lonleft, lathigh, lonleftmid).ActivityType("running").Do()
	datarunning2, err := strava.NewSegmentsService(client).Explore(lathighmid, lonleftmid, lathigh, lonmid).ActivityType("running").Do()
	datarunning3, err := strava.NewSegmentsService(client).Explore(lathighmid, lonmid, lathigh, lonmidright).ActivityType("running").Do()
	datarunning4, err := strava.NewSegmentsService(client).Explore(lathighmid, lonmidleft, lathigh, lonright).ActivityType("running").Do()
	
	datarunning5, err := strava.NewSegmentsService(client).Explore(latmid, lonleft, lathighmid, lonleftmid).ActivityType("running").Do()
	datarunning6, err := strava.NewSegmentsService(client).Explore(latmid, lonleftmid, lathighmid, lonmid).ActivityType("running").Do()
	datarunning7, err := strava.NewSegmentsService(client).Explore(latmid, lonmid, lathighmid, lonmidright).ActivityType("running").Do()
	datarunning8, err := strava.NewSegmentsService(client).Explore(latmid, lonmidleft, lathighmid, lonright).ActivityType("running").Do()
	
	datarunning9, err := strava.NewSegmentsService(client).Explore(latmidlow, lonleft, latmid, lonleftmid).ActivityType("running").Do()
	datarunning10, err := strava.NewSegmentsService(client).Explore(latmidlow, lonleftmid, latmid, lonmid).ActivityType("running").Do()
	datarunning11, err := strava.NewSegmentsService(client).Explore(latmidlow, lonmid, latmid, lonmidright).ActivityType("running").Do()
	datarunning12, err := strava.NewSegmentsService(client).Explore(latmidlow, lonmidleft, latmid, lonright).ActivityType("running").Do()
	
	datarunning13, err := strava.NewSegmentsService(client).Explore(latlow, lonleft, latmidlow, lonleftmid).ActivityType("running").Do()
	datarunning14, err := strava.NewSegmentsService(client).Explore(latlow, lonleftmid, latmidlow, lonmid).ActivityType("running").Do()
	datarunning15, err := strava.NewSegmentsService(client).Explore(latlow, lonmid, latmidlow, lonmidright).ActivityType("running").Do()
	datarunning16, err := strava.NewSegmentsService(client).Explore(latlow, lonmidleft, latmidlow, lonright).ActivityType("running").Do()
	
	datariding1, err := strava.NewSegmentsService(client).Explore(lathighmid, lonleft, lathigh, lonleftmid).ActivityType("riding").Do()
	datariding2, err := strava.NewSegmentsService(client).Explore(lathighmid, lonleftmid, lathigh, lonmid).ActivityType("riding").Do()
	datariding3, err := strava.NewSegmentsService(client).Explore(lathighmid, lonmid, lathigh, lonmidright).ActivityType("riding").Do()
	datariding4, err := strava.NewSegmentsService(client).Explore(lathighmid, lonmidleft, lathigh, lonright).ActivityType("riding").Do()
	
	datariding5, err := strava.NewSegmentsService(client).Explore(latmid, lonleft, lathighmid, lonleftmid).ActivityType("riding").Do()
	datariding6, err := strava.NewSegmentsService(client).Explore(latmid, lonleftmid, lathighmid, lonmid).ActivityType("riding").Do()
	datariding7, err := strava.NewSegmentsService(client).Explore(latmid, lonmid, lathighmid, lonmidright).ActivityType("riding").Do()
	datariding8, err := strava.NewSegmentsService(client).Explore(latmid, lonmidleft, lathighmid, lonright).ActivityType("riding").Do()
	
	datariding9, err := strava.NewSegmentsService(client).Explore(latmidlow, lonleft, latmid, lonleftmid).ActivityType("riding").Do()
	datariding10, err := strava.NewSegmentsService(client).Explore(latmidlow, lonleftmid, latmid, lonmid).ActivityType("riding").Do()
	datariding11, err := strava.NewSegmentsService(client).Explore(latmidlow, lonmid, latmid, lonmidright).ActivityType("riding").Do()
	datariding12, err := strava.NewSegmentsService(client).Explore(latmidlow, lonmidleft, latmid, lonright).ActivityType("riding").Do()
	
	datariding13, err := strava.NewSegmentsService(client).Explore(latlow, lonleft, latmidlow, lonleftmid).ActivityType("riding").Do()
	datariding14, err := strava.NewSegmentsService(client).Explore(latlow, lonleftmid, latmidlow, lonmid).ActivityType("riding").Do()
	datariding15, err := strava.NewSegmentsService(client).Explore(latlow, lonmid, latmidlow, lonmidright).ActivityType("riding").Do()
	datariding16, err := strava.NewSegmentsService(client).Explore(latlow, lonmidleft, latmidlow, lonright).ActivityType("riding").Do()
	
	
	var list [32][]*strava.SegmentExplorerSegment
	list[0], list[1], list[2], list[3], list[4], list[5], list[6], list[7] = datarunning1, datarunning2, datarunning3, datarunning4, datarunning5, datarunning6, datarunning7, datarunning8
	list[8], list[9], list[10], list[11], list[12], list[13], list[14], list[15] = datarunning9, datarunning10, datarunning11, datarunning12, datarunning13, datarunning14, datarunning15, datarunning16
	list[16], list[17], list[18], list[19], list[20], list[21], list[22], list[23] = datariding1, datariding2, datariding3, datariding4, datariding5, datariding6, datariding7, datariding8
	list[24], list[25], list[26], list[27], list[28], list[29], list[30], list[31] = datariding9, datariding10, datariding11, datariding12, datariding13, datariding14, datariding15, datariding16
	
	
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var i int = 1
	var category string
	for _, element := range list {
		if i <= 4 {
			category = "running"
		} else {
			category = "cycling"
		}
		db, err = sql.Open("postgres", "postgres://user:pass@86.87.235.82:8082/twitter?sslmode=disable")
		
		for _, segment := range element {
			fmt.Printf("Fetching new leaderboard...\n")
			results, err := strava.NewSegmentsService(client).GetLeaderboard(segment.Id).Do()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			var start[2]float64 = segment.StartLocation
			//var end[2]float64 = segment.EndLocation
			
			for _, e := range results.Entries {
				var name string = e.AthleteName
				
				var t time.Time	= e.StartDateLocal
				var timeErr error
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
				tx, err := db.Begin()
				if err != nil {
					fmt.Println(err)
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
		i+=1
	}
	fmt.Println("done")
}
