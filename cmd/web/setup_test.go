package main

import (
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

type testHandler struct {
}

func (mh *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
