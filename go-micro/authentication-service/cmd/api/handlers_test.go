package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func Test_Authenticate(t *testing.T) {
	// Given
	jsonToReturn := `
		{
			"error": false,
			"message": "some message"
		}
	`

	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
			Header:     make(http.Header),
		}
	})

	testApp.Client = client

	postBody := map[string]any{
		"email":    "me@here.com",
		"password": "verysecret",
	}

	body, err := json.Marshal(postBody)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/authenticate", bytes.NewReader(body))
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(testApp.Authenticate)

	// When
	handler.ServeHTTP(rr, req)

	// Then
	if rr.Code != http.StatusAccepted {
		t.Errorf("expected http.StatusAccepted but got %d", rr.Code)
	}
}
