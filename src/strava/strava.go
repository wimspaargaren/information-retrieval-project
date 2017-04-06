package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/strava/go.strava"
)

func main() {
	var accessToken string
	var swlat, swlon, nelat, nelon float64 = 52.323358, 4.743425, 52.421035, 5.032503
	
	
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
	
	fmt.Printf("Fetching segment info...\n")
	datarunning, err := strava.NewSegmentsService(client).Explore(swlat, swlon, nelat, nelon).ActivityType("running").Do()
	datariding, err := strava.NewSegmentsService(client).Explore(swlat, swlon, nelat, nelon).ActivityType("biking").Do()
	
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	for _, segment := range datarunning {
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
	for _, segment := range datariding {
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
