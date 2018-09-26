package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

func login(c *gin.Context){
	w:=c.Writer
	r:=c.Request
	c.Header("Content-Type", "text/html; charset=utf-8")
	session, err := Store.Get(r, "zqxes_get_name_session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session.Values["username"]!=nil {
		fmt.Println("登陆过了")
		c.Redirect(302,"http://10.99.2.212:82")
	}else {

		session.Values["username"] = "dd"
		session.Save(r, w)
		c.String(200,"登陆成功"+ session.Values["username"].(string) +"<a href='/logout'>exit</a>")
	}
}
