package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// This is here so we can actually see that the responses that have been cached don't actually get here
		fmt.Println("The request actually got here")

		w.Write([]byte("You got here"))
	})

	http.ListenAndServe(":8000", mux)
}
