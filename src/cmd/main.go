package main

import (
	"fmt"
	"log"
	"m/internal/cliente"
	"m/internal/extrato"
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
  r.Use(middlewares.LoggingMiddleware)
  
  clientRepo := cliente.ClienteRepository{}
	transacao_handler := transacao.NewHandler(&clientRepo)
  extrato_handler := extrato.NewHandler(&clientRepo)

  // Set routes
	transacao_handler.SetRoutes(r)
  extrato_handler.SetRoutes(r)

	port := 8000
	log.Println(fmt.Sprintf("Server up and listening on port: %d...", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
