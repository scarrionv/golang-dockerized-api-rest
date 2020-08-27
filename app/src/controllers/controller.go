package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/scarrionv/simple-api/app/src/services"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetUsers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := services.GetUsers()
	log.Info("[Controllers] -> [GetUsers]: obtained users: ", users)
	_ = json.NewEncoder(w).Encode(users)
}

func CreateControllers(port string) {
	log.Info("[createController]: Creating controllers in localhost:" + port)
	router := mux.NewRouter().StrictSlash(true)
	log.Info("[createController]: New endpoint /user")
	router.HandleFunc("/users", GetUsers).Methods("GET")
	log.Info(http.ListenAndServe(":"+port, router))
}
