package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/strava/go.strava"
)

func main() {
	var accessToken string
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
	datarunning1, err := strava.NewSegmentsService(client).Explore(latmid, lonleft, lathigh, latmid).ActivityType("running").Do()
	datariding1, err := strava.NewSegmentsService(client).Explore(latmid, lonleft, lathigh, latmid).ActivityType("biking").Do()
	datarunning2, err := strava.NewSegmentsService(client).Explore(latmid, lonmid, lathigh, lonright).ActivityType("running").Do()
	datariding2, err := strava.NewSegmentsService(client).Explore(latmid, lonmid, lathigh, lonright).ActivityType("biking").Do()
	datarunning3, err := strava.NewSegmentsService(client).Explore(latlow, lonleft, latmid, lonmid).ActivityType("running").Do()
	datariding3, err := strava.NewSegmentsService(client).Explore(latlow, lonleft, latmid, lonmid).ActivityType("biking").Do()
	datarunning4, err := strava.NewSegmentsService(client).Explore(latlow, lonmid, latmid, lonright).ActivityType("running").Do()
	datariding4, err := strava.NewSegmentsService(client).Explore(latlow, lonmid, latmid, lonright).ActivityType("biking").Do()
	
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//Quadrant 1
	for _, segment := range datarunning1 {
		fmt.Printf("Fetching new leaderboard...\n")
		results, err := strava.NewSegmentsService(client).GetLeaderboard(segment.Id).Do()
		fmt.Printf("Printing running leaderboard...\n")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, e := range results.Entries {
			fmt.Printf("%5d: %5d %s\n", e.Rank, e.ElapsedTime, e.AthleteName)
		}
	}
	for _, segment := range datariding1 {
		fmt.Printf("Fetching new leaderboard...\n")
		results, err := strava.NewSegmentsService(client).GetLeaderboard(segment.Id).Do()	
		fmt.Printf("Printing riding leaderboard...\n")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, e := range results.Entries {
			fmt.Printf("%5d: %5d %s\n", e.Rank, e.ElapsedTime, e.AthleteName)
		}
	}

	//Quadrant 2
	for _, segment := range datarunning2 {
		fmt.Printf("Fetching new leaderboard...\n")
		results, err := strava.NewSegmentsService(client).GetLeaderboard(segment.Id).Do()
		fmt.Printf("Printing running leaderboard...\n")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, e := range results.Entries {
			fmt.Printf("%5d: %5d %s\n", e.Rank, e.ElapsedTime, e.AthleteName)
		}
	}
	for _, segment := range datariding2 {
		fmt.Printf("Fetching new leaderboard...\n")
		results, err := strava.NewSegmentsService(client).GetLeaderboard(segment.Id).Do()	
		fmt.Printf("Printing riding leaderboard...\n")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, e := range results.Entries {
			fmt.Printf("%5d: %5d %s\n", e.Rank, e.ElapsedTime, e.AthleteName)
		}
	}
	//Quadrant 3
	for _, segment := range datarunning3 {
		fmt.Printf("Fetching new leaderboard...\n")
		results, err := strava.NewSegmentsService(client).GetLeaderboard(segment.Id).Do()
		fmt.Printf("Printing running leaderboard...\n")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, e := range results.Entries {
			fmt.Printf("%5d: %5d %s\n", e.Rank, e.ElapsedTime, e.AthleteName)
		}
	}
	for _, segment := range datariding3 {
		fmt.Printf("Fetching new leaderboard...\n")
		results, err := strava.NewSegmentsService(client).GetLeaderboard(segment.Id).Do()	
		fmt.Printf("Printing riding leaderboard...\n")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, e := range results.Entries {
			fmt.Printf("%5d: %5d %s\n", e.Rank, e.ElapsedTime, e.AthleteName)
		}
	}
	//Quadrant 4
	for _, segment := range datarunning4 {
		fmt.Printf("Fetching new leaderboard...\n")
		results, err := strava.NewSegmentsService(client).GetLeaderboard(segment.Id).Do()
		fmt.Printf("Printing running leaderboard...\n")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, e := range results.Entries {
			fmt.Printf("%5d: %5d %s\n", e.Rank, e.ElapsedTime, e.AthleteName)
		}
	}
	for _, segment := range datariding4 {
		fmt.Printf("Fetching new leaderboard...\n")
		results, err := strava.NewSegmentsService(client).GetLeaderboard(segment.Id).Do()	
		fmt.Printf("Printing riding leaderboard...\n")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, e := range results.Entries {
			fmt.Printf("%5d: %5d %s\n", e.Rank, e.ElapsedTime, e.AthleteName)
		}
	}
}
