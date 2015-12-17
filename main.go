package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Printf("server error: %s\n", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("---***### Incoming request ###***---\n")
	fmt.Printf("Request path: %v%v\n", r.Host, r.URL.Path)
	fmt.Println("Headers:")
	for k, v := range r.Header {
		fmt.Printf("    %v:\n", k)
		for _, hv := range v {
			fmt.Printf("          %v\n", hv)
		}
	}
	fmt.Println("\n\n")
	w.WriteHeader(http.StatusNoContent)
}
