package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer() http.Handler {
	users := NewUsers()

	router := gin.Default()
	router.POST("/users", withUsersCache(controllerGetUsers, users))
	router.POST("/users/create", withUsersCache(controllerCreateUser, users))

	return router
}
