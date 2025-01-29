package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendGetRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `{"message": "test request"}`)
	}))
	defer ts.Close()

	responseBody, err := sendGetRequest(ts.URL)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	expected := `{"message": "test request"}`
	actual := string(responseBody)
	assert.Equal(t, expected, actual)
}

func TestSendPostRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		body, _ := io.ReadAll(r.Body)
		if !bytes.Equal(body, requestBody) {
			t.Errorf("Expected request body %s, got %s", string(requestBody), string(body))
		}

		fmt.Fprint(w, `{"message": "test request"}`)
	}))
	defer ts.Close()

	requestBody, err := sendPostRequest(ts.URL, requestBody)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	expected := `{"message": "test request"}`
	actual := string(requestBody)
	assert.Equal(t, expected, actual)
}
