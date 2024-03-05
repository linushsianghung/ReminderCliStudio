package main

import "net/http"

func main() {
	http.HandleFunc("/hello", helloFunc)
	http.ListenAndServe(":8888", nil)
}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello World!"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
