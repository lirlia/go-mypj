package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/lirlia/go-mypj/voyage/kadai2/model"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		querys, _ := url.ParseQuery(r.URL.RawQuery)

		if querys.Get("url") != "" {

			urls := strings.Split(querys.Get("url"), ",")
			c := make(chan model.Page, len(urls))
			wg := &sync.WaitGroup{}

			for _, u := range urls {
				url := u
				wg.Add(1)
				go func() {
					defer wg.Done()
					p := model.Get(url)
					c <- *p
				}()
			}

			wg.Wait()

			// 終了までまつ
			var Pages []model.Page

			for {
				Pages = append(Pages, <-c)
				if len(c) == 0 {
					close(c)
					break
				}
			}

			enc := json.NewEncoder(w)
			enc.SetIndent("", "    ")
			enc.Encode(Pages)
		}
	})

	http.ListenAndServe("127.0.0.1:8080", nil)
}
