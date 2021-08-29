// fib_test.go
package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKurl(t *testing.T) {
	resBody := "Hello, client"
	port := ":8080"
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, resBody)
	})

	l, _ := net.Listen("tcp", port)
	ts := httptest.Server{
		Listener: l,
		Config:   &http.Server{Handler: handler},
	}
	ts.Start()
	defer ts.Close()

	// 200 ok
	k := &kurlObj{
		method: "GET",
		fqdn:   "http://localhost:8080",
	}
	err := k.sendRequest()
	assert.NoError(t, err)
	assert.Equal(t, k.getResponse().statusCode, 200)
	assert.Equal(t, k.getResponse().payload, payload([]byte(resBody)))
	assert.Equal(t, k.getResponse().payload.stringer(), resBody)

	// 404 not found
	k = &kurlObj{
		method: "GET",
		fqdn:   "http://localhost:8080/helllo",
	}
	resBody = "404 not found"
	err = k.sendRequest()
	assert.NoError(t, err)
	assert.Equal(t, k.getResponse().payload, payload([]byte(resBody)))
	assert.Equal(t, k.getResponse().payload.stringer(), resBody)
}
