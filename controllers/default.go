package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"runtime"
	"strconv"
	"strings"
	"github.com/astaxie/beego/httplib"
	"time"
	"github.com/readhelper/beedemo/logger"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Ctx.WriteString("hello,beego")
}

func (c *MainController) GetGID() {
	var id = strconv.Itoa(getGID())
	c.Ctx.WriteString(id)
}

func getGID() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recover:panic info:%v", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func (c *MainController) RestGet() {
	req := httplib.Get("http://localhost:12345").Debug(true).SetTimeout(3 * time.Second, 3 * time.Second)
	ret, err := req.String()
	if err != nil {
		logger.Error(" httplib.Get err:", err)
		c.Ctx.WriteString("connection refused")
	}
	c.Ctx.WriteString(ret)
}
