package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// http://xkcd.com/614/info.0.json

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

// get xkcds and add the to an array of type xkcd
func getXkcds() []XKCD {
	var xkcds []XKCD
	var xkcd XKCD

	for i := 1; i < 10; i++ {
		count := strconv.Itoa(i)

		response, err := http.Get(fmt.Sprintf("http://xkcd.com/%s/info.0.json", count))

		if err != nil {
			fmt.Println(err)
		}

		bytes, _ := ioutil.ReadAll(response.Body)

		response.Body.Close()

		json.Unmarshal(bytes, &xkcd)

		if err != nil {
			fmt.Println(err)
		}
		xkcds = append(xkcds, xkcd)

		response.Body.Close()

		// time.Sleep(3 * time.Second)
	}

	return xkcds

}

func main() {

	var words [2]string
	words[0] = "Kevin"
	words[1] = "Muhia"

	primes := [6]int{2, 3, 5, 7, 11, 13}

	fmt.Printf("Array of  strings : %v \n", words)
	fmt.Printf("Array of integers :  %v \n", primes)
	fmt.Println(len(getXkcds()))

}
