package main

import (
	"../controllers"
	"../services"
	"log"
	"net/http"
)

func main() {
	services.GetUserServiceInstance().Init()
	http.HandleFunc("/", controllers.IndexHandler)
	http.HandleFunc("/string", controllers.StringHandler)
	http.HandleFunc("/json", controllers.JsonHandler)
	http.HandleFunc("/json-array", controllers.JsonArrayHandler)
	log.Println("Starting listening :4444")
	log.Fatal(http.ListenAndServe(":4444", nil))
}
