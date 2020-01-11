package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAttachDevice(t *testing.T) {
	attach := AttachDevice{
		DeviceID:       "Test_Device_ID",
		DeviceName:     "Test_Device",
		OrganisationID: "5e11ecc919fa7f2152331103",
	}

	body, err := json.Marshal(attach)
	if err != nil {
		t.Fatal(err)
	}
	r := newRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/organisations/attach/device", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Status should be 201, got %d", w.Code)
	}
}
