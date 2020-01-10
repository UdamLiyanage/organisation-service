package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidDeviceGet(t *testing.T) {
	//Object ID - 5e11ecc919fa7f2152331103 is always available as a test document
	r := newRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/organisations/5e11ecc919fa7f2152331103", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Status should be 200, got %d", w.Code)
	}
}

func TestInvalidDeviceGet(t *testing.T) {
	r := newRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/organisations/000000000000000000000000", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusNotFound {
		t.Errorf("Status should be 404, got %d", w.Code)
	}
}
