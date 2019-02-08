package controllers

import (
	"dexter/service/interfaces"
	"dexter/service/interfaces/impl"
	"fmt"
	"net/http"

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
