package main

import (
    "fmt"
    "net/http"
    "time"
    "encoding/json"
    "io/ioutil"
    "log"

    "gorilla/mux"
)

type Transacao struct {
    Valor int `json:"valor"`
    Tipo string `json:"tipo"`
    Descricao string `json:"descricao"`
}

type RetornoTransacao struct {
    Limite int `json:"limite"`
    Saldo int `json:"saldo"`
}

type Extrato struct {
    Total int `json:"total"`
    Saldo int `json:"saldo"`
    DataExtrato time.Time `json:"data_extrato""`
    Limite int `json:"limite"`
    UltimasTransacoes []Transacao `json:"ultimas_transacoes"`
}

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Println("/hello endpoint called")
    fmt.Fprintf(w, "hello\n")
}

func clientesTransacoes(w http.ResponseWriter, r *http.Request) {
    fmt.Println("/clientes/{id}/transacoes endpoint called")
    vars := mux.Vars(r)
    
    fmt.Fprintf(rw, "hello\n")
    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        panic(err)
    }
    log.Println(string(body))
    var t Transacao
    err = json.Unmarshal(body, &t)
    if err != nil {
        panic(err)
    }
    log.Println(t.Valor)
    log.Println(vars["tipo"])
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/clientes/{id}/transacoes", clientesTransacoes)
    fmt.Println("Server up and listening...")
    http.ListenAndServe(":8000", nil)
}