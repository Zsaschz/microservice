package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Helo World")
		d, err := ioutil.ReadAll(r.Body)

		if (err != nil) {
			http.Error(w, "Error", http.StatusBadRequest)
		}

		fmt.Fprintf(w, "Helloww %s", d)
	})

	http.ListenAndServe(":8080", nil)
}