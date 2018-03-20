package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"
)

type SitemapIndex struct{
	Locations[]Location `xml:"sitemap"`
}

type Location struct{
	Loc string `xml:"loc"`
}

func (l Location) String() string{
	return fmt.Sprintf(l.Loc)
}

func main() {
	//basic get request
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _:= ioutil.ReadAll(resp.Body)

	//free up the resources that grabbed the information
	resp.Body.Close()


	var s SitemapIndex
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Locations)

}
