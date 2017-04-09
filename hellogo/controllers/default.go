package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"net"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Ctx.WriteString("hello,beego get")
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
