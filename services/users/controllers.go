package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func controllerCreateUser(c *gin.Context, users Users) {
	info := userInfo{}

	if err := c.BindJSON(&info); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var id int = users.CreateUser(info)
	c.JSON(http.StatusAccepted, gin.H{"id": id})
}

func controllerGetUsers(c *gin.Context, users Users) {
	c.JSON(http.StatusOK, users.GetUsers())
}
