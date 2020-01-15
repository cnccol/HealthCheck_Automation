package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"time"
	"fmt"
)

type Data struct {
	Id int `json:"Id"`
	VideoSize int `json:"VideoSize"`
}

func (p Data) String() string{
	return fmt.Sprintf("{Id: %v, VideoSize: %v}", p.Id, p.VideoSize)
}

var datos []Data

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datos)
}

func createData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data Data
	_ = json.NewDecoder(r.Body).Decode(&data)
	datos = append(datos, data)
	json.NewEncoder(w).Encode(&data)
}

func Run(){
	router := mux.NewRouter()

	datos = append(datos, Data{Id: 1, VideoSize: 2})
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