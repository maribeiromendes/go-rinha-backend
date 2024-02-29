package cliente

import "fmt"

type ClienteRepository struct {}

func (r *ClienteRepository) GetById(id int) (interface{}, error) {
  if id > 5 {
    return nil, fmt.Errorf("cliente com id: %d nao encontrado", id)
  }

  return &Cliente {Limite: 100, Saldo: 0}, nil
}
