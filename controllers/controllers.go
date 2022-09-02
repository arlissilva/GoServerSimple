package controllers

import (
	"fmt"
	"net/http"
)

func RequestController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		fmt.Fprintf(w, "Metodo POST.")
	case "GET":
		fmt.Fprintf(w, "Metodo GET.")
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
