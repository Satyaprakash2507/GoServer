package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!!!!!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) //it will look for static file and by default pick the index.html
	http.Handle("/", fileServer)                        //handle root route which is just above
	http.HandleFunc("/form", formHandler)               // handle formhandler
	http.HandleFunc("/hello", helloHandler)             //handle hellohandler

	fmt.Printf("Starting server at port 8080 \n")
	//starting server on 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
