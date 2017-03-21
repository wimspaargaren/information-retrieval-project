package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	resp, err := http.Get("https://en.wikipedia.org/wiki/List_of_dance_style_categories")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	r, errorrr := regexp.Compile("<li>(.*?)</li>")
	if errorrr != nil {
		fmt.Println(errorrr)
		return
	}
	b := string(body)
	res := r.FindAllString(b, -1)
	var titles string
	for i := 0; i < len(res); i++ {
		result := strings.Split(res[i], ">")
		title := result[2]
		// fmt.Println(title)
		if title != "" {
			titles = titles + title[:len(title)-3] + ","
		}
	}
	fmt.Println(titles)
}
