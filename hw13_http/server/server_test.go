package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleGetRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handleRequest))
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL+"/resource", nil)
	if err != nil {
		t.Errorf("не удалось создать запросЖ %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("Ошибка выполнения запроса: %v", err)
	}
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestHandlePostRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handleRequest))
	defer ts.Close()

	req, err := http.NewRequest("POST", ts.URL+"/resource", bytes.NewBufferString(""))
	if err != nil {
		t.Errorf("не удалось создать запрос %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("Ошибка выполнения запроса: %v", err)
	}
	defer res.Body.Close()

	assert.Equal(t, http.StatusCreated, res.StatusCode)
}

func TestHandleInvalidMethod(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handleRequest))
	defer ts.Close()

	req, err := http.NewRequest("PUT", ts.URL+"/resource", nil)
	if err != nil {
		t.Errorf("не удалось создать запрос %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("Ошибка выполнения запроса: %v", err)
	}
	defer res.Body.Close()

	assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)
}
