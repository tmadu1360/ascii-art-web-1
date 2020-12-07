package main

import (
	"fmt"
	"log"
	"net/http"

	asciiart "./Ascii"
)

func startPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.ServeFile(w, r, "www/error404.html")
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "www/")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		text := r.FormValue("text")
		font := r.FormValue("style")
		out := asciiart.AsciiArt(text, font)
		fmt.Fprintf(w, out)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	//http.Handle("/", http.FileServer(http.Dir("www/")))
	fs := http.FileServer(http.Dir("www/css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	police := http.FileServer(http.Dir("www/font"))
	http.Handle("/font/", http.StripPrefix("/font/", police))
	img := http.FileServer(http.Dir("www/image"))
	http.Handle("/image/", http.StripPrefix("/image/", img))

	http.HandleFunc("/", startPage)

	fmt.Printf("Starting server for testing HTTP on port 8080...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
