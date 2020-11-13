package main

import (
	"log"
	"net/http"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
)

func transformJSON(data []string) (jsonData []byte, err error) {
	json := simplejson.New()
	json.Set("data", data)
	jsonData, err = json.MarshalJSON()
	return
}

func GetAllUsers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	data, err := transformJSON(AllUsers())

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.WriteHeader(http.StatusOK)
		writer.Write(data)
	}
}

func startServer(router *mux.Router) {
	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users/all", GetAllUsers)
	http.Handle("/", router)
	startServer(router)
}
