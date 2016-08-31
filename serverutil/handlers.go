package serverutil

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

var database = CreateDatabase()

//Handler basic one to show a message
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to goencryption")
}

//StoreData manages the encryption and storage of the data.
func StoreData(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		values, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Error:", err)
			return
		}

		// Read the id
		if len(values.Get("id")) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Error:", "Wrong input id.")
			return
		}

		//Read the payload
		lengthBody, _ := strconv.Atoi(r.Header.Get("Content-Length"))
		payload := make([]byte, lengthBody)
		leng, err := r.Body.Read(payload)
		if leng == 0 {
			fmt.Println(err)
		}

		//generate the key, encrypy the text and add it to the storage
		key := RandomKey(32)
		cipherData := Encrypt(key, payload)
		AddElement(database, values.Get("id"), cipherData)

		// Sending back the information (key)
		w.WriteHeader(http.StatusOK)
		w.Write(key)

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error: Only POST accepted")
	}

}

//GetData manages fetching the data, its decryption and sending it back
func GetData(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		values, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Error:", err)
			return
		}

		// Read the id
		if len(values.Get("id")) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Error:", "Wrong input id.")
			return
		}

		// Read the key
		lengthBody, _ := strconv.Atoi(r.Header.Get("Content-Length"))
		key := make([]byte, lengthBody)
		leng, err := r.Body.Read(key)
		if leng == 0 {
			fmt.Println(err)
		}

		// Search for the cipher data in the database and decrypt it
		cipherData := GetElement(database, values.Get("id"))
		data := Decrypt(key, cipherData)

		// Sending back the information (decrypted data)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, data)

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error: Only POST accepted")
	}
}
