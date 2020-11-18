package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// https://www.admfactory.com/how-to-trim-whitespace-from-string-in-golang/

type SiteMapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Loc        []string `xml:"url>loc"`     // locations
	Lastmod    []string `xml:"url>lastmod"` // keywords
	Changefreq []string `xml:"url>changefreq"`
	Priority   []string `xml:"url>priority"` // title
}
type NewsMap struct {
	Keyword  string
	Location string
}

func main() {
	var s SiteMapIndex
	var n News
	newsMap := make(map[string]NewsMap)
	resp, _ := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")

	// defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)

	xml.Unmarshal(bytes, &s)

	// loop over a slice
	for _, Location := range s.Locations {
		strTrimSpace := strings.TrimSpace(Location) // remove whitespaces at begining and end of the url

		// u, err := url.Parse(strTrimSpace)
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Println(u)

		response, error := http.Get(strTrimSpace)
		if error != nil {
			fmt.Println(error)
		}
		// defer response.Body.Close()

		bytes, _ := ioutil.ReadAll(response.Body)
		xml.Unmarshal(bytes, &n)

		for index, _ := range n.Priority {
			newsMap[n.Priority[index]] = NewsMap{n.Priority[index], n.Loc[index]}
		}

		for index, value := range newsMap {
			fmt.Println("\n\n\n", index)
			fmt.Println("\n\n\n", value.Keyword)
			fmt.Println("\n\n\n", value.Location)
		}

	}
}
