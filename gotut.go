package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"
)

// type SitemapIndex struct{
// 	Locations[]Location `xml:"sitemap"`
// }

// type Location struct{
// 	Loc string `xml:"loc"`
// }

// func (l Location) String() string{
// 	return fmt.Sprintf(l.Loc)
// }

type SitemapIndex struct{
	Locations []string `xml:"sitemap>loc"`
}

type News struct{
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>Keywords"`
	Locations []string`xml:"url>location"`
}

func main() {
	var s SitemapIndex
	var n News
	//basic get request
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _:= ioutil.ReadAll(resp.Body)

	//free up the resources that grabbed the information
	resp.Body.Close()


	// var s SitemapIndex
	xml.Unmarshal(bytes, &s)

	// fmt.Println(s.Locations)

	//loop over sitemaps and display them
	// for _, Location := range s.Locations{
	// 	fmt.Printf("\n %s", Location)
	// }

	for _, Location := range s.Locations{
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
		fmt.Printf("\n %s", Location)
	}

}
