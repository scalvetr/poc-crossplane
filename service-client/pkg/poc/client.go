package poc

import (
	"io"
	"net/http"
)

type Client struct {
	BaseURL string
	Tenant  string
	http.Client
}

func (c *Client) sendRequest(method string, url string, body io.Reader, headers map[string]string) (resp *http.Response, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-tenant", c.Tenant)
	for key, element := range headers {
		req.Header.Set(key, element)
	}
	return c.Do(req)
}
func (c *Client) Get(url string) (resp *http.Response, err error) {
	return c.sendRequest("GET", url, nil, nil)
}
func (c *Client) Post(url string, body io.Reader) (resp *http.Response, err error) {
	return c.sendRequest("POST", url, body, map[string]string{"Content-Type": "application/json"})
}
func (c *Client) Put(url string, body io.Reader) (resp *http.Response, err error) {
	return c.sendRequest("PUT", url, body, map[string]string{"Content-Type": "application/json"})
}
func (c *Client) Delete(url string) (resp *http.Response, err error) {
	return c.sendRequest("DELETE", url, nil, nil)
}
