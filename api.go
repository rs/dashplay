package api

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"

	"appengine"
	"appengine/urlfetch"
)

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		path := base64.URLEncoding.EncodeToString([]byte(req.URL.Path[1:]))
		switch req.Method {
		case "OPTIONS", "GET", "PUT", "DELETE", "PATCH":
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		r, err := http.NewRequest(req.Method, fmt.Sprintf("https://dashplay.firebaseio.com/urls/%s.json", path), req.Body)
		r.Header = req.Header
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		c := appengine.NewContext(req)
		client := urlfetch.Client(c)
		res, err := client.Do(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for name, values := range res.Header {
			w.Header().Set(name, values[0])
		}
		w.WriteHeader(res.StatusCode)
		w.Write(body)
	})
}
