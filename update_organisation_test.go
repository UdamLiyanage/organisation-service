package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestUpdateDeviceValid(t *testing.T) {
	organisation := Organisation{
		Name: "Updated Name",
	}
	body, err := json.Marshal(organisation)
	if err != nil {
		t.Fatal(err)
	}
	r := newRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/organisations/5e11ecc919fa7f2152331103", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Status should be 200, got %d", w.Code)
	}
}

func TestUpdateDeviceInvalid(t *testing.T) {
	organisation := Organisation{
		Name: "Updated Name",
	}
	body, err := json.Marshal(organisation)
	if err != nil {
		t.Fatal(err)
	}
	r := newRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/organisations/000000000000000000000000", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Status should be 200, got %d", w.Code)
	}
	var response map[string]string
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	value, exists := response["MatchedCount"]
	if !exists {
		t.Errorf("Wrong Response Format")
	}
	count, _ := strconv.Atoi(value)
	if count != 0 {
		t.Errorf("Operation Not Working Properly!")
	}
}
