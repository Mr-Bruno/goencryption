package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to goencryption")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/StoreData", StoreData)
	http.HandleFunc("/GetData", GetData)
	http.ListenAndServe(":8080", nil)

}

//StoreData manages the encryption and storage of the data.
func StoreData(w http.ResponseWriter, r *http.Request) {

}

//GetData manages fetching the data, its decryption and sending it back
func GetData(w http.ResponseWriter, r *http.Request) {

}
