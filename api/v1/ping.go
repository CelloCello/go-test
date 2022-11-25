package v1

import (
	"net/http"

	"go-test/pkg/net"

	"github.com/gin-gonic/gin"
)

type PingResponse struct {
	net.Response
}

// TestPing godoc
// @Summary ping example
// @Description do ping
// @Tags AAA
// @Param Id query int true "Id" 参数 ：@Param 参数名 位置（query 或者 path或者 body） 类型 是否必需 注释
// @Produce json
// @Success 200 {object} PingResponse
// @Router /ping [get]
func TestPing(c *gin.Context) {
	resp := PingResponse{}
	resp.Message = "pong"
	c.JSON(http.StatusOK, resp)
}
