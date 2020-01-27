package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"time"
	"github.com/cnccol/HealthCheck_Automation/Common"
)

var datos []Common.Data

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datos)
}

func createData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data Common.Data
	_ = json.NewDecoder(r.Body).Decode(&data)
	datos = append(datos, data)
	json.NewEncoder(w).Encode(&data)
}

func Run(){
	router := mux.NewRouter()

	datos = append(datos, Common.Data{Id: 1, VideoSize: 2})
	router.HandleFunc("/data", getData).Methods("GET")
	router.HandleFunc("/data", createData).Methods("POST")
	srv := &http.Server{
		Handler: router,
		Addr: ":9000",
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	srv.ListenAndServe()
}