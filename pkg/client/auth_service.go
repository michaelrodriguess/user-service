package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type AuthClient struct {
	baseURL string
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type RegisterResponse struct {
	Token string `json:"token"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func NewAuthClient() *AuthClient {
	authServiceURL := os.Getenv("AUTH_SERVICE_URL")
	if authServiceURL == "" {
		authServiceURL = "http://localhost:8080"
	}
	return &AuthClient{
		baseURL: authServiceURL,
	}
}

func (c *AuthClient) GenerateToken(email, password, role string) (string, error) {
	reqBody := RegisterRequest{
		Email:    email,
		Password: password,
		Role:     role,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(fmt.Sprintf("%s/register", c.baseURL),
		"application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return "", errors.New("failed to generate token, status: " + resp.Status)
	}

	var respData RegisterResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return "", err
	}

	return respData.Token, nil
}
