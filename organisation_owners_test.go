package main

import (
	"encoding/json"
	"net/http"
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
	w := testRequestStatusCode("POST", "/organisations/attach/user", body, http.StatusCreated, t)
	testRequestBody(w, "ModifiedCount", 0, t)
}

func TestRemoveAttachedUser(t *testing.T) {
	remove := make(map[string]string)
	remove["organisation_id"] = "5e11ecc919fa7f2152331103"
	remove["user"] = "test_user_id"
	body, err := json.Marshal(&remove)
	if err != nil {
		t.Fatal(err)
	}
	w := testRequestStatusCode("POST", "/organisations/remove/user", body, http.StatusNotFound, t)
	testRequestBody(w, "ModifiedCount", 0, t)
}
