package main

import (
	//"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml")

type SitemapIndex struct{
	Location []string `xml:"sitemap>loc"`
}

type News struct{
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

/* NOT 100% SURE HOW THIS WORKS BUT IT 
TELLS HOW TO REPRESENT THE LOCATION OJECT
AS A STRING */
// func (l Location) String() string{
// 	return fmt.Sprintf(l.Loc)
// }

func main(){
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	// var stringBody string = string(bytes)
	// fmt.Println(stringBody)
	resp.Body.Close()

	var n News
	var s SitemapIndex
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Location{
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
	}

}