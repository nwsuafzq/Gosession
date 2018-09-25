package main

import "github.com/gin-gonic/gin"

func logout(c *gin.Context){
	c.Header("Content-Type", "text/html; charset=utf-8")
	w:=c.Writer
	r:=c.Request
	session, _ := store.Get(r, "zqxes_get_name_session")
	name:= session.Values["username"].(string)

	session.Options.MaxAge = -1
	session.Save(r, w)
	c.String(200,"退出成功"+name)
}
