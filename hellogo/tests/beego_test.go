package tests

import (
	"path/filepath"
	"runtime"
	"testing"
	"github.com/astaxie/beego"
	"fmt"
	"bytes"
	"mime/multipart"
	"net/http"
	"io/ioutil"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	apppath = filepath.Join(apppath, "conf", "app_win.conf")
	//beego.BeeApp.LoadAppConfig("ini",apppath)
	//beego.TestBeegoInit(apppath)
}

// TestBeego is a sample to run an endpoint test
func TestBeegoUpload(t *testing.T) {
	var buffer bytes.Buffer

	w := multipart.NewWriter(&buffer)
	//  Write fields and files w.CreateFormField( " input1 " )
	w.WriteField("input1", "value1")

	write,err :=ioutil.TempFile("c:/","rnd1.txt")
	fmt.Println(err)

	write.WriteString("ssss")
	write.Close()

	w.CreateFormFile("file", "c:/rnd.txt")

	fmt.Println(w.FormDataContentType())
	fmt.Println(string(buffer.Bytes()))

	resp, err := http.Post("http://www.baidu.com/", w.FormDataContentType(), &buffer)
	fmt.Println(err)
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}
func TestBeegoHandle(t *testing.T) {

}

// TestBeego is a sample to run an endpoint test
func TestBeego(t *testing.T) {
	var value string
	value = beego.AppConfig.String("appname")
	fmt.Println(value)

	beego.AppConfig.Set("appname", "newappname")
	value = beego.AppConfig.String("appname")
	fmt.Println(value)

	value = beego.AppConfig.String("myset")
	fmt.Println(value)

	beego.AppConfig.Set("myset", "newsetvalue")
	value = beego.AppConfig.String("myset")
	fmt.Println(value)
}