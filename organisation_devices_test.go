package main

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestAttachDevice(t *testing.T) {
	attach := AttachDevice{
		DeviceID:       "test_id",
		DeviceName:     "Test_Device",
		OrganisationID: "5e11ecc919fa7f2152331103",
	}

	body, err := json.Marshal(attach)
	if err != nil {
		t.Fatal(err)
	}
	w := testRequestStatusCode("POST", "/organisations/attach/device", body, http.StatusCreated, t)
	testRequestBody(w, "ModifiedCount", 0, t)
}

func TestRemoveDevice(t *testing.T) {
	remove := make(map[string]string)
	remove["organisation_id"] = "5e11ecc919fa7f2152331103"
	remove["device_id"] = "test_id"
	body, err := json.Marshal(&remove)
	if err != nil {
		t.Fatal(err)
	}
	w := testRequestStatusCode("POST", "/organisations/remove/device", body, http.StatusNotFound, t)
	testRequestBody(w, "ModifiedCount", 0, t)
}
