package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml")

type SitemapIndex struct{
	Location []Location `xml:"sitemap"`
}

type Location struct{
	Loc string `xml:"loc"`
}

/* NOT 100% SURE HOW THIS WORKS BUT IT 
TELLS HOW TO REPRESENT THE LOCATION OJECT
AS A STRING */
func (l Location) String() string{
	return fmt.Sprintf(l.Loc)
}

func main(){
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	// var stringBody string = string(bytes)
	// fmt.Println(stringBody)
	resp.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Location)
}