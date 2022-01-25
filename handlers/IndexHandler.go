package handlers

import "net/http"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello, World!"))
}
