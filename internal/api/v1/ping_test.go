package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-test/pkg/net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var gGin *gin.Engine
var gBasePath string

func performRequest(method, path string, data any) *httptest.ResponseRecorder {
	jsonData, _ := json.Marshal(data)
	dataBytes := bytes.NewBuffer(jsonData)
	req, _ := http.NewRequest(method, gBasePath+path, dataBytes)
	w := httptest.NewRecorder()
	gGin.ServeHTTP(w, req)
	return w
}

func netInit() {
	g := net.Init()
	basePath := "/api/v1"
	SetRoutes(basePath)
	gGin = g
	gBasePath = basePath
}

func TestPing(t *testing.T) {
	netInit()
	w := performRequest("GET", "/ping", nil)

	assert.Equal(t, http.StatusOK, w.Code)

	resp := PingResponse{}
	err := json.NewDecoder(w.Body).Decode(&resp)
	fmt.Println(w)
	assert.Nil(t, err)
	assert.Equal(t, "pong", resp.Message)
}
