package extrato

import "time"
import "m/internal/transacao"

type Extrato struct {
  Saldo             Saldo                   `json:"saldo"`
  UltimasTransacoes []transacao.Transacao   `json:"ultimas_transacoes"`
}

type Saldo struct {
  Total             int       `json:"total"`
  Limite            int       `json:"limite"`
  DataExtrato       time.Time `json:"data_extrato"`
}
