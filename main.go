package main

import (
	"github.com/gin-gonic/gin"
	//"github.com/gorilla/sessions"
	"sessiontest/login"

)

// var Store = sessions.NewCookieStore([]byte("xes-something-very-secret"))



func main() {
	router := gin.Default()

	router.GET("/",login.Checklogin)
	router.GET("/login",login.Login)
	router.GET("/logout",login.Logout)

	router.Run(":82")

}
