package extrato

import (
	"m/internal/cliente"
	"m/internal/common"
	"m/internal/transacao"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type ExtratoHandler struct {
  clienteRepository common.Repository 
}

func NewHandler(clienteRepository common.Repository) *ExtratoHandler {
	return &ExtratoHandler{clienteRepository: clienteRepository}
}

func (h *ExtratoHandler) SetRoutes(router *mux.Router) {
  router.HandleFunc("/clientes/{id}/extrato", h.ExtratoDoCliente).Methods("GET")
}

func (h *ExtratoHandler) ExtratoDoCliente(w http.ResponseWriter, r *http.Request) {
  idCliente, err := common.GetClientId(w, r)
  client, err := h.clienteRepository.GetById(idCliente)
  if err != nil {
    http.Error(w, err.Error(), http.StatusNotFound)
    return
  }
  cliente := client.(*cliente.Cliente) 


  //TODO: Get client transaction history
  transactions := []transacao.Transacao {}
  
  result := Extrato {
    Saldo: Saldo{ 
      Total: cliente.Saldo,
      Limite: cliente.Limite,
      DataExtrato: time.Now().UTC(),
    },
    UltimasTransacoes: transactions,
  }

  // Return response
  common.WriteJson[Extrato](w, result, http.StatusOK)
}
