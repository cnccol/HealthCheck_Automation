package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
)

type Data struct {
	Id int `json:"id"`
	VideoSize int `json:"videoSize"`
}

var datos []Data

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datos)
}

func Run(){
	router := mux.NewRouter()

	datos = append(datos, Data{Id: 1, VideoSize: 2})
	router.HandleFunc("/data", getData).Methods("GET")

	http.ListenAndServe(":9000", router)
}