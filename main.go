// xkcd is a tool that dowloads each URL of each comic once, creates an offline index,
// then using that index prints the URL and the transcript of each comic that matches
// a search term provided on the command line.
// Made with help from the Golang Community, built based off an exercise in GoPL textbook
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const xkcdURL = "https://xkcd.com/"

type xkcdComic struct {
	Transcript string `json:"transcript"`
	ComicNum   int    `json:"num"`
}

func main() {
	result, err := populateComics()
	if err != nil {
		log.Fatal(err)
	}
	searchXkcd(result, os.Args[1:])
}

func populateComics() ([]xkcdComic, error) {
	var comics []xkcdComic
	var u string
	for i := 1; i <= 1626; i++ {
		u = xkcdURL + strconv.Itoa(i) + "/info.0.json"
		resp, err := http.Get(u)
		if err != nil {
			return nil, err
		}
		var c xkcdComic
		switch resp.StatusCode {
		case 200:
			err := json.NewDecoder(resp.Body).Decode(&c)
			resp.Body.Close()
			if err != nil {
				return nil, err
			}
			comics = append(comics, c)
		default:
			resp.Body.Close()
			fmt.Println(u)
		}
	}
	return comics, nil
}

func searchXkcd(comics []xkcdComic, terms []string) {
	for _, c := range comics {
		for _, t := range terms {
			if strconv.Itoa(c.ComicNum) == t {
				fmt.Printf("URL: https://xkcd.com/ \b%v/info.0.json,\nTranscript: %s\n\n", c.ComicNum, c.Transcript)
			}
		}
	}
}

// STILL TO ADD: make populateComics more simplified with defer, using simplified function
// func populateSingleComic() xkcdComic {

//}
