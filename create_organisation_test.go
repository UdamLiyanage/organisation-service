package main

import (
	"encoding/json"
	"net/http"
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
	createRecorder := testRequestStatusCode("POST", "/organisations", body, http.StatusCreated, t)

	_ = json.NewDecoder(createRecorder.Body).Decode(&organisation)
	testRequestStatusCode("DELETE", "/organisations/"+organisation.ID, nil, http.StatusNotFound, t)
}
