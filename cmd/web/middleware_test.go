package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNosurf(t *testing.T) {
	var myHandler testHandler

	h := NoSurf(&myHandler)

	switch v := h.(type) {
	case http.Handler:
		// do nothing - test passed
	default:
		t.Error(fmt.Errorf("type is not http.Handler, but is %T", v))
	}

}

func TestSessionLoad(t *testing.T) {
	var myHandler testHandler

	h := SessionLoad(&myHandler)

	switch v := h.(type) {
	case http.Handler:
		// do nothing - test passed
	default:
		t.Error(fmt.Errorf("type is not http.Handler, but is %T", v))
	}

}
