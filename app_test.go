package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	server := GetServer()
	testServer := httptest.NewServer(server)
	defer testServer.Close()
	_, err := http.Get(fmt.Sprintf("%s/ping", testServer.URL))
	if err != nil {
		t.Fatalf(err.Error())
	}
}
