package controllers

import (
	"fmt"
	"net/http"

	"github.com/praveenk007/dexter/service/interfaces"
	"github.com/praveenk007/dexter/service/interfaces/impl"

	"github.com/gorilla/mux"
)

type ConfigApi struct {
}

func NewConfigApi() *ConfigApi {
	return &ConfigApi{}
}
func (a *ConfigApi) GetConfig(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go is cool!")
	cService := getConfigService()
	cService.GetConfig(mux.Vars(r)["id"], mux.Vars(r)["type"])
}

func getConfigService() interfaces.IConfigService {
	return impl.NewConfigService()
}
