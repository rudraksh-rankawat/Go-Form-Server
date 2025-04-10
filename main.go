package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(	w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "world")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "error in Parseform(): %v", err)
		return
	}
	fmt.Fprintf(w, "POST request and form submission successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)

}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Server started at 8080...\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal()
	}
}
