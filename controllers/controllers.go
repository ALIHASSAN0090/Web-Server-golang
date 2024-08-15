package controllers

import (
	"fmt"
	"net/http"
)

func HelloHandler(res http.ResponseWriter, req *http.Request) {
	// Check if the URL path is correct
	if req.URL.Path != "/hello" {
		http.Error(res, "404 not found", http.StatusNotFound)
		return
	}

	// Check if the HTTP method is GET
	if req.Method != "GET" {
		http.Error(res, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	// Write a simple "HELLO" response
	fmt.Fprintf(res, "HELLO")
}

func FormHandler(res http.ResponseWriter, req *http.Request) {
	// Check if the URL path is correct
	if req.URL.Path != "/form" {
		http.Error(res, "404 not found", http.StatusNotFound)
		return
	}

	// Ensure the method is POST
	if req.Method != "POST" {
		http.Error(res, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	if err := req.ParseForm(); err != nil {
		http.Error(res, fmt.Sprintf("Parse Form Error: %v", err), http.StatusInternalServerError)
		return
	}

	// Log that the post request was successful
	fmt.Println("Post requested successfully")

	// Extract the form values
	name := req.FormValue("name")
	address := req.FormValue("address")

	// Respond with the submitted name and address
	fmt.Fprintf(res, "Name is: %s\n", name)
	fmt.Fprintf(res, "Address is: %s\n", address)
}
