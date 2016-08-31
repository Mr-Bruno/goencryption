package main

import (
	"net/http"

	"github.com/Mr-Bruno/goencryption/serverutil"
)

func main() {
	http.HandleFunc("/", serverutil.Handler)
	http.HandleFunc("/StoreData", serverutil.StoreData)
	http.HandleFunc("/GetData", serverutil.GetData)
	http.ListenAndServe(":8080", nil)

}
