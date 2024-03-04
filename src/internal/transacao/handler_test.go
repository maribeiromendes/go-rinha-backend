package transacao

import (
	"bytes"
	"m/internal/cliente"
	"net/http"
	"net/http/httptest"
	"testing"

  "m/pkg/testutil"
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
  
  handler := NewHandler(&FakeClienteRepository{})
  r := testutil.CreateRouter(handler)
	req, _ := http.NewRequest("POST", "/clientes/1/transacoes", nil)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	testutil.AssertEquals(rr.Code, expectedStatusCode, t)
}

func TestTransacoesDoClienteWithInvalidBody(t *testing.T) {
	expectedStatusCode := http.StatusBadRequest

  handler := NewHandler(&FakeClienteRepository{})
  r := testutil.CreateRouter(handler)
	req, _ := http.NewRequest("POST", "/clientes/1/transacoes", bytes.NewBuffer([]byte(`"test": "not valid"`)))

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	testutil.AssertEquals(rr.Code, expectedStatusCode, t)
}

func TestCredit(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedBody := `{"limite":100,"saldo":1000}`

  handler := NewHandler(&FakeClienteRepository{})
  r := testutil.CreateRouter(handler)
	req, _ := http.NewRequest("POST", "/clientes/1/transacoes", bytes.NewBuffer([]byte(`{"valor": 1000, "tipo": "c", "descricao": "descricao"}`)))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	testutil.AssertEquals(rr.Code, expectedStatusCode, t)
	testutil.AssertEquals(rr.Body.String(), expectedBody, t)
}

func TestDebitWhenNoLimit(t *testing.T) {
  expectedStatusCode := http.StatusUnprocessableEntity
  expectedBody := `{"mensagem":"sem limite"}`

  handler := NewHandler(&FakeClienteRepository{})
  r := testutil.CreateRouter(handler)
  req, _ := http.NewRequest("POST", "/clientes/1/transacoes", bytes.NewBuffer([]byte(`{"valor": 101, "tipo": "d", "descricao": "descricao"}`)))
  rr := httptest.NewRecorder()
  r.ServeHTTP(rr, req)

  testutil.AssertEquals(rr.Code, expectedStatusCode, t)
  testutil.AssertEquals(rr.Body.String(), expectedBody, t)
}

func TestDebitWithLimit(t *testing.T) {
  expectedStatusCode := http.StatusOK
	expectedBody := `{"limite":100,"saldo":-100}`

  handler := NewHandler(&FakeClienteRepository{})
  r := testutil.CreateRouter(handler)
  req, _ := http.NewRequest("POST", "/clientes/1/transacoes", bytes.NewBuffer([]byte(`{"valor": 100, "tipo": "d", "descricao": "descricao"}`)))
  rr := httptest.NewRecorder()
  r.ServeHTTP(rr, req)

  testutil.AssertEquals(rr.Code, expectedStatusCode, t)
  testutil.AssertEquals(rr.Body.String(), expectedBody, t)
}
