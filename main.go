package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("Server error: %s\n", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("---***### Incoming request ###***---")
	fmt.Println()

	fmt.Printf("Request path: %v%v\n", r.Host, r.URL.Path)
	fmt.Println("Headers:")
	for k, v := range r.Header {
		fmt.Printf("    %v:\n", k)
		for _, hv := range v {
			fmt.Printf("          %v\n", hv)
		}
	}

	// Processing the form
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Failed to parse the form")
	}

	if len(r.Form) > 0 {
		fmt.Println("Fields in the Form:")
		for k := range r.Form {
			fmt.Println(k)
		}
	}

	// Extracting the payload
	payload := r.Form["payload"]

	if payload != nil {
		fmt.Println("Payload:")
		fmt.Printf("%v\n", payload)
	}

	fmt.Println()
	w.WriteHeader(http.StatusOK)
}
