package model

import (
	"net/http"
	"time"

	"golang.org/x/net/html"
)

type Page struct {
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	OgTitle     string `json:"og_title"`
	OgImage     string `json:"og_image"`
	Status      string `json:"status"`
	Reason      string `json:"reason"`
}

func isDescription(attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if attr.Key == "name" && attr.Val == "description" {
			/* KeyがnameでValがdescriptionか？ */
			return true
		}
	}
	return false
}

func isOgTitle(attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if attr.Key == "property" && attr.Val == "og:title" {
			/* KeyがnameでValがdescriptionか？ */
			return true
		}
	}
	return false
}

func isOgImage(attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if attr.Key == "property" && attr.Val == "og:image" {
			/* KeyがnameでValがdescriptionか？ */
			return true
		}
	}
	return false
}

func Get(url string) *Page {
	page := &Page{}
	page.URL = url

	c := &http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := c.Get(url)
	if resp != nil {
		page.Status = resp.Status
	}
	if err != nil {
		page.Reason = err.Error()
		return page
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		page.Reason = err.Error()
		return page
	}

	var f func(*html.Node)
	// fはDOMツリーを再帰的にトラバースするための手続きです。
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			page.Title = n.FirstChild.Data
		}

		if n.Type == html.ElementNode && n.Data == "meta" {
			if isDescription(n.Attr) {
				for _, attr := range n.Attr {
					// contentの中身を取得
					page.Description = attr.Val
				}
			}

			if isOgTitle(n.Attr) {
				for _, attr := range n.Attr {
					// contentの中身を取得
					page.OgTitle = attr.Val
				}
			}

			if isOgImage(n.Attr) {
				for _, attr := range n.Attr {
					// contentの中身を取得
					page.OgImage = attr.Val
				}
			}
		}

		// 再帰的にノードをおっていくために、次のノードを探し、
		// ノードが存在すれば再びこのfを実行する、ということをしています。
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return page
}
