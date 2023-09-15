package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := "3000"
	if val, ok := os.LookupEnv("PORT"); ok {
		port = val
	}

	infoLog := log.New(os.Stdout, "\033[92mINFO\033[0m\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "\033[91mERROR\033[0m\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", MainPageHandler)
	mux.HandleFunc("/download", DownloadHandler)
	mux.HandleFunc("/ascii-art", AsciiArtHandler)
	infoLog.Printf("Starting server on http://localhost:%s", port)
	err := http.ListenAndServe(":"+port, mux)
	errorLog.Fatal(err)
}
