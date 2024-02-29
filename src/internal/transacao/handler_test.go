package transacao

import (
	"bytes"
	"m/internal/cliente"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type FakeClienteRepository struct {}

func (r *FakeClienteRepository) GetById(id int) (interface{}, error) {
  if id == 1 {
    return &cliente.Cliente {Limite: 100, Saldo: 0}, nil
  }

  return nil,nil
}

func TestTransacoesDoClienteWithoutBody(t *testing.T) {
	expectedStatusCode := http.StatusBadRequest
  
  r := createRouter(&FakeClienteRepository{})
	req, _ := http.NewRequest("POST", "/clientes/1/transacoes", nil)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert_equal(rr.Code, expectedStatusCode, t)
}

func TestTransacoesDoClienteWithInvalidBody(t *testing.T) {
	expectedStatusCode := http.StatusBadRequest

  r := createRouter(&FakeClienteRepository{})
	req, _ := http.NewRequest("POST", "/clientes/1/transacoes", bytes.NewBuffer([]byte(`"test": "not valid"`)))

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert_equal(rr.Code, expectedStatusCode, t)
}

func TestCredit(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedBody := `{"limite":100,"saldo":1000}`

  r := createRouter(&FakeClienteRepository{})
	req, _ := http.NewRequest("POST", "/clientes/1/transacoes", bytes.NewBuffer([]byte(`{"valor": 1000, "tipo": "c", "descricao": "descricao"}`)))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert_equal(rr.Code, expectedStatusCode, t)
	assert_equal(rr.Body.String(), expectedBody, t)
}

func TestDebitWhenNoLimit(t *testing.T) {
  expectedStatusCode := http.StatusUnprocessableEntity
  expectedBody := `{"mensagem":"sem limite"}`

  r := createRouter(&FakeClienteRepository{})
  req, _ := http.NewRequest("POST", "/clientes/1/transacoes", bytes.NewBuffer([]byte(`{"valor": 101, "tipo": "d", "descricao": "descricao"}`)))
  rr := httptest.NewRecorder()
  r.ServeHTTP(rr, req)

  assert_equal(rr.Code, expectedStatusCode, t)
  assert_equal(rr.Body.String(), expectedBody, t)
}

func TestDebitWithLimit(t *testing.T) {
  expectedStatusCode := http.StatusOK
	expectedBody := `{"limite":100,"saldo":-100}`

  r := createRouter(&FakeClienteRepository{})
  req, _ := http.NewRequest("POST", "/clientes/1/transacoes", bytes.NewBuffer([]byte(`{"valor": 100, "tipo": "d", "descricao": "descricao"}`)))
  rr := httptest.NewRecorder()
  r.ServeHTTP(rr, req)

  assert_equal(rr.Code, expectedStatusCode, t)
  assert_equal(rr.Body.String(), expectedBody, t)
}


func createRouter(repository *FakeClienteRepository) *mux.Router {
	sut := NewHandler(repository)
  r := mux.NewRouter()
  r.HandleFunc("/clientes/{id}/transacoes", sut.TransacoesDoCliente)

  return r
}

func assert_equal[T int | string](actual T, expected T, t *testing.T) {
	if actual != expected {
		t.Errorf("assert equal failed: actual: %v expected: %v", actual, expected)
	}
}
