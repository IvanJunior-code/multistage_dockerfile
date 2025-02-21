package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Página de teste!</h1>")
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("Servidor em execução.")
	http.ListenAndServe(":8080", nil)
}