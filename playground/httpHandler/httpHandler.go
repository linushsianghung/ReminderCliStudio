package main

import "net/http"

func httpGetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Amazing!"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
