package testutil

import (
	"m/internal/common"

	"github.com/gorilla/mux"
)

func CreateRouter(handler common.Handler) *mux.Router {
  r := mux.NewRouter()
  handler.SetRoutes(r)

  return r
}
