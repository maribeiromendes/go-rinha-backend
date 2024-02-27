package transacao

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type TransacaoHandler struct{}

func NewHandler() *TransacaoHandler {
	return &TransacaoHandler{}
}

func (h *TransacaoHandler) SetRoutes(router *mux.Router) {
	router.HandleFunc("/clientes/{id}/transacoes", h.TransacoesDoCliente).Methods("POST")

}

func (h *TransacaoHandler) TransacoesDoCliente(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "request must contain a body", http.StatusBadRequest)
		return
	}

	//vars := mux.Vars(r)

	decoder := json.NewDecoder(r.Body)

	var t Transacao
	err := decoder.Decode(&t)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		log.Printf(err.Error())
		return
	}
	//log.Println(t.Valor)
	//log.Println(vars["tipo"])

	result := RetornoTransacao{Limite: 10, Saldo: 5}
	jsonResponse, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

}
