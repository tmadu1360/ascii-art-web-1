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

//hey coucou

func startPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.ServeFile(w, r, "www/error404.html")
		return
	}

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
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		text := r.FormValue("text")
		font := r.FormValue("style")
		//out := asciiart.AsciiArt(text, font)
		result := struct {
			Data string
		}{
			Data: asciiart.AsciiArt(text, font),
		}
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
			err = parsedTemplate.Execute(w, result)
			if err != nil {
				log.Println("Error executing template :", err)
				return
			}
		}
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
	js := http.FileServer(http.Dir("www/js"))
	http.Handle("/js/", http.StripPrefix("/js/", js))

	http.HandleFunc("/", startPage)

	fmt.Printf("Starting server for testing HTTP on port 8080...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
