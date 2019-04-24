package controllers

import (
	"fmt"
	"net/http"

	"github.com/praveenk007/dexter/dao"

	"github.com/praveenk007/dexter/service/interfaces"
	"github.com/praveenk007/dexter/service/interfaces/impl"

	"github.com/gorilla/mux"
)

type ConfigApi struct {
	session *dao.Session
}

func NewConfigApi(session *dao.Session) *ConfigApi {
	return &ConfigApi{session}
}
func (a *ConfigApi) GetConfig(w http.ResponseWriter, r *http.Request) {
	cService := getConfigService(a.session)
	result := cService.GetConfig(mux.Vars(r)["id"], mux.Vars(r)["type"])
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(*result))
}

func getConfigService(session *dao.Session) interfaces.IConfigService {
	return impl.NewConfigService(session)
}
