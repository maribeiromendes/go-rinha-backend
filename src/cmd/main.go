package main

import (
	"fmt"
	"log"
	"m/internal/middlewares"
	"m/internal/transacao"
	"net/http"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("/hello endpoint called")
	fmt.Fprintf(w, "hello\n")

}

func main() {
	r := mux.NewRouter()
	log.Println("Registering routes")
	r.HandleFunc("/", hello)
  r.Use(middlewares.Logging)
	transacao_handler := transacao.NewHandler()
	transacao_handler.SetRoutes(r)

	port := 8000
	log.Println(fmt.Sprintf("Server up and listening on port: %d...", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
