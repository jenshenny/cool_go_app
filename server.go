package main

import (
	"fmt" //print useful data
	"log" //print fatal errors
	"net/http" //functionality for creating HTTP client/server
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)

	//HandleFunc add route handlers to the web server
	//second argument accepts a function that holds the business logic to correctly respond to the request.
	//this function accepts a ResponseWriter to send a response back and a Request object that provides more information about the request itself
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")

	//ListenAndServe allows us to start the web server and specify the port to listen for incoming requests
	//port parameter needs to be passed as a string prepended by colon punctuation.
	//second parameter accepts a handler to configure the server for HTTP/2.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
			http.Error(w, "404 not found.", http.StatusNotFound)
			return
	}

	if r.Method != "GET" {
			http.Error(w, "Method is not supported.", http.StatusNotFound)
			return
	}


	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}
