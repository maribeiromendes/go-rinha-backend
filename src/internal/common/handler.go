package common

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler interface {
  SetRoutes(router *mux.Router)
}

type ErrorResponse struct {
	Message string `json:"mensagem"`
}

func WriteJson[T any](w http.ResponseWriter, content T, status  int) {
  w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	jsonResponse, _ := json.Marshal(content)
	w.Write(jsonResponse)
}

func WriteErrorJson(w http.ResponseWriter, message string, status int) {
  content := ErrorResponse {Message: message}
  WriteJson(w, content, status)
}

func ValidateBody[T any](w http.ResponseWriter, r *http.Request, out *T) bool {
  if r.Body == nil {
    WriteErrorJson(w, "requisicao invalida", http.StatusBadRequest)
    return false
  }
  decoder := json.NewDecoder(r.Body)

  err := decoder.Decode(&out)
  if err != nil {
    WriteErrorJson(w, "requisicao invalida", http.StatusBadRequest)
    log.Printf(err.Error())
    return false
  }
  return true
}

func GetClientId(w http.ResponseWriter, request *http.Request) (int, error) {
	vars := mux.Vars(request)
  idStr := vars["id"]
  idCliente, err := strconv.Atoi(idStr)
  if err != nil {
    return 0, fmt.Errorf("id: '%s' invalido", idStr) 
  }

  return idCliente, nil
}
