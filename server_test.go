package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldShouldSuceed(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testServer.Close()

	testClient := testServer.Client()

	fmt.Println(testServer.URL)

	response, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "Hello World!", string(body))
	assert.Equal(t, 200, response.StatusCode)
}

func TestHelloWorldShouldFail(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testServer.Close()

	testClient := testServer.Client()

	fmt.Println(testServer.URL)

	body := strings.NewReader("some body")

	response, err := testClient.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 405, response.StatusCode)
}

func TestHealthShouldSuceed(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHealth))
	defer testServer.Close()

	testClient := testServer.Client()

	fmt.Println(testServer.URL)

	response, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "ok!", string(body))
	assert.Equal(t, 200, response.StatusCode)
}

func TestHealthShouldFail(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHealth))
	defer testServer.Close()

	testClient := testServer.Client()

	fmt.Println(testServer.URL)

	body := strings.NewReader("some body")

	response, err := testClient.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 405, response.StatusCode)
}

func TestNewEndpointSuceed(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleNewEndpoint))
	defer testServer.Close()

	testClient := testServer.Client()

	fmt.Println(testServer.URL)

	response, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "New Endpoint!", string(body))
	assert.Equal(t, 200, response.StatusCode)
}

func TestNewEndpointFail(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleNewEndpoint))
	defer testServer.Close()

	testClient := testServer.Client()

	fmt.Println(testServer.URL)

	body := strings.NewReader("some body")

	response, err := testClient.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 405, response.StatusCode)
}
