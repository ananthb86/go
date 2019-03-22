package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//REST endpoint
const HackerNewsURL string = "https://hacker-news.firebaseio.com/v0/"

const ListStoriesURL string = HackerNewsURL + "askstories.json"
const GetStoryURL string = HackerNewsURL + "item/%v.json"

type Story struct {
	Title string
	Text  string
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

func main() {

	stories := httpGet(ListStoriesURL)
	var storyIds []int

	json.Unmarshal(stories, &storyIds)

	for i, story := range storyIds {
		storyURL := fmt.Sprintf(GetStoryURL, story)
		storyStr := httpGet(storyURL)
		var story Story
		json.Unmarshal(storyStr, &story)

		fmt.Println(i+1, story.Title)
		//fmt.Println(story.Text)
	}

}
