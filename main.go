package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const HackerNewsURL string = "https://hacker-news.firebaseio.com/v0/"

const ListStoriesURL string = HackerNewsURL + "askstories.json"
const GetStoryURL string = HackerNewsURL + "item/%v.json"

type Story struct {
	Title string
	Text  string
	Link string
}


func httpGet(url string) []byte {
	client := &http.Client{Timeout: 10 * time.Second}

	res, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	return body
}

func httpGetAsync(url string, ch chan<- []byte) {
	client := &http.Client{Timeout: 10 * time.Second}

	res, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	ch <- body
}

func main() {
	stories := httpGet(ListStoriesURL)
	var storyIds []int
	json.Unmarshal(stories, &storyIds)
	ch := make(chan []byte)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	for _, story := range storyIds {
		storyURL := fmt.Sprintf(GetStoryURL, story)
		go httpGetAsync(storyURL, ch)
		
		
	}

	for range storyIds{
		var story Story
		json.Unmarshal(<-ch, &story)
		fmt.Fprintln(w, story.Title, story.Link, r.URL.Path)

	}
})

	http.ListenAndServe(":80", nil)

}
