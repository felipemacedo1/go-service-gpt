package client

import (
	"fmt"
	"net/http"
)

type JavaClient struct {
	baseURL string
}

func NewJavaClient(baseURL string) *JavaClient {
	return &JavaClient{
		baseURL: baseURL,
	}
}

func (c *JavaClient) ValidateToken(token string) (bool, error) {
	req, err := http.NewRequest("GET", c.baseURL, nil)
	if err != nil {
		return false, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}
	return false, fmt.Errorf("invalid token. Status code: %d", resp.StatusCode)
}
