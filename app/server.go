package main

import (
	"fmt"
	"net/http"
	"encoding/json")
	// "github.com/mmcdole/gofeed")

func index_handler(w http.ResponseWriter, r *http.Request){
	//Fprintf first arguement is what to write to in
	//this case teh ResponseWriter
	
	//NOTE: EVERYTHING MUST BE HTML IF YOU WANT TO RETURN HTML
	fmt.Fprintf(w, "<h1>Go is pretty cool<h1>")
	fmt.Fprintf(w, "<h6>You can even use %s and html tags like %s </h6>", "variables", "<strong> this </strong>")
}

// func rss(w http.ResponseWriter, r *http.Request){
// 	fp := gofeed.NewParser()
// 	feed, _ := fp.ParseURL("https://www.androidpolice.com/feed/")
// 	fmt.Println(feed)
// 	fmt.Fprintf(w, feed.Title)
// }

func data(w http.ResponseWriter, r *http.Request){
	people := [6]int{2, 3, 5, 7, 11, 13}
	json.NewEncoder(w).Encode(people)
}

func main(){
	//specifies the path and what function to run
	http.HandleFunc("/", index_handler)

	//rss feed
	// http.HandleFunc("/rss", rss)


	http.HandleFunc("/data", data)
	//creates the server
	http.ListenAndServe(":8000", nil)
}