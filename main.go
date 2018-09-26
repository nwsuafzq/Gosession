package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"

)

var Store = sessions.NewCookieStore([]byte("xes-something-very-secret"))



func main() {
	router := gin.Default()

	router.GET("/",checklogin)
	router.GET("/login",login)
	router.GET("/logout",logout)

	router.Run(":82")

}
