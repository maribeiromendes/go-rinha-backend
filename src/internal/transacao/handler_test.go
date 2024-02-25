package transacao

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestTransacoesDoClienteWithoutBody(t *testing.T)  {
  expectedStatusCode := http.StatusBadRequest
  req, _ := http.NewRequest("POST", "/clientes/1/transacoes", nil) 
  rr := httptest.NewRecorder()
  sut := NewHandler()

  handler := http.HandlerFunc(sut.TransacoesDoCliente)

  handler.ServeHTTP(rr, req)

  assert_equal(rr.Code, expectedStatusCode, t)
}

func TestTransacoesDoClienteWithInvalidBody(t *testing.T)  {
  expectedStatusCode := http.StatusBadRequest
  req, _ := http.NewRequest("POST", "/clientes/1/transacoes", bytes.NewBuffer([]byte(`"test": "not valid"`))) 
  req.Header.Set("Content-Type", "application/json; charset=UTF-8")
  rr := httptest.NewRecorder()
  sut := NewHandler()

  handler := http.HandlerFunc(sut.TransacoesDoCliente)

  handler.ServeHTTP(rr, req)
  assert_equal(rr.Code, expectedStatusCode, t)
}

func TestTransacoesDoClienteWithValidBody(t *testing.T) {
  expectedStatusCode := http.StatusOK
  expectedBody := `{"limite":10,"saldo":5}` 

  req, _ := http.NewRequest("POST", "/clientes/1/transacoes", bytes.NewBuffer([]byte(`{"valor": 1000, "tipo": "c", "descricao": "descricao"}`))) 
  rr := httptest.NewRecorder()
  sut := NewHandler()

  handler := http.HandlerFunc(sut.TransacoesDoCliente)

  handler.ServeHTTP(rr, req)

  assert_equal(rr.Code, expectedStatusCode, t)
  assert_equal(rr.Body.String(), expectedBody, t)
}


func assert_equal[T int | string](expected T, actual T, t *testing.T) {
  if actual != expected {
    t.Errorf("assert equal failed: actual: %v expected: %v", actual, expected)
  }

}
