package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

func main() {
	color.Red("Merhaba Dünya!")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Merhaba Dünya!")
}
