package main

import (
	"fmt"
	"net/http"

	"github.com/goji/httpauth"
)

func main() {
	// HandlerFunc は与えられたパターン（/hello) をDefaultServeMuxに登録します
	// http.HandleFunc("/hello", helloHandler)
	authHandler := httpauth.SimpleBasicAuth("username", "password")

	http.ListenAndServe(":8080",
		authHandler(middleware2(middleware1(http.HandlerFunc(helloHandler)))))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hello" && r.Method == "GET" {
		fmt.Println("Hello")
	}
}

func middleware1(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware1")
		h.ServeHTTP(w, r)
		fmt.Println("middleware1")
	})
}

func middleware2(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware2")
		h.ServeHTTP(w, r)
		fmt.Println("middleware2")
	})
}
