package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"m/internal/transacao"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("/hello endpoint called")
	fmt.Fprintf(w, "hello\n")
}

func main() {
	r := mux.NewRouter()
	log.Println("Registering routes")
	r.HandleFunc("/", hello)
	transacao_handler := transacao.NewHandler()
	transacao_handler.SetRoutes(r)

	port := 8000
	log.Println(fmt.Sprintf("Server up and listening on port: %d...", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
