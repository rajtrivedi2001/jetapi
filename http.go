package jetapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"runtime"
)

func (a *JetApi) CreateRequest(method, path string, query url.Values, body io.Reader) *http.Request {
	reqUrl := a.baseUrl + path + "?" + query.Encode()

	req, err := http.NewRequest(method, reqUrl, body)
	if err != nil {
		log.Fatalln("Unable to create HTTP request: %v", err)
	}

	req.Header.Set("User-Agent", fmt.Sprintf("%s/github.com/kiasaki/jetapi", runtime.Version()))
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if a.token != nil {
		req.Header.Set("Authorization", fmt.Sprintf("bearer %s", a.token.Token))
	}

	return req
}

func (a *JetApi) CreateGetRequest(path string, query url.Values) *http.Request {
	return a.CreateRequest("GET", path, query, nil)
}

func (a *JetApi) CreatePostRequest(path string, query url.Values, body interface{}) (*http.Request, error) {
	bodyContent, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader := bytes.NewReader(bodyContent)
	return a.CreateRequest("POST", path, query, bodyReader), nil
}

func (a *JetApi) DoRequest(req *http.Request, model interface{}) error {
	//fmt.Println("Jet Api Requesting: " + req.URL.String())
	resp, err := a.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode > 399 {
		return errors.New(
			fmt.Sprintf("Error %d: %s", resp.StatusCode, string(respBody)))
	}

	// fmt.Println(string(respBody))
	return json.Unmarshal(respBody, model)
}
