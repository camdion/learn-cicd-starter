package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	h := http.Header{}
	_, err := GetAPIKey(h)
	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("No auth header expected but got: %v", err)
	}
	h.Add("Authorization", "foobar")
	_, err = GetAPIKey(h)
	if err == nil {
		t.Fatal("Expected error: Authorization not formatted with ApiKey")
	}
	h.Set("Authorization", "ApiKey foobar")
	key, err := GetAPIKey(h)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if key != "foobar" {
		t.Fatalf("Unexpected key returned: %v", key)
	}
	fmt.Printf("TestGetAPIKey: ApiKey: %v\n", key)
}
