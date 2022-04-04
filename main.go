package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type BSN struct {
	Nummer string `json:"Nummer"`
	Naam   string `json:"Naam"`
}

var BSNnummers []BSN

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/bsnnummers", returnAllBSN)
	myRouter.HandleFunc("/bsntoevoegen", BSNToevoegen).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func BSNToevoegen(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var bsn BSN
	json.Unmarshal(reqBody, &bsn)
	BSNnummers = append(BSNnummers, bsn)
	json.NewEncoder(w).Encode(bsn)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the BSN World!")
	fmt.Println("Einde behaald")
}
func returnAllBSN(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Einde behaald deel 2: returnAllBSN")
	json.NewEncoder(w).Encode(BSNnummers)

}
func main() {
	BSNnummers = []BSN{
		{Nummer: "123456789", Naam: "Noa"},
		{Nummer: "987654321", Naam: "Hog Rider"},
	}

	// Eventueel vanuit bestand lezen
	// data, err := ioutil.ReadFile("data.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var slice []string
	// err = json.Unmarshal(data, &slice)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	HandleRequests()
}
