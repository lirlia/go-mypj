// fib_test.go
package client

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
	k := &KurlObj{
		Method: "GET",
		Fqdn:   "http://localhost:8080",
	}
	err := k.SendRequest()
	assert.NoError(t, err)
	assert.Equal(t, k.GetResponse().StatusCode, 200)
	assert.Equal(t, k.GetResponse().Payload, Payload([]byte(resBody)))
	assert.Equal(t, k.GetResponse().Payload.Stringer(), resBody)

	// 404 not found
	k = &KurlObj{
		Method: "GET",
		Fqdn:   "http://localhost:8080/helllo",
	}
	resBody = "404 not found"
	err = k.SendRequest()
	assert.NoError(t, err)
	assert.Equal(t, k.GetResponse().Payload, Payload([]byte(resBody)))
	assert.Equal(t, k.GetResponse().Payload.Stringer(), resBody)
}
