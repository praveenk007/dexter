package main

import (
	"fmt"
	"net/http"

	"github.com/praveenk007/dexter/controllers"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting app ...")
	//init router
	router := mux.NewRouter()

	//add routes
	router.HandleFunc("/api/v1/conf/{type}/{id}", controllers.NewConfigApi().GetConfig).Methods("GET")
	http.ListenAndServe(":8080", router)
}
