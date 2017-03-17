package main

import (
	sql "database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	opts := struct {
		TCKey    string
		TCSecret string
		TAToken  string
		TASecret string
	}{}
	flag.StringVar(&opts.TCKey, "tckey", opts.TCKey, "input for Twitter consumer key")
	flag.StringVar(&opts.TCSecret, "tcsecret", opts.TCSecret, "input for Twitter consumer secret")
	flag.StringVar(&opts.TAToken, "tatoken", opts.TAToken, "input for Twitter access token")
	flag.StringVar(&opts.TASecret, "tasecret", opts.TASecret, "input for Twitter access secret")
	flag.Parse()

	var err error
	if os.Getenv("DEV") == "TRUE" {
		db, err = sql.Open("postgres", "postgres://user:pass@86.87.235.82:8082/twitter?sslmode=disable")
	} else {
		db, err = sql.Open("postgres", "postgres://postgres:postgres@127.0.0.1/twitter?sslmode=disable")
	}

	consumerKey := opts.TCKey
	consumerSecret := opts.TCSecret
	accessToken := opts.TAToken
	accessSecret := opts.TASecret

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter Client
	client := twitter.NewClient(httpClient)

	// Convenience Demux demultiplexed stream messages
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {

		if tweet.Coordinates != nil {
			tx, err := db.Begin()
			output, err := json.Marshal(tweet)
			if err != nil {
				fmt.Printf("error: %v\n", err)
			}

			var t time.Time
			var timeErr error
			t, timeErr = time.Parse(time.RubyDate, tweet.CreatedAt)
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
			dbErr := tx.QueryRow("INSERT INTO data (tweet, text, long, lat, day, daypart) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
				output, tweet.Text, tweet.Coordinates.Coordinates[0], tweet.Coordinates.Coordinates[1], day.String(), daypart).Scan(&id)
			if dbErr != nil {
				fmt.Println(dbErr)
			}
			tx.Commit()

			fmt.Println("=======================================================")
			fmt.Println(tweet.Text)
			fmt.Println(tweet.Coordinates.Coordinates[0])
			fmt.Println(tweet.Coordinates.Coordinates[1])
			fmt.Println("In database as ", id)
		}

	}
	demux.DM = func(dm *twitter.DirectMessage) {
		fmt.Println(dm.SenderID)
	}
	demux.Event = func(event *twitter.Event) {
		fmt.Printf("%#v\n", event)
	}

	fmt.Println("Starting Stream...")

	// FILTER
	filterParams := &twitter.StreamFilterParams{
		//Track: []string{"geocode:52.367424,4.892607,20km"},
		Locations: []string{"4.7685", "52.3216", "5.0173", " 52.4251"},
		//	Locations:     []string{"-74", "40", "-73", "41"},
		StallWarnings: twitter.Bool(true),
	}
	stream, err := client.Streams.Filter(filterParams)
	if err != nil {
		log.Fatal(err)
	}

	// USER (quick test: auth'd user likes a tweet -> event)
	// userParams := &twitter.StreamUserParams{
	// 	StallWarnings: twitter.Bool(true),
	// 	With:          "followings",
	// 	Language:      []string{"en"},
	// }
	// stream, err := client.Streams.User(userParams)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// SAMPLE
	// sampleParams := &twitter.StreamSampleParams{
	// 	StallWarnings: twitter.Bool(true),
	// }
	// stream, err := client.Streams.Sample(sampleParams)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Receive messages until stopped or stream quits
	go demux.HandleChan(stream.Messages)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")
	stream.Stop()
}
