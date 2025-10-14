package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/a-h/go-workshop/100/102/04-web-testing/models"
)

func New(baseURL string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: http.DefaultClient,
	}
}

type Client struct {
	BaseURL    string
	HTTPClient Doer
}

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

func (c *Client) UsersGet() (output models.UsersGetResponse, err error) {
	req, err := http.NewRequest(http.MethodGet, c.BaseURL+"/users", nil)
	if err != nil {
		return output, err
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return output, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return output, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return output, err
	}
	return output, nil
}

func (c *Client) UsersPost(input models.UsersPostRequest) (err error) {
	body, err := json.Marshal(input)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, c.BaseURL+"/users", bytes.NewReader(body))
	if err != nil {
		return err
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}
