package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ingojaeckel/golang-packages/fib"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Running")
	})
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		n, _ := strconv.ParseInt(r.URL.Query().Get("n"), 10, 64)
		fmt.Fprintf(w, "fib(%d)=%d", n, fib.Fib(n))
	})
	fmt.Println("Running..")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
