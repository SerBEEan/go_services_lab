package users

import (
	"github.com/gin-gonic/gin"
)

type controllerCallback func(c *gin.Context, users Users)

func withUsersCache(controller controllerCallback, users Users) gin.HandlerFunc {
	return func(c *gin.Context) {
		controller(c, users)
	}
}
