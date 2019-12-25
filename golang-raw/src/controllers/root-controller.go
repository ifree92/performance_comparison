package controllers

import (
	"../services"
	"encoding/json"
	"io"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "Golang root")
}

func StringHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	_, _ = io.WriteString(w, "Hello world!")
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {
	jsonText, _ := json.Marshal(services.GetUserServiceInstance().GetUser())
	_, _ = w.Write(jsonText)
}

func JsonArrayHandler(w http.ResponseWriter, r *http.Request) {
	jsonText, _ := json.Marshal(services.GetUserServiceInstance().GetUsers())
	_, _ = w.Write(jsonText)
}
