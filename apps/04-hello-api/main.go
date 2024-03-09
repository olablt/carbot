package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
)

func main() {
	svc := &Service{}

	mux := &http.ServeMux{}
	mux.Handle("POST /notes", Handle(svc.CreateNote))
	mux.Handle("PUT /notes/{noteID}", Handle(svc.UpdateNote))

	requests := []struct {
		method string
		url    string
		data   string
	}{
		{http.MethodPost, "/notes", `{"note": "Hello world!"}`},
		{http.MethodPut, "/notes/1023", `{"note": "Updated content!"}`},
	}

	for _, r := range requests {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(r.method, r.url, strings.NewReader(r.data))

		mux.ServeHTTP(rr, req)

		res := rr.Result()
		defer res.Body.Close()

		b, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalf("failed to read response body: %v", err)
		}

		fmt.Printf("response status: %d, response body: %s\n", res.StatusCode, b)
	}
}
