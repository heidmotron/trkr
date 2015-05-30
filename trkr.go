package trkr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://www.pivotaltracker.com/services/v5/"
)

// Client responsible for communicating with the Tracker API
type Client struct {
	client *http.Client

	BaseURL *url.URL
	Token   string
}

// NewClient creates a Client with defaults
func NewClient(c *http.Client) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)

	if c == nil {
		c = http.DefaultClient
	}

	return &Client{
		client:  c,
		BaseURL: baseURL,
	}
}

// NewRequest generates an http.Request with all the correct headers filled out
func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	pathURL, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	rURL := c.BaseURL.ResolveReference(pathURL)
	req, _ := http.NewRequest(method, rURL.String(), buf)

	if c.Token != "" {
		req.Header.Add("X-TrackerToken", c.Token)
	}

	req.Header.Add("User-Agent", "trkr-go/0.1")
	return req, nil
}

// Do Exposes http.Client's Do
func (c *Client) Do(r *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if c := resp.StatusCode; c < 200 && c > 299 {
		return resp, fmt.Errorf("Status code invalid: %d", c)
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		return resp, err
	}

	return resp, nil

}
