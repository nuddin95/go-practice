package main

import (
	"fmt"
	"net/http")
	// "github.com/mmcdole/gofeed")

func index_handler(w http.ResponseWriter, r *http.Request){
	//Fprintf first arguement is what to write to in
	//this case teh ResponseWriter
	
	fmt.Fprintf(w, "Go is pretty cool")
}

// func rss(w http.ResponseWriter, r *http.Request){
// 	fp := gofeed.NewParser()
// 	feed, _ := fp.ParseURL("https://www.androidpolice.com/feed/")
// 	fmt.Println(feed)
// 	fmt.Fprintf(w, feed.Title)
// }

func main(){
	//specifies the path and what function to run
	http.HandleFunc("/", index_handler)

	//rss feed
	// http.HandleFunc("/rss", rss)

	//creates the server
	http.ListenAndServe(":8000", nil)
}