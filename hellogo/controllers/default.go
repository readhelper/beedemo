package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"net"
	"time"
"runtime"
	"bytes"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	println(getGID())
	println(Goid())
	c.Ctx.WriteString("hello,beego get")
}


func getGID() uint64 {
    b := make([]byte, 64)
    b = b[:runtime.Stack(b, false)]
    b = bytes.TrimPrefix(b, []byte("goroutine "))
    b = b[:bytes.IndexByte(b, ' ')]
    n, _ := strconv.ParseUint(string(b), 10, 64)
    return n
}

func Goid() int {
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

func (c *MainController) Post() {
	c.Ctx.WriteString("hello,beego post")
}

func (c *MainController) RestGet() {
	ret := httpDo("GET", "http://localhost:12345", nil)
	c.Ctx.WriteString(ret)
}

func (c *MainController) RestPost() {
	ret := httpDo("POST", "http://localhost:12345", nil)
	c.Ctx.WriteString(ret)
}

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, time.Second * 10)
}

func (c *MainController) RestDo() {
	transport := &http.Transport{
		Dial:              dialTimeout,
	}
	ret := httpDo("POST", "http://localhost:12345", transport)
	c.Ctx.WriteString(ret)
}

func (c *MainController) RestDefault() {
	transport := http.DefaultTransport
	ret := httpDo("POST", "http://localhost:12345", transport)
	c.Ctx.WriteString(ret)
}

func httpDo(method, url string, tr http.RoundTripper) string {
	client := &http.Client{}
	if tr != nil {
		client.Transport = tr
	}

	req, err := http.NewRequest(method, url, strings.NewReader(""))
	if err != nil {
		fmt.Println("http.NewRequest error", err)
	}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll error", err)
	}
	return string(body)
}
