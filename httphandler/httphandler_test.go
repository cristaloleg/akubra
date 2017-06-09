package httphandler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnEntityTooLargeCode(t *testing.T) {
	request := httptest.NewRequest("POST", "http://somepath", nil)
	request.Header.Set("Content-Length", "4096")
	handler := &Handler{bodyMaxSize: 1024, maxConcurrentRequests: 10}
	writer := httptest.NewRecorder()
	handler.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusRequestEntityTooLarge, writer.Code)
}

func TestShouldReturnBadRequestOnUnparsableContentLengthHeader(t *testing.T) {
	request := httptest.NewRequest("POST", "http://somepath", nil)
	request.Header.Set("Content-Length", "strange-content-header")
	handler := &Handler{bodyMaxSize: 1024, maxConcurrentRequests: 10}
	writer := httptest.NewRecorder()
	handler.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusBadRequest, writer.Code)
}

func TestShouldReturnServiceNotAvailableOnTooManyRequests(t *testing.T) {
	request := httptest.NewRequest("GET", "http://somepath", nil)
	handler := &Handler{bodyMaxSize: 1024, maxConcurrentRequests: 0}
	writer := httptest.NewRecorder()
	handler.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusServiceUnavailable, writer.Code)
}
