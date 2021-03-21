package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestHandler(t *testing.T) {

    req, err := http.NewRequest(
        http.MethodGet,
        "http://localhost:9000/",
        nil,
    )

    if err != nil {
        t.Fatalf("Could not create a request %v", err)
    }

    rec := httptest.NewRecorder()
    helloWorldEndPoint(rec, req)

    if rec.Code != http.StatusOK {
        t.Errorf("accepted status 200, got %v", rec.Code)
    }

    if !strings.Contains(rec.Body.String(), "hello world") {
        t.Errorf("unexpected body in response %q", rec.Body.String())
    }

}
