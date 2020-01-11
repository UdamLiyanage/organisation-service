package main

import (
	"net/http"
	"testing"
)

func TestValidDeviceGet(t *testing.T) {
	//Object ID - 5e11ecc919fa7f2152331103 is always available as a test document
	testRequestStatusCode("GET", "/organisations/5e11ecc919fa7f2152331103", nil, http.StatusOK, t)
}

func TestInvalidDeviceGet(t *testing.T) {
	testRequestStatusCode("GET", "/organisations/000000000000000000000000", nil, http.StatusNotFound, t)
}
