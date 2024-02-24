package transacao

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestTransacao(t *testing.T)  {
  req, err := http.NewRequest("GET", "/health", nil) 
  if err != nil {
    t.Fatal(err)
  }

  rr := httptest.NewRecorder()
  sut := NewHandler()
  log.Println("Created handler")
  //TODO: Having a null exception. The handler is being executed, but the request is probably invalid
  handler := http.HandlerFunc(sut.TransacoesDoCliente)

  handler.ServeHTTP(rr, req)

  if status := rr.Code; status != http.StatusOK {
    t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
  }

}
