package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandleChatIntegration(t *testing.T) {
	// Create a mock HTTP server to simulate the OpenAI API
	mockOpenAIServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate a successful response from OpenAI
		response := `{"choices": [{"message": {"content": "Mock response from OpenAI"}}]}`
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	defer mockOpenAIServer.Close()

	// Replace the OpenAI API URL with the mock server URL
	openAIURL = mockOpenAIServer.URL

	// Create a new Echo instance
	e := echo.New()

	// Create a chat handler
	chatHandler := &ChatHandler{}

	// Create a request body
	reqBody := map[string]string{
		"message": "Test message",
	}
	jsonReq, _ := json.Marshal(reqBody)

	// Create a request to the /chat endpoint
	req := httptest.NewRequest(http.MethodPost, "/chat", bytes.NewBuffer(jsonReq))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the HandleChat function
	err := chatHandler.HandleChat(c)

	// Assert that there are no errors
	assert.NoError(t, err)

	// Assert that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, rec.Code)

	// Assert that the response body is not empty
	assert.NotEmpty(t, rec.Body.String())

	var responseData map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &responseData)
	if err != nil {
		fmt.Println("Erro ao fazer o unmarshal do corpo da resposta:", err)
		return
	}
	if _, ok := responseData["response"]; !ok {
		t.Errorf("Chave 'response' não encontrada no JSON")
		return
	}
}