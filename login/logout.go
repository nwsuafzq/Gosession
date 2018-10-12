package login

import (
	"github.com/gin-gonic/gin"
	"sessiontest/constant"
	"log"
)

func Logout(c *gin.Context){
	c.Header("Content-Type", "text/html; charset=utf-8")
	w:=c.Writer
	r:=c.Request
	session, _ := constant.Store.Get(r, "aaaa")
	name:= session.Values["username"].(string)
	log.Println(session.IsNew)
	session.Options.MaxAge = -1
	session.Save(r, w)
	c.String(200,"退出成功"+name+" 点击这里登录<a href='/login'>login</a>")
}
