package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, "Listening on :%s\n", "8080")
	hostname, _ := os.Hostname()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "I'm %s\n", hostname)
		fmt.Fprintf(w, "I'm %s\n", hostname)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
