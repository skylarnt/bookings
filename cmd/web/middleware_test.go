package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myHttp myHandler
	h := NoSurf(&myHttp)

	switch v := h.(type) {
	case http.Handler:
		//do nothing
	default:
		t.Error(fmt.Sprintf("Type is not http.handler, but %T", v))

	}
}

func TestSessionLoad(t *testing.T) {
	var myHttp myHandler
	h := SessionLoad(&myHttp)

	switch v := h.(type) {
	case http.Handler:
		//do nothing
	default:
		t.Error(fmt.Sprintf("Type is not http.handler, but %T", v))

	}
}
