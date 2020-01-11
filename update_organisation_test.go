package main

import (
	"encoding/json"
	"net/http"
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
	w := testRequestStatusCode("PUT", "/organisations/5e11ecc919fa7f2152331103", body, http.StatusOK, t)
	testRequestBody(w, "ModifiedCount", 0, t)
}

func TestUpdateDeviceInvalid(t *testing.T) {
	organisation := Organisation{
		Name: "Updated Name",
	}
	body, err := json.Marshal(organisation)
	if err != nil {
		t.Fatal(err)
	}
	w := testRequestStatusCode("PUT", "/organisations/000000000000000000000000", body, http.StatusOK, t)
	testRequestBody(w, "MatchedCount", 0, t)
}
