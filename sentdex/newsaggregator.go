package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"
)

type NewsMap struct {
	Keyword  string
	Location string
}
type SiteMapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Loc        []string `xml:"url>loc"`     // locations
	Lastmod    []string `xml:"url>lastmod"` // keywords
	Changefreq []string `xml:"url>changefreq"`
	Priority   []string `xml:"url>priority"` // title
}

type NewsAggPage struct {
	Title string
	News  map[string]NewsMap
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	var s SiteMapIndex
	var n News
	newsMap := make(map[string]NewsMap)
	resp, _ := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")
	// defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)

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
		bytes, _ := ioutil.ReadAll(response.Body)
		xml.Unmarshal(bytes, &n)

		for index, _ := range n.Priority {
			newsMap[n.Priority[index]] = NewsMap{n.Priority[index], n.Loc[index]}
		}
	}
	p := NewsAggPage{Title: "Amazing news Aggregator", News: newsMap}
	t, _ := template.ParseFiles("basictemplating.html")
	t.Execute(w, p)
}
func main() {
	http.HandleFunc("/agg", newsAggHandler)
	http.ListenAndServe(":5000", nil)
}
