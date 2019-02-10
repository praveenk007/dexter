package main

import (
	"fmt"
	"net/http"

	"github.com/praveenk007/dexter/dao"

	"github.com/praveenk007/dexter/controllers"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting app ...")
	session, error := dao.NewSession("localhost:27017")
	if error != nil {
		fmt.Errorf("Error occurred when conn to db %s", error)
	}

	//init router
	router := mux.NewRouter()

	//add routes
	router.HandleFunc("/api/v1/conf/{type}/{id}", controllers.NewConfigApi(session).GetConfig).Methods("GET")
	http.ListenAndServe(":8080", router)
}
