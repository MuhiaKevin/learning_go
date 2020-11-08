package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SiteMapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf((l.Loc))
}

// func index_handler(w http.ResponseWriter, r *http.Request) {

// 	fmt.Fprintf(w, "Woa go is neat")
// }

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)

	var s SiteMapIndex
	xml.Unmarshal(bytes, &s)

	// fmt.Println(s.Locations)

	for _, Location := range s.Locations {
		fmt.Printf("\n%s")
	}

	// http.HandleFunc("/", index_handler)
	// http.ListenAndServe(":8000", nil)
}
