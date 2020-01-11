package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAttachOwner(t *testing.T) {
	attach := AttachOwner{
		UserID:         "test_user_id",
		UserName:       "test_user_name",
		OrganisationID: "5e11ecc919fa7f2152331103",
	}

	body, err := json.Marshal(attach)
	if err != nil {
		t.Fatal(err)
	}
	r := newRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/organisations/attach/user", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Status should be 201, got %d", w.Code)
	}
}

func TestRemoveAttachedUser(t *testing.T) {
	remove := make(map[string]string)
	remove["organisation_id"] = "5e11ecc919fa7f2152331103"
	remove["user"] = "test_user_id"
	body, err := json.Marshal(&remove)
	if err != nil {
		t.Fatal(err)
	}
	r := newRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/organisations/remove/user", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	if w.Code != http.StatusNotFound {
		t.Errorf("Status should be 404, got %d", w.Code)
	}
}
