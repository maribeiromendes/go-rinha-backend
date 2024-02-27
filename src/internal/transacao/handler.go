package transacao

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

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
  // Validate request
  var t Transacao
  if !ValidateBody(w, r, &t) {
    return
  }

  cliente := GetClient(w, r)
  if cliente == nil {
    http.Error(w, "usuario nao encontrado", http.StatusNotFound)
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
    http.Error(w, "sem limte", http.StatusUnprocessableEntity)
    return
  }

  // TODO: Update balance

  // Return response
	result := RetornoTransacao{ Limite: 1000, Saldo: novoSaldo }
  WriteJson[RetornoTransacao](w, result)
}

func WriteJson[T any](w http.ResponseWriter, content T) {
	w.Header().Set("Content-Type", "application/json")

	jsonResponse, _ := json.Marshal(content)
	w.Write(jsonResponse)
}

func ValidateBody[T any](w http.ResponseWriter, r *http.Request, out *T) bool {
  if r.Body == nil {
    http.Error(w, "requisicao invalida", http.StatusBadRequest)
    return false
  }
  decoder := json.NewDecoder(r.Body)

  err := decoder.Decode(&out)
  if err != nil {
    http.Error(w, "requisicao invalida", http.StatusBadRequest)
    log.Printf(err.Error())
    return false
  }

  return true
}

func GetClient(w http.ResponseWriter, request *http.Request) *Cliente {
  // Mocks a database, must be replaced by repository method later
	vars := mux.Vars(request)
  idCliente, err := strconv.Atoi(vars["id"])
  if err != nil || (idCliente < 0 || idCliente > 5) {
    return nil
  }
  return &Cliente { idCliente, 10000, 0}
}
