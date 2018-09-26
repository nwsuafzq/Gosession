package main

import "github.com/gin-gonic/gin"

func logout(c *gin.Context){
	c.Header("Content-Type", "text/html; charset=utf-8")
	w:=c.Writer
	r:=c.Request
	session, _ := Store.Get(r, "zqxes_get_name_session")
	name:= session.Values["username"].(string)

	session.Options.MaxAge = -1
	session.Save(r, w)
	c.String(200,"退出成功"+name+" 点击这里登录<a href='/login'>login</a>")
}
