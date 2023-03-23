package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, world, this is environment %s", os.Getenv("environment"))
	})

	fmt.Println("Server started on port 8080")
	http.ListenAndServe("0.0.0.0:8080", nil)
}
