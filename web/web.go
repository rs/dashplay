package api

import "net/http"

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// This is a dummy go server, everything is handled by web.yaml
		w.Write([]byte("Hello World"))
	})
}
