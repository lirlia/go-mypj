package client

import (
	"fmt"
	"io"
	"net/http"

	"github.com/lirlia/go-mypj/voyage/kurl/config"
)

type KurlObj struct {
	Method  string
	Fqdn    string
	Payload []byte
	Headers Headers
	Response
}

type Header struct {
	Key   string
	Value string
}

type Response struct {
	Payload    Payload
	StatusCode int
}

type Payload []byte

func (p Payload) Stringer() string {
	return string(p)
}

type Headers []Header

func (k *KurlObj) AddHeader(key string, value string) {
	k.Headers = append(k.Headers, Header{key, value})
}

func (h Headers) existsHeader(keyName string) bool {
	for _, header := range h {
		if header.Key == keyName {
			return true
		}
	}
	return false
}

// リクエスト組み立て(get/post)

// リクエスト発行
func (k *KurlObj) SendRequest() error {
	req, err := http.NewRequest(k.Method, k.Fqdn, nil)

	if !k.Headers.existsHeader("Content-Type") {
		k.AddHeader("Content-Type", config.DEFAULT_HEADER_CONTENT_TYPE)
	}

	// headerの付与
	for _, header := range k.Headers {
		req.Header.Add(header.Key, header.Value)
	}

	// responseの格納
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	k.Response.StatusCode = res.StatusCode

	switch res.StatusCode {
	case 404:
		k.Response.Payload = []byte("404 not found")

	default:
		defer res.Body.Close()
		k.Response.Payload, err = io.ReadAll(res.Body)

		if err != nil {
			return err
		}
	}
	return nil
}

func (k *KurlObj) GetResponse() Response {
	return k.Response
}

func (k *KurlObj) Run() error {
	err := k.SendRequest()
	if err != nil {
		return err
	}

	res := k.GetResponse()
	fmt.Println(res.Payload.Stringer())
	return nil
}

func CheckBeforeCallRun(k *KurlObj) error {
	if !ValidateFqdn(k.Fqdn) {
		return fmt.Errorf("FQDN is invalid. URL must be set http[s]://xxxxxxx .")
	}
	return nil
}

// URLが正しいかをチェックします
func ValidateFqdn(f string) bool {
	// TODO
	return true
}
