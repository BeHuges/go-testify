package main

import (
	"net/http"
	"net/http/httptest"

	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenStatusOkAndNotEmpty(t *testing.T) {
	totalStatus := http.StatusOK
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	actualStatus := responseRecorder.Code
	require.Equal(t, totalStatus, actualStatus, "expected status code: %d, got %d", http.StatusOK, actualStatus)
	require.NotEmpty(t, responseRecorder.Body, "body is empty")
}

func TestMainHandlerWhenCityNotMoscow(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=5&city=syktyvkar", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	actualCount := len(strings.Split(body, ","))
	assert.Equal(t, totalCount, actualCount, "Status code:%d, wrong count value", responseRecorder.Code)
}
