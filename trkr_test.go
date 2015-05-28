package trkr

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var (
	mux    *http.ServeMux
	ts     *httptest.Server
	client *Client
)

func setup() {
	mux = http.NewServeMux()
	ts = httptest.NewServer(mux)
	client = NewClient(nil)
	client.Token = "blarstacoman"
	url, _ := url.Parse(ts.URL)
	client.BaseURL = url
}

func teardown() {
	ts.Close()
}

func testMethod(t *testing.T, r *http.Request, wants string) {

}

func TestNewClient(t *testing.T) {
	c := NewClient(nil)
	if got, want := c.BaseURL.String(), defaultBaseURL; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}
}

func TestNewRequest(t *testing.T) {
	c := NewClient(nil)
	c.Token = "blarstacoman"
	req, err := c.NewRequest("GET", "foo", nil)
	if err != nil {
		t.Errorf("Pre-mature error %v", err)
	}

	if got, want := req.URL.String(), defaultBaseURL+"foo"; got != want {
		t.Errorf("Request url is %v, wanted %v", got, want)
	}

	if req.Header.Get("X-TrackerToken") == "" {
		t.Error("X-TrackerToken is not defined")
	}

	if got, want := req.Header.Get("User-Agent"), "trkr-go/0.1"; got != want {
		t.Errorf("UserAgent is %#v, want %#v", got, want)
	}
}

func TestNewRequest_badURL(t *testing.T) {
	c := NewClient(nil)
	_, err := c.NewRequest("GET", "%AR", nil) // bad hex string

	if err == nil {
		t.Errorf("expecting error but got nil")
	}
}

func TestClient_Do(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"name": "test"}`)
	})

	s := new(Story)
	req, _ := client.NewRequest("GET", "/", nil)
	resp, err := client.Do(req, s)

	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if got, wants := resp.StatusCode, 200; got != wants {
		t.Errorf("Status code is %v, wants %v", got, wants)
	}

	if s == nil {
		t.Error("Object was supposed to be not nil")
	}
}

func TestClient_DoBlank(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, ``)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	resp, err := client.Do(req, nil)

	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if got, wants := resp.StatusCode, 200; got != wants {
		t.Errorf("Status code is %v, wants %v", got, wants)
	}
}
