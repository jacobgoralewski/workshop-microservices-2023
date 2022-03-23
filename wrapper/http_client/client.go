package http_client

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpClient struct {
	client *http.Client
}

func NewHttpClient() *HttpClient {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	return &HttpClient{
		client: &http.Client{
			Timeout:   10 * time.Second,
			Transport: t,
		},
	}
}

func (hc HttpClient) Get(url string, container interface{}) error {
	response, err := hc.client.Get(url)
	if err != nil {
		return errors.Wrapf(err, "failed to do GET request to %s", url)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.Wrapf(err, "failed to do read body from: %s", url)
	}
	defer response.Body.Close()

	err = json.Unmarshal(body, container)
	if err != nil {
		return errors.Wrapf(err, "failed to do unmarshal body from: %s", url)
	}

	return nil
}
