package client

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Client provides functionality to interact with the encryption-server
type Client interface {
	// Store accepts an id and a payload in bytes and requests that the
	// encryption-server stores them in its data store
	Store(id, payload []byte) (aesKey []byte, err error)

	// Retrieve accepts an id and an AES key, and requests that the
	// encryption-server retrieves the original (decrypted) bytes stored
	// with the provided id
	Retrieve(id, aesKey []byte) (payload []byte, err error)
}

//MyClient Implements the functionality to interact with the encryption-server
// through the Client interface
type MyClient struct {
}

// Store functionality to store data
func (Mc MyClient) Store(id, payload []byte) (aesKey []byte, err error) {

	url := "http://localhost:8080/StoreData?id=" + string(id)
	aeskey, err := DoPost(url, payload)
	return aeskey, err
}

// Retrieve functionality to get data
func (Mc MyClient) Retrieve(id, aesKey []byte) (payload []byte, err error) {

	url := "http://localhost:8080/GetData?id=" + string(id)
	payload, err = DoPost(url, aesKey)
	return payload, err
}

//DoPost encapsulates the POST requests to do into the server
func DoPost(url string, body []byte) ([]byte, error) {

	client := &http.Client{}
	r, _ := http.NewRequest("POST", url, bytes.NewReader(body))
	r.Header.Add("Content-Type", "application/octet-stream")
	r.Header.Add("Content-Length", strconv.Itoa(len(body)))
	response, _ := client.Do(r)

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("Response gave an error status")
	}

	defer response.Body.Close()

	ResponseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Error processing body")
	}

	return ResponseBody, nil
}
