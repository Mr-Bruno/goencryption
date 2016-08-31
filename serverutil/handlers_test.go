package serverutil

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestHttpFlow(t *testing.T) {

	id := "123"
	OriginalText := string("It should appear this")
	body := []byte(OriginalText)

	//Store the data
	url := "http://localhost:8080/StoreData?id=" + id
	r, _ := http.NewRequest("POST", url, bytes.NewReader(body))
	r.Header.Add("Content-Type", "application/octet-stream")
	r.Header.Add("Content-Length", strconv.Itoa(len(body)))

	StoreRR := httptest.NewRecorder()
	StoreHandler := http.HandlerFunc(StoreData)
	StoreHandler.ServeHTTP(StoreRR, r)

	key, _ := ioutil.ReadAll(StoreRR.Body)
	body = key

	//Gets the data
	urlget := "http://localhost:8080/GetData?id=" + id
	reqg, _ := http.NewRequest("POST", urlget, bytes.NewReader(body))
	reqg.Header.Add("Content-Type", "application/octet-stream")
	reqg.Header.Add("Content-Length", strconv.Itoa(len(body)))

	GetRR := httptest.NewRecorder()
	GetHandler := http.HandlerFunc(GetData)
	GetHandler.ServeHTTP(GetRR, reqg)

	BackToOriginal, _ := ioutil.ReadAll(GetRR.Body)

	if OriginalText != string(BackToOriginal) {
		t.Error("Expected content", OriginalText, "and got", string(BackToOriginal))
	}

}
