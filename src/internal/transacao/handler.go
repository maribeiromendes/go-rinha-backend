package transacao

import (
	"m/internal/cliente"
	"m/internal/common"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)


type TransacaoHandler struct {
  clienteRepository common.Repository 
}

func NewHandler(clienteRepository common.Repository) *TransacaoHandler {
	return &TransacaoHandler{clienteRepository: clienteRepository}
}

func (h *TransacaoHandler) SetRoutes(router *mux.Router) {
  router.HandleFunc("/clientes/{id}/transacoes", h.TransacoesDoCliente).Methods("POST")
}

func (h *TransacaoHandler) TransacoesDoCliente(w http.ResponseWriter, r *http.Request) {
  // Validate request
  var t Transacao
  if !common.ValidateBody(w, r, &t) {
    return
  }

  idCliente, err := common.GetClientId(w, r)
  client, err := h.clienteRepository.GetById(idCliente)
  if err != nil {
    http.Error(w, err.Error(), http.StatusNotFound)
    return
  }
  cliente := client.(*cliente.Cliente) 

  // Validate limit
  var valor int
  if strings.ToLower(t.Tipo) == "c" {
    valor = t.Valor
  } else {
    valor = -t.Valor
  }
  novoSaldo := cliente.Saldo + valor
  if novoSaldo < - cliente.Limite {
    common.WriteErrorJson(w, "sem limite", http.StatusUnprocessableEntity)
    return
  }

  // TODO: Update balance

  // Return response
	result := RetornoTransacao{ Limite: cliente.Limite, Saldo: novoSaldo }
  common.WriteJson[RetornoTransacao](w, result, http.StatusOK)
}
