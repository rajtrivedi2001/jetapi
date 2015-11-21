package jetapi

import (
	"net/url"
	"time"
)

type JetToken struct {
	Token  string    `json:"id_token"`
	Expiry time.Time `json:"expires_on"`
}

func (a *JetApi) GetToken() (*JetToken, error) {
	req, err := a.CreatePostRequest("/token", url.Values{}, map[string]string{
		"user": a.key,
		"pass": a.secret,
	})
	if err != nil {
		return nil, err
	}

	var token *JetToken
	return token, a.DoRequest(req, &token)
}

func (a *JetApi) EnsureValidToken() error {
	if a.token == nil || time.Now().After(a.token.Expiry) {
		token, err := a.GetToken()
		if err != nil {
			return err
		}
		a.token = token
	}
	return nil
}
