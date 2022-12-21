package cmc_client

import (
	"fmt"
	"io"
	"net/http"
)

type CMCClient struct {
	client  *http.Client
	baseURL string
	apiKey  string
}

func New(baseURL, apiKey string) *CMCClient {
	return &CMCClient{client: &http.Client{}, baseURL: baseURL, apiKey: apiKey}
}

func (c *CMCClient) Get(path string) (body string, err error) {
	req, err := http.NewRequest("GET", c.baseURL+path, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", c.apiKey)
	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(resp.Status)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(respBody), nil
}
