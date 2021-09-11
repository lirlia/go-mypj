package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/lirlia/go-mypj/voyage/kadai2/model"
	"golang.org/x/net/html"
)

func Get(url string) (*model.Page, error) {
	page := new(model.Page)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var f func(*html.Node)
	// fはDOMツリーを再帰的にトラバースするための手続きです。
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			page.Title = n.FirstChild.Data
		}

		// 再帰的にノードをおっていくために、次のノードを探し、
		// ノードが存在すれば再びこのfを実行する、ということをしています。
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return page, nil
}
func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		querys, _ := url.ParseQuery(r.URL.RawQuery)
		if querys.Get("url") != "" {
			p, err := Get(querys.Get("url"))
			if err != nil && err != io.EOF {
				panic(err)
			}

			msg := fmt.Sprintf("%v", p.Title)
			w.Write([]byte(msg))

		}
	})

	http.ListenAndServe(":8080", nil)
}
