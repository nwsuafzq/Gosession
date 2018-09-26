package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"

)

func checklogin(c *gin.Context) {
	w:=c.Writer
	r:=c.Request
	c.Header("Content-Type", "text/html; charset=utf-8")
	session, err := Store.Get(r, "zqxes_get_name_session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	name, ok := session.Values["username"].(string)
	//session.Values["username"] = "dd"
	//session.Save(r, w)
	fmt.Printf("aaaaaaaa"+"%t\n",ok)
	fmt.Println("1")

	if ok {
		fmt.Printf("bbbbbbbbbb"+"%t\n",ok)
		fmt.Println("2")
		c.String(200,"hello, "+string(name)+"<a href='/logout'>退出登录</a>")

	}else{
		fmt.Println("3")
		//io.WriteString(w,"请登录!!")
		c.String(200,"你还没登录 点击这里登录, "+"<a href='/login'>LOGIN</a>")

		//c.Redirect(302,"http://10.99.2.212:82/login")
	}
}

