package main

import (
	"fmt"
	"log"
	"net/http"
)

func abc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 NOT FOUND", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "forms.html")

	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err : %v", err)
			return
		}

		fmt.Fprintf(w, "Post from website r.postfrom = %v\n", r.PostForm)
		name := r.FormValue("name")
		address := r.FormValue("address")

		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)

	default:
		fmt.Fprintf(w, "Only Get and Post")
	}
}

func main() {
	http.HandleFunc("/", abc)

	fmt.Printf("Starting server \n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
