package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1> Welcome to the Homepage! </h1>")
}

func aboutUs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1> About us: </h1>")
	fmt.Fprintln(w, "<p> We are amazing! </p>")
}

func products(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1> Products </h1>")
	fmt.Fprintln(w, "<h3> Coming soon... </h3>")
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", aboutUs)
	http.HandleFunc("/products", products)

	fmt.Println("Server running!")

	http.ListenAndServe(":7777", nil)
}
