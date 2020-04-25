package app

import (
	"github.com/wi7sonjoseph/bookstore_users-api/controllers/ping"
	"github.com/wi7sonjoseph/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/user/:user_id", users.GetUser) //Need to fix this
	router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}
