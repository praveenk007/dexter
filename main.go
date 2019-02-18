package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/spf13/viper"

	"github.com/praveenk007/dexter/dao"

	"github.com/praveenk007/dexter/controllers"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting app ....")
	fmt.Println(viper.GetString("DB_HOST"))
	bindEnvVariables()
	session, error := dao.NewSession("mongodb://pk.local:27017")
	if error != nil {
		fmt.Println(error)
		fmt.Errorf("Error occurred when conn to db %s", error)
	}
	fmt.Println("Connected to mongo server")
	env := viper.GetString("ENV")

	initConfig(env)
	fmt.Println("Initialized configuration for environment " + env)

	//init router
	router := mux.NewRouter()

	//add routes
	router.HandleFunc("/api/v1/conf/{type}/{id}", controllers.NewConfigApi(session).GetConfig).Methods("GET")
	http.ListenAndServe(":"+strconv.Itoa(viper.GetInt("PORT")), router)
}

func bindEnvVariables() {
	viper.BindEnv("ENV")
	viper.BindEnv("DB_HOST")
}

func initConfig(env string) {
	viper.SetConfigType("yaml")
	switch environment := env; environment {
	case "PROD":
		viper.SetConfigName("prod")
	case "DEV":
		viper.SetConfigName("local")
	default:
		viper.SetConfigName("local")
	}

	viper.AddConfigPath("configs/")
	viper.ReadInConfig()
}
