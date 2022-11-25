package v1

import (
	"bytes"
	"encoding/json"
	"go-test/pkg/net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string, data any) *httptest.ResponseRecorder {
	jsonData, _ := json.Marshal(data)
	dataBytes := bytes.NewBuffer(jsonData)
	req, _ := http.NewRequest(method, path, dataBytes)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestPing(t *testing.T) {
	g := net.Init()
	w := performRequest(g, "GET", "/api/v1/ping", nil)

	assert.Equal(t, http.StatusOK, w.Code)

	response := net.Response{}
	// err := json.Unmarshal([]byte(w.Body.String()), &response)
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.Nil(t, err)
	assert.Equal(t, "pong", response.Message)
}
