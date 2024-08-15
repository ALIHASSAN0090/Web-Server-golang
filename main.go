package main

import (
	"fmt"
	"log"
	"net/http"

	"server/controllers"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", controllers.HelloHandler)
	http.HandleFunc("/form", controllers.FormHandler)

	fmt.Println("Creating Server")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Print(err)
	}
}
