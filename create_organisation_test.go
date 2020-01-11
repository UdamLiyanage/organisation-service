package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateOrganisationAndDelete(t *testing.T) {
	organisation := Organisation{
		Name:    "Test Organisation",
		Devices: nil,
		Users:   nil,
	}
	body, err := json.Marshal(organisation)
	if err != nil {
		t.Fatal(err)
	}
	r := newRouter()
	createRecorder := httptest.NewRecorder()
	createRequest, _ := http.NewRequest("POST", "/organisations", bytes.NewBuffer(body))
	r.ServeHTTP(createRecorder, createRequest)
	if createRecorder.Code != http.StatusCreated {
		t.Errorf("Status should be 201, got %d", createRecorder.Code)
	}

	_ = json.NewDecoder(createRecorder.Body).Decode(&organisation)
	deleteRequest, _ := http.NewRequest("DELETE", "/organisations/"+organisation.ID, nil)
	deleteRecorder := httptest.NewRecorder()
	r.ServeHTTP(deleteRecorder, deleteRequest)
	if deleteRecorder.Code != http.StatusNotFound {
		t.Errorf("Status should be 404, got %d", deleteRecorder.Code)
	}
}
