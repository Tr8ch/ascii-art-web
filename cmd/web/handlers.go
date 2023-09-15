package main

import (
	"ascii-art-web/internal"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Error(w, http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		mainPage, err := template.ParseFiles("./static/templates/main_page.html")
		if err != nil {
			Error(w, http.StatusInternalServerError)
			return
		}
		err = mainPage.Execute(w, nil)
		if err != nil {
			Error(w, http.StatusInternalServerError)
			return
		}
	default:
		Error(w, http.StatusMethodNotAllowed)
		return
	}
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	Request := struct {
		SimpleText string `json:"text"`
		Font       string `json:"font"`
	}{}

	type response struct {
		AsciiText string `json:"asciiText"`
	}

	if http.MethodPost != r.Method {
		Error(w, http.StatusMethodNotAllowed)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		Error(w, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(data, &Request)
	if err != nil {
		Error(w, http.StatusBadRequest)
		return
	}

	asciiMaker := internal.NewAsciiArt(Request.Font)
	asciiText, errNum := asciiMaker.AsciiMake(Request.SimpleText)
	if errNum == 400 {
		badreq := struct {
			Message string `json:"message"`
		}{
			Message: "invalid Input: Please enter valid ASCII characters only",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(errNum)
		err := json.NewEncoder(w).Encode(badreq)
		if err != nil {
			Error(w, http.StatusInternalServerError)
			return
		}
		return
	}
	if errNum == 500 {
		Error(w, http.StatusInternalServerError)
		return
	}

	Response := response{
		AsciiText: asciiText,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(Response)
	if err != nil {
		Error(w, http.StatusInternalServerError)
		return
	}
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Error(w, 405)
		return
	}

	err := r.ParseForm()
	if err != nil {
		Error(w, 500)
	}
	data := r.FormValue("downloadText")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii.txt")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprint(len(data)))

	w.Write([]byte(data))
}
