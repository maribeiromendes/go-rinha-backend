package transacao

import "time"

type Transacao struct {
	Valor     int    `json:"valor"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
}

type RetornoTransacao struct {
	Limite int `json:"limite"`
	Saldo  int `json:"saldo"`
}

type Extrato struct {
	Total             int         `json:"total"`
	Saldo             int         `json:"saldo"`
	DataExtrato       time.Time   `json:"data_extrato""`
	Limite            int         `json:"limite"`
	UltimasTransacoes []Transacao `json:"ultimas_transacoes"`
}
