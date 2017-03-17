package main

import (
	sql "database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {

	var err error
	db, err = sql.Open("postgres", "postgres://user:pass@86.87.235.82:8082/twitter?sslmode=disable")

	rows, err := db.Query("select id, tweet from data where id < 100")
	if err != nil {
		fmt.Println("Could not retrieve EANs from user")
	}

	tx, err := db.Begin()
	for rows.Next() {
		// var tweet twitter.Tweet
		var i int
		err = rows.Scan(&i)
		if err != nil {
			fmt.Println("this did not work")
		}

		tweet := twitter.Tweet{}
		err := json.Unmarshal(data, &tweet)
		if err != nil {
			fmt.Println("unmarshalling error", err)
		}
		var t time.Time
		var timeErr error
		const longForm = "Jan 2 12:00:04 +0000 2006"
		t, timeErr = time.Parse(time.RubyDate, tweet.CreatedAt)
		if timeErr != nil {
			fmt.Println(timeErr)
		}
		x := t.Weekday()
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
		// fmt.Println(tweet.Coordinates.Coordinates[1], tweet.Coordinates.Coordinates[0])
		var id int
		dbErr := tx.QueryRow("UPDATE data SET text=$1, lat=$3, long=$4, day=$5, daypart=$6 WHERE id=$2 RETURNING id", tweet.Text, i, tweet.Coordinates.Coordinates[1], tweet.Coordinates.Coordinates[0], x.String(), daypart).Scan(&id)
		if dbErr != nil {
			fmt.Println(dbErr, "Error for tweet ", id)
		}
	}
	tx.Commit()
}
