package transacao


type Transacao struct {
	Valor     int    `json:"valor"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
}

type RetornoTransacao struct {
	Limite int `json:"limite"`
	Saldo  int `json:"saldo"`
}

