package main

import (
	"log"
	"net/http"
	bniCallback "sakti-cashmanagement/bni/callback"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	/* BNI */
	router.HandleFunc("/sakti-cashmanagement/api/v1/bni-callback", bniCallback.AddNew).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", router))

}
