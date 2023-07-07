package main

import (
	"fmt"
	"net/http"
)

func handle(response http.ResponseWriter, request *http.Request){
	fmt.Fprintf(response, "Server is started successfully")
}

func main(){
	http.HandleFunc("/", handle)
	http.ListenAndServe(":5000", nil)
}

//http://localhost:5000