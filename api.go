package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
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
		client := &http.Client{
			Timeout: 2 * time.Second,
		}
		res, err := client.Do(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		for name, values := range res.Header {
			w.Header().Set(name, values[0])
		}
		w.WriteHeader(res.StatusCode)
		w.Write(body)
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
