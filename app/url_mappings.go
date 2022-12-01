package app

import (
	"github.com/roziawilliams/bookstore_users-api/controllers/ping"
	"github.com/roziawilliams/bookstore_users-api/controllers/users"
)

func mapUrls() {
	//every Get req against /ping is handled by Ping
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)

	//router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", users.CreateUser)

}
