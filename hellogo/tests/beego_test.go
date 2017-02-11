package tests

import (
	"path/filepath"
	"runtime"
	"testing"
	"github.com/astaxie/beego"
	"fmt"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	apppath = filepath.Join(apppath,"conf","app_win.conf")
	beego.LoadAppConfig("ini",apppath)
	//beego.TestBeegoInit(apppath)
}

// TestBeego is a sample to run an endpoint test
func TestBeego(t *testing.T) {
	var value string
	value =beego.AppConfig.String("appname")
	fmt.Println(value)

	beego.AppConfig.Set("appname","newappname")
	value =beego.AppConfig.String("appname")
	fmt.Println(value)

	value =beego.AppConfig.String("myset")
	fmt.Println(value)

	beego.AppConfig.Set("myset","newsetvalue")
	value =beego.AppConfig.String("myset")
	fmt.Println(value)

}
