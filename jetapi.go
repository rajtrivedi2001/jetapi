package jetapi

import (
	"net/http"
)

type JetApi struct {
	key     string
	secret  string
	baseUrl string
	client  *http.Client
	token   *JetToken
}

func NewJetApi(key, secret, baseUrl string) *JetApi {
	return &JetApi{key, secret, baseUrl, &http.Client{}, nil}
}
