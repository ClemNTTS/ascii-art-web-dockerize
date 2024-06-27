package main

import (
	ascii "ascii-art-web/asciiart"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	// Load CSS
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", HandleHome)

	// Initiate the server
	server := &http.Server{
		Addr:              ":3000",           // Address of the server (port is for example)
		ReadHeaderTimeout: 10 * time.Second,  // Time allowed to read headers
		WriteTimeout:      10 * time.Second,  // Max time to write response
		IdleTimeout:       120 * time.Second, // Max time between two requests
		MaxHeaderBytes:    1 << 20,           // 1 MB, maximum bytes server will read
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

type PageData struct {
	Resultat string
	Title    string
}

var (
	tpl, _    = template.ParseFiles("body.html")
	err404, _ = template.ParseFiles("err404.html")
	data      = PageData{
		Resultat: "",
		Title:    "Ascii-Art",
	}
)

// Handler of the main page
func HandleHome(w http.ResponseWriter, r *http.Request) {
	// Management of 404 error
	if r.URL.Path != "/" {
		err := err404.Execute(w, nil)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	//check if we are in a post methode
	if r.Method == http.MethodPost {
		// Get the user input
		opt := r.FormValue("options")
		result := r.FormValue("send_value")
		if result != "" {
			// Create the result in Ascii-art
			tmp, err := ascii.AsciiArt(result, opt)
			if err != nil {
				data.Resultat = "Error 500 : Wrong input !"
				log.Printf("Error %v", http.StatusInternalServerError)
			} else {
				data.Resultat = tmp
			}
		}
	}

	// Run the body.html template with data
	err := tpl.Execute(w, data)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	} else {
		log.Printf("Status OK : %v", http.StatusOK)
	}
}
