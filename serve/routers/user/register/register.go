package register

import (
	"errors"
	"esp_webrtc/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Register_data struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func Register(c *gin.Context) {

	var data Register_data
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	if !models.Create_user(data.Name, data.Password, data.Email) {
		err := errors.New("create user error")
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})

}
