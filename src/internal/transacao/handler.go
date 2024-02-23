package transacao

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) SetRoutes(router *mux.Router) {
	router.HandleFunc("/clientes/{id}/transacoes", h.TransacoesDoCliente).Methods("POST")

}

func (h *Handler) TransacoesDoCliente(w http.ResponseWriter, r *http.Request) {
	log.Println("/clientes/{id}/transacoes endpoint called")
	vars := mux.Vars(r)

	fmt.Fprintf(w, "hello\n")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var t Transacao
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	log.Println(t.Valor)
	log.Println(vars["tipo"])
}
