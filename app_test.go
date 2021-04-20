package main

import (
	"ecommerce/controllers"
	"gopkg.in/yaml.v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/test-get", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.TestGet)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := `GET successful`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestConfig(t *testing.T) {
	type testConfig struct {
		TestEnv string `yaml:"test_env"`
	}
	var tc testConfig
	conf := `test_env: test`
	if err := yaml.Unmarshal([]byte(conf), &tc); err != nil {
		t.Errorf("Error unmarshaling config, %v", err)
	}

	expectedValue := "test"
	if tc.TestEnv != expectedValue {
		t.Errorf("Expected config to be %v, but got %v instead", expectedValue, tc.TestEnv)
	}
}
