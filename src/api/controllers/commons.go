package controllers

import (
	"encoding/json"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
)

// RequestBody ...
type RequestBody map[string]interface{}

// RequestHeaders ...
type RequestHeaders map[string]string

//PerformRequest used for make a request in tests
func PerformRequest(method, target string, body RequestBody, headers RequestHeaders, engine *gin.Engine) *httptest.ResponseRecorder {
	requestBody := ""
	if body != nil {
		requestBodyBytes, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}
		requestBody = string(requestBodyBytes)
	}
	payload := strings.NewReader(requestBody)
	req := httptest.NewRequest(method, target, payload)

	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	res := httptest.NewRecorder()
	engine.ServeHTTP(res, req)
	return res
}
