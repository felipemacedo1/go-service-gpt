package client

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type JavaClient struct {
	baseURL    string
	logger *logrus.Logger
}

func NewJavaClient(baseURL string, logger *logrus.Logger) *JavaClient {
	return &JavaClient{
		baseURL:    baseURL,
		logger: logger,
	}
}

func (c *JavaClient) ValidateToken(token string) (bool, error) {
	c.logger.Info("Validating token with Java service")
	req, err := http.NewRequest("GET", c.baseURL, nil)
	if err != nil {
		c.logger.WithError(err).Error("Error creating request to Java service")
		return false, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		c.logger.WithError(err).Error("Error sending request to Java service")
		return false, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		c.logger.WithField("status", resp.StatusCode).Info("Token validated successfully by Java service")
		return true, nil
	}
	err = fmt.Errorf("invalid token. Status code: %d", resp.StatusCode)
	c.logger.WithError(err).WithField("status", resp.StatusCode).Error("Token validation failed by Java service")
	return false, err
}
