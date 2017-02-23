package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

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
			fmt.Println("=======================================================")
			fmt.Println(tweet.Text)
			fmt.Println(tweet.Coordinates.Coordinates[0])
			fmt.Println(tweet.Coordinates.Coordinates[1])
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
