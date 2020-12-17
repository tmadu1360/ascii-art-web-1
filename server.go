package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"

	asciiart "./Ascii"
)

//Main Page function
func startPage(w http.ResponseWriter, r *http.Request) {
	//Error404 if you try anything else than "/" path
	if r.URL.Path != "/" {
		http.ServeFile(w, r, "www/error404.html")
		return
	}

	//Handle GET Methode, serve the page and set {{.Data}} to blank.
	switch r.Method {
	case "GET":
		result := struct {
			Data string
		}{
			Data: " ",
		}
		parsedTemplate, _ := template.ParseFiles("www/index.html")
		err := parsedTemplate.Execute(w, result)
		if err != nil {
			log.Println("Error executing template :", err)
			return
		}
	case "POST": //Handle POST Method, serve the template with the data we want.
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		//Get data from the html web page.
		text := r.FormValue("text")
		font := r.FormValue("style")
		//Define the {{.Data}} to put the ascii art in the template
		result := struct {
			Data string
		}{
			Data: asciiart.AsciiArt(text, font),
		}
		//Handle download button
		if r.FormValue("download") == "txt" {
			f, err := os.Open("output.txt")
			if err != nil {
				fmt.Println(err)
			}
			fileinfo, err := f.Stat()
			if err != nil {
				fmt.Println(err)
			}
			size := fileinfo.Size()
			arr := make([]byte, size)
			f.Read(arr)
			file := strings.NewReader(string(arr))
			f.Close()
			w.Header().Set("Content-Disposition", "attachment; filename=file.txt")
			w.Header().Set("Content-Type", "text")
			w.Header().Set("Content-Length", strconv.Itoa(len(arr)))
			io.Copy(w, file)
		} else {
			file, err := os.Create("output.txt")
			fmt.Fprintf(file, result.Data)
			if err != nil {
				fmt.Println(err)
			}
			file.Close()
			parsedTemplate, _ := template.ParseFiles("www/index.html")
			err = parsedTemplate.Execute(w, result) //execute the template with the data.
			if err != nil {
				log.Println("Error executing template :", err)
				return
			}
		}
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.") //Error if there's any other method than GET | POST used
	}
}

func main() {
	//Handle all of the files, css js images fonts and html
	fs := http.FileServer(http.Dir("www/css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	police := http.FileServer(http.Dir("www/font"))
	http.Handle("/font/", http.StripPrefix("/font/", police))
	img := http.FileServer(http.Dir("www/image"))
	http.Handle("/image/", http.StripPrefix("/image/", img))
	js := http.FileServer(http.Dir("www/js"))
	http.Handle("/js/", http.StripPrefix("/js/", js))

	http.HandleFunc("/", startPage) //handle the main page

	fmt.Printf("Starting server for testing HTTP on port 8080... (go to localhost:8080 on your web browser)\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
