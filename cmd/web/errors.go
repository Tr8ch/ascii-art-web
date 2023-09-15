package main

import (
	"html/template"
	"net/http"
)

func Error(w http.ResponseWriter, code int) {
	ErrStruct := struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	}{
		Message: http.StatusText(code),
		Code:    code,
	}
	w.WriteHeader(ErrStruct.Code)
	notFound, err := template.ParseFiles("./static/templates/err.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	notFound.Execute(w, ErrStruct)
}
