package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_httpGetHandlerR(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Error(err)
	}

	httpGetHandler(rr, req)
	if rr.Result().StatusCode != http.StatusOK {
		t.Errorf("Expect 200 but get %d: ", rr.Result().StatusCode)
	}

	defer rr.Result().Body.Close()
	expected := "Amazing!"
	b, err := io.ReadAll(rr.Result().Body)
	if err != nil {
		t.Error(err)
	}

	if string(b) != expected {
		t.Errorf("Expect %s but get %s", expected, string(b))
	}

}

func Test_httpGetHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	server := httptest.NewServer(http.HandlerFunc(httpGetHandler))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expect 200 but get %d: ", resp.StatusCode)
	}

	defer resp.Body.Close()
	expected := "Amazing!"
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	if string(b) != expected {
		t.Errorf("Expect %s but get %s", expected, string(b))
	}
	// tests := []struct {
	// 	name string
	// 	args args
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		httpGetHandler(tt.args.w, tt.args.r)
	// 	})
	// }
}
