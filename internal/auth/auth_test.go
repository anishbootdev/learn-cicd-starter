package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetEmptyHeaderMapApiKey(t *testing.T) {
	_, err := GetAPIKey(http.Header{})
	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected: %v, got: %v", ErrNoAuthHeaderIncluded, err)
	}
}

func TestGetCorrectApi(t *testing.T) {
	apiKey, err := GetAPIKey(http.Header{
		"Authorization": []string{"ApiKey LOLOLOL"},
	})
	if !reflect.DeepEqual(apiKey, "LOLOLOL") {
		t.Fatalf("expected: %v, got: %v", "LOLOLOL", apiKey)
	}
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}
}
