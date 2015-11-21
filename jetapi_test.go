package jetapi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewJetApiHasKeyAndSecret(t *testing.T) {
	api := NewJetApi("key", "secret", "baseurl")
	assert.NotEqual(t, api.key, "api key should be set")
	assert.NotEqual(t, api.secret, "api secret should be set")
	assert.NotEqual(t, api.baseUrl, "api baseUrl should be set")
}

func TestNewJetApiSetsClient(t *testing.T) {
	api := NewJetApi("", "", "")
	assert.NotNil(t, api.client, "api http client is not nil")
}

func TestNewJetApiHasNoToken(t *testing.T) {
	api := NewJetApi("", "", "")
	assert.Nil(t, api.token, "api token is nil initially")
}
