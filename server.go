package main

import (
	"fmt"
	"log"
	"net/http"

	asciiart "./Ascii"
)

func startPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Page not found.", http.StatusNotFound)
	}

	if r.Method != "GET" {
		http.Error(w, "Method not supported !", http.StatusNotFound)
	}

	http.ServeFile(w, r, "start.html")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		text := r.FormValue("text")
		font := r.FormValue("font")
		out := asciiart.AsciiArt(text, font)
		fmt.Fprintf(w, out)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", startPage)
	http.HandleFunc("/form", postHandler)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
