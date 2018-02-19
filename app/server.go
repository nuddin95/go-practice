package main

import (
	"fmt"
	"net/http")

func index_handler(w http.ResponseWriter, r *http.Request){
	//Fprintf first arguement is what to write to in
	//this case teh ResponseWriter
	
	fmt.Fprintf(w, "Go is pretty cool")
}

func main(){
	//specifies the path and what function to run
	http.HandleFunc("/", index_handler)

	//creates the server
	http.ListenAndServe(":8000", nil)
}