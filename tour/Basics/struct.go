package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type XKCD struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

const xkcdURL string = "https://xkcd.com/info.0.json"

func main() {
	var xkcd XKCD

	response, err := http.Get(xkcdURL)
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	bytes, _ := ioutil.ReadAll(response.Body)

	json.Unmarshal(bytes, &xkcd)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(xkcd)
}
