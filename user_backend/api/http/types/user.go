package types

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PostRegisterUserHandlerRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func CreatePostRegisterUserHandlerRequest(r *http.Request) (*PostRegisterUserHandlerRequest, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while decoding json: %v", err)
	}

	defer r.Body.Close()

	var req PostRegisterUserHandlerRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		return nil, fmt.Errorf("error while decoding json: %v", err)
	}
	return &req, nil
}

type PostRegisterUserHandlerResponse struct {
	SessionID string `json:"sessionID"`
}
