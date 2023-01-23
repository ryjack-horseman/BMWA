package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "hello world")
		fmt.Printf("Number of Bytes written: %d\n", n)
		if err != nil {
			fmt.Println(err)
		}
	})

	http.ListenAndServe(":8080", nil)
}
