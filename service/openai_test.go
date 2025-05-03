package service

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockHTTPClient is a mock implementation of the HTTPClient interface for testing.
type MockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func TestOpenAIService_SendMessage(t *testing.T) {
	t.Run("Successful response", func(t *testing.T) {
		mockResponse := `{"choices": [{"message": {"content": "Test response"}}], "usage":{"completion_tokens": 10}}`
		mockHTTPClient := &MockHTTPClient{
			DoFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBufferString(mockResponse)),
				}, nil
			},
		}

		openAIService := NewOpenAIService("test-api-key", mockHTTPClient)
		response, err := openAIService.SendMessage("Test message")

		assert.NoError(t, err)
		assert.Equal(t, "Test response", response.Choices[0].Message.Content)
		assert.Equal(t, 10, response.Usage.CompletionTokens)
	})

	t.Run("Error response", func(t *testing.T) {
		mockHTTPClient := &MockHTTPClient{
			DoFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusInternalServerError,
					Body:       io.NopCloser(bytes.NewBufferString(`{"error": "Internal Server Error"}`)),
				}, nil
			},
		}

		openAIService := NewOpenAIService("test-api-key", mockHTTPClient)
		_, err := openAIService.SendMessage("Test message")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to send message to OpenAI")
	})
	
		t.Run("HTTP Client Error", func(t *testing.T) {
		mockHTTPClient := &MockHTTPClient{
			DoFunc: func(req *http.Request) (*http.Response, error) {
				return nil, errors.New("network error")
			},
		}

		openAIService := NewOpenAIService("test-api-key", mockHTTPClient)
		_, err := openAIService.SendMessage("Test message")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "network error")
	})
}