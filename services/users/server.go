package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer() http.Handler {
	users := NewUsers()

	router := gin.Default()
	router.GET("/users", withUsersCache(controllerGetUsers, users))
	router.POST("/users", withUsersCache(controllerCreateUser, users))

	return router
}
