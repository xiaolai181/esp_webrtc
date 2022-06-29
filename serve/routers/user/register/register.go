package register

import (
	"errors"
	"esp_webrtc/models"
	"net/http"
	"strconv"

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

func Login(c *gin.Context) {
	var data Register_data
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	if ok := models.Vaild_User(data.Name, data.Password); ok {
		c.JSON(http.StatusOK, gin.H{"msg": "ok"})

	}
}

func GetUserEmailById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "id is empty"})
		return
	}
	user_id, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	user := models.GetUserById(user_id)
	c.JSON(http.StatusOK, gin.H{"email": user.Email})
}
