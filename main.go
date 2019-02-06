package main

import (
	"dexter/controllers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting app ...")
	//init router
	router := mux.NewRouter()

	//add routes
	router.HandleFunc("/api/v1/getconf", controllers.NewConfigApi().GetConfig).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func getConfig() {

}
