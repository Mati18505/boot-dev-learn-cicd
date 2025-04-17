package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAPIKey(t *testing.T) {
	h := make(http.Header)
	
	h.Add("Authorization", "ApiKey 1234")

	_, err := GetAPIKey(h)
	assert.NoError(t, err)
}

func TestGetApiKeyNoHeader(t *testing.T) {
	_, err := GetAPIKey(nil)
	assert.Error(t, err)
}

func TestGetApiKeyNoAuthHeader(t *testing.T) {
	_, err := GetAPIKey(make(http.Header))
	assert.Error(t, err)
}

func TestGetApiKeyMallformed(t *testing.T) {
	h := make(http.Header)
	
	h.Add("Authorization", "ApiKey ")

	_, err := GetAPIKey(h)
	assert.NoError(t, err)
}