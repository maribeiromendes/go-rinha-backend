package common

type Repository interface {
  GetById(id int) (interface{}, error)
}
