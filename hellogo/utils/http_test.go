package utils

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"net"
	"time"
	"testing"
)

var num = 0

var seconds = 0
var max = 100

func TestHttpGet(t *testing.T) {
	var tr = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   time.Second,
			KeepAlive: 1 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 1 * time.Second,
	}
	go timerBack()
	//var url = "http://localhost:4001/v2/keys/huawei/?recursive=true"
	var url = "http://www.readhelper.cn/"
	for num = 0; num < 10000 && seconds < max; num++ {
		httpGet(tr, url, num)
		//time.Sleep(time.Second)
	}
}

func timerBack() {
	timer1 := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timer1.C:
			seconds++
			if (seconds % 10 == 0) {
				fmt.Println("[", seconds, "] seconds process [", num, "] requests")
			}
			if num >= max {
				break
			}
		}
	}
}

func httpGet(tr *http.Transport, url string, index int) {
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := tr.RoundTrip(req)
	if err != nil {
		fmt.Println("http.Get error", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http.Get error", err)
	}
	url = string(body)
	if (index % 200 == 1) {
		fmt.Println("[", index, "] content is ", string(body)[0:10])
	}
}
