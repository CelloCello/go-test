package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-test/internal/models"
	"go-test/pkg/database"
	"go-test/pkg/net"
)

type UsersResp struct {
	net.Response
	Payload []models.User `json:"payload"`
}

type CreateUserReq struct {
	Name string `jaon:"name" binding:"required"`
}

type CreateUserResp struct {
	net.Response
}

// TestPing godoc
// @Summary User
// @Description Get user list
// @Tags User
// @Produce json
// @Success 200 {object} UsersResp
// @Router /users [get]
func GetUsers(c *gin.Context) {
	db := database.GetDB()
	users := []models.User{}

	result := db.Order("created_at desc").Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "failed to get users",
			"error": result.Error.Error(),
		})
		return
	}

	resp := UsersResp{}
	resp.Message = "success"
	resp.Payload = users
	c.JSON(http.StatusOK, resp)
}

// TestPing godoc
// @Summary User
// @Description Create a user
// @Tags User
// @Param data body CreateUserReq true  "User JSON"
// @Produce json
// @Success 200 {object} CreateUserResp
// @Router /users [post]
func CreateUser(c *gin.Context) {
	db := database.GetDB()
	user := CreateUserReq{}
	resp := CreateUserResp{}

	if err := c.ShouldBindJSON(&user); err != nil {
		resp.Message = "no binding data"
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	newUser := models.User{}
	newUser.Name = user.Name
	r := db.Create(&newUser)
	if r.Error != nil {
		resp.Message = "failed to create link"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	log.Printf("created user: id: %s, name: %s", newUser.ID, newUser.Name)

	resp.Message = "success"
	resp.Payload = user
	c.JSON(http.StatusOK, resp)
}
