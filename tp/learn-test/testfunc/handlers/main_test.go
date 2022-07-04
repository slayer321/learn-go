package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestDoubleHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "localhost:8080/double?v=2", nil)
	if err != nil {
		log.Fatalf("Cloud not created request: %v", err)
	}

	rec := httptest.NewRecorder()

	doubleHandler(rec, req)

	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK: got %v", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	d, err := strconv.Atoi(string(bytes.TrimSpace(b)))
	if err != nil {
		t.Fatalf("expected an Interger; got %s", b)
	}

	if d != 4 {
		t.Fatalf("expected double to be 4;got %v", d)
	}

}
