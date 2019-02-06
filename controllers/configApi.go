package controllers

import (
	"dexter/service"
	"fmt"
	"net/http"
)

type ConfigApi struct {
}

func NewConfigApi() *ConfigApi {
	return &ConfigApi{}
}
func (a *ConfigApi) GetConfig(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go is cool!")
	service.NewConfigService().GetBrokerConfig()
}
