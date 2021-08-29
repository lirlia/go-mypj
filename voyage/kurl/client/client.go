package client

import (
	"io"
	"net/http"

	"github.com/lirlia/go-mypj/voyage/kurl/config"
)

type kurlObj struct {
	method  string
	fqdn    string
	payload []byte
	headers headers
	response
}

type header struct {
	key   string
	value string
}

type response struct {
	payload    payload
	statusCode int
}

type payload []byte

func (p payload) stringer() string {
	return string(p)
}

type headers []header

func (k *kurlObj) addHeader(key string, value string) {
	k.headers = append(k.headers, header{key, value})
}

func (h headers) existsHeader(keyName string) bool {
	for _, header := range h {
		if header.key == keyName {
			return true
		}
	}
	return false
}

// リクエスト組み立て(get/post)

// リクエスト発行
func (k *kurlObj) sendRequest() error {
	req, err := http.NewRequest(k.method, k.fqdn, nil)

	if !k.headers.existsHeader("Content-Type") {
		k.addHeader("Content-Type", config.DEFAULT_HEADER_CONTENT_TYPE)
	}

	// headerの付与
	for _, header := range k.headers {
		req.Header.Add(header.key, header.value)
	}

	// responseの格納
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	k.response.statusCode = res.StatusCode

	switch res.StatusCode {
	case 404:
		k.response.payload = []byte("404 not found")

	default:
		defer res.Body.Close()
		k.response.payload, err = io.ReadAll(res.Body)

		if err != nil {
			return err
		}
	}
	return nil
}

func (k *kurlObj) getResponse() response {
	return k.response
}
