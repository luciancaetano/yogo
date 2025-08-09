package apiclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// ==========================
// Client
// ==========================

type Client struct {
	baseURL    string
	httpClient *http.Client
	headers    map[string]string
}

// NewClient cria um cliente para a API.
// Exemplo: c := NewClient("https://api.exemplo.com", nil)
func NewClient(baseURL string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{
		baseURL:    strings.TrimRight(baseURL, "/"),
		httpClient: httpClient,
		headers:    make(map[string]string),
	}
}

// SetHeader adiciona/atualiza um header global no cliente
func (c *Client) SetHeader(key, value string) {
	c.headers[key] = value
}

// ==========================
// GetUser
// ==========================

type GetUserResponse200 struct {
	Email string `json:"email"`
	Id    int    `json:"id"`
	Name  string `json:"name"`
}

type GetUserResponse404 struct {
	Error string `json:"error"`
}

// GetUser chama o endpoint GET /users/{id}.
func (c *Client) GetUser(ctx context.Context, id any) (any, int, error) {
	url := fmt.Sprintf("%s/users/%v", c.baseURL, id)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return nil, 0, err
	}

	for k, v := range c.headers {
		req.Header.Set(k, v)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {

	case 200:
		var data200 GetUserResponse200
		if err := json.NewDecoder(resp.Body).Decode(&data200); err != nil {
			return nil, 200, err
		}
		return &data200, 200, nil

	case 404:
		var data404 GetUserResponse404
		if err := json.NewDecoder(resp.Body).Decode(&data404); err != nil {
			return nil, 404, err
		}
		return &data404, 404, nil

	default:
		body, _ := io.ReadAll(resp.Body)
		return string(body), resp.StatusCode, nil
	}
}

// ==========================
// CreateUser
// ==========================

type CreateUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type CreateUserResponse201 struct {
	Id int `json:"id"`
}

type CreateUserResponse400 struct {
	Error string `json:"error"`
}

// CreateUser chama o endpoint POST /users.
func (c *Client) CreateUser(ctx context.Context, body CreateUserRequest) (any, int, error) {
	url := fmt.Sprintf("%s/users", c.baseURL)

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, 0, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, 0, err
	}

	for k, v := range c.headers {
		req.Header.Set(k, v)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {

	case 201:
		var data201 CreateUserResponse201
		if err := json.NewDecoder(resp.Body).Decode(&data201); err != nil {
			return nil, 201, err
		}
		return &data201, 201, nil

	case 400:
		var data400 CreateUserResponse400
		if err := json.NewDecoder(resp.Body).Decode(&data400); err != nil {
			return nil, 400, err
		}
		return &data400, 400, nil

	default:
		body, _ := io.ReadAll(resp.Body)
		return string(body), resp.StatusCode, nil
	}
}
