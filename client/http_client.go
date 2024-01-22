package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func newBasicClient(clientID string, clientSecret string) *basicClient {
	return &basicClient{
		clientID:     clientID,
		clientSecret: clientSecret,
		client: http.Client{
			Timeout: time.Second * 5,
		},
	}
}

type basicClient struct {
	clientID     string
	clientSecret string

	client http.Client
}

func (c *basicClient) Post(uri string, payload url.Values) (res *http.Response, body string, err error) {
	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewReader([]byte(payload.Encode())))
	if err != nil {
		return
	}

	req.SetBasicAuth(c.clientID, c.clientSecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err = c.client.Do(req)
	if err != nil {
		return
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	res.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))

	return res, string(bodyBytes), err
}
