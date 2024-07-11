package main;

import (
	"net/http"
	"log"
	"fmt"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hallo :O")
	log.Printf("Got request on index :O")
}

func main() {
	http.HandleFunc("/", index)
	log.Printf("Running Backend at 8787")
	if err := http.ListenAndServe(":8787", nil); err != nil {
		log.Fatal(err)
	}
}
