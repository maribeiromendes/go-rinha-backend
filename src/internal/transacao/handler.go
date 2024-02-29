package transacao

import (
	"encoding/json"
	"fmt"
	"log"
	"m/internal/cliente"
	"m/internal/common"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type ErrorResponse struct {
	Message string `json:"mensagem"`
}

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
  if !ValidateBody(w, r, &t) {
    return
  }

  cliente, err := h.getClient(w, r)
  if err != nil {
    http.Error(w, err.Error(), http.StatusNotFound)
    return
  }

  // Validate limit
  var valor int
  if strings.ToLower(t.Tipo) == "c" {
    valor = t.Valor
  } else {
    valor = -t.Valor
  }
  novoSaldo := cliente.Saldo + valor
  if novoSaldo < - cliente.Limite {
    writeErrorJson(w, "sem limite", http.StatusUnprocessableEntity)
    return
  }

  // TODO: Update balance

  // Return response
	result := RetornoTransacao{ Limite: cliente.Limite, Saldo: novoSaldo }
  writeJson[RetornoTransacao](w, result, http.StatusOK)
}

func writeJson[T any](w http.ResponseWriter, content T, status  int) {
  w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	jsonResponse, _ := json.Marshal(content)
	w.Write(jsonResponse)
}

func writeErrorJson(w http.ResponseWriter, message string, status int) {
  content := ErrorResponse {Message: message}
  writeJson(w, content, status)
}

func ValidateBody[T any](w http.ResponseWriter, r *http.Request, out *T) bool {
  if r.Body == nil {
    writeErrorJson(w, "requisicao invalida", http.StatusBadRequest)
    return false
  }
  decoder := json.NewDecoder(r.Body)

  err := decoder.Decode(&out)
  if err != nil {
    writeErrorJson(w, "requisicao invalida", http.StatusBadRequest)
    log.Printf(err.Error())
    return false
  }
  return true
}

func (h *TransacaoHandler) getClient(w http.ResponseWriter, request *http.Request) (*cliente.Cliente, error) {
	vars := mux.Vars(request)
  idStr := vars["id"]
  idCliente, err := strconv.Atoi(idStr)
  if err != nil {
    return nil, fmt.Errorf("id: '%s' invalido", idStr) 
  }
  client, err := h.clienteRepository.GetById(idCliente)
  return client.(*cliente.Cliente), err 
}
