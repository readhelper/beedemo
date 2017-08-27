package controllers

import (
	"testing"
	"github.com/astaxie/beego/context"
	"net/http/httptest"
	"strings"
	"github.com/readhelper/beedemo/assert"
)

var main = MainController{}
var rw = httptest.NewRecorder()

func initMock() {
	main.Ctx = context.NewContext()
	rw = httptest.NewRecorder()
	main.Ctx.Reset(rw, httptest.NewRequest("get", "/", strings.NewReader("")))
}

func TestGet(t *testing.T) {
	initMock()

	main.Get()
	ret := rw.Body.String()
	assert.Equal(t, "hello,beego", ret)
}

func TestGetGID(t *testing.T) {
	gid := getGID()
	assert.Equal(t, true, gid > 0)
}

func TestRestGet(t *testing.T) {
	initMock()

	main.RestGet()
	ret := rw.Body.String()
	assert.Equal(t, "connection refused", ret)
}