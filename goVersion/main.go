package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var (
	BaseHTMLpath string = filepath.Join("templates", "BASE.html")
)

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	indexPage, err := template.ParseFiles(BaseHTMLpath, filepath.Join("templates", "index.html"))
	if err != nil {
		log.Println("Error parsing files (/)")
	}
	err = indexPage.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Println("Error executing template (/)")
	}
}

func AboutHandler(w http.ResponseWriter, req *http.Request) {
	aboutPage, err := template.ParseFiles(BaseHTMLpath, filepath.Join("templates", "about.html"))
	if err != nil {
		log.Println("Error parsing files (/about)")
	}
	err = aboutPage.ExecuteTemplate(w, "about.html", nil)
	if err != nil {
		log.Println("Error executing template (/about)")
	}
}

func main() {
	// LOGS
	err := os.MkdirAll("logs", os.ModePerm)
	logfile, err := os.Create(filepath.Join("logs", fmt.Sprintf("server_logs_%v.log", time.Now().Unix())))
	if err != nil {
		fmt.Println("Could not create logfile. Panicking...")
		panic(err)
	}
	log.SetOutput(logfile)

	// ROUTES
	serveMux := http.NewServeMux()
	serveMux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	serveMux.HandleFunc("/", IndexHandler)
	serveMux.HandleFunc("/about", AboutHandler)

	// SERVER
	server := &http.Server{
		Addr:           ":8000",
		Handler:        serveMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	serverError := server.ListenAndServe()
	if serverError != nil {
		log.Fatal("Fatal server error ! : ", serverError)
	}
}
