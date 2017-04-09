package tests

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"net"
	"testing"
)

func TestRequetWithDial(t *testing.T) {
	var N = 10
	for i := 0; i < N; i++ {
		transport := &http.Transport{
			Dial:func(network, addr string) (net.Conn, error) {
				d := net.Dialer{Timeout: time.Second, }
				return d.Dial(network, addr)
			},
		}
		httpDo("POST", "http://192.168.0.4:12345", transport)
		time.Sleep(time.Millisecond * 100)
	}
	time.Sleep(time.Second * 1000)
}

func TestRequetWithEmpty(t *testing.T) {
	var N = 100000
	for i := 0; i < N; i++ {
		transport := &http.Transport{}
		httpDo("POST", "http://localhost:12345", transport)
		time.Sleep(time.Millisecond * 10)
	}
	time.Sleep(time.Second * 1000)
}
func TestRequetWithDefault(t *testing.T) {
	var N = 100000
	for i := 0; i < N; i++ {
		transport := http.DefaultTransport
		httpDo("POST", "http://localhost:12345", transport)
		time.Sleep(time.Millisecond * 10)
	}
	time.Sleep(time.Second * 1000)
}
func TestRequetToBeego(t *testing.T) {
	var N = 100000
	for i := 0; i < N; i++ {
		transport := &http.Transport{}
		httpDo("GET", "https://beego.me/", transport)
		time.Sleep(time.Millisecond * 10)
	}
	time.Sleep(time.Second * 1000)
}
func TestRequetToBaidu(t *testing.T) {
	var N = 100000
	for i := 0; i < N; i++ {
		transport := &http.Transport{}
		httpDo("GET", "http://www.baidu.com/", transport)
		time.Sleep(time.Millisecond * 10)
	}
	time.Sleep(time.Second * 1000)
}

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, time.Second * 10)
}

func httpDo(metod, url string, tr http.RoundTripper) {
	client := &http.Client{}
	client.Transport = tr
	req, err := http.NewRequest(metod, url, strings.NewReader(""))
	if err != nil {
		fmt.Println("http.NewRequest error", err)
		return
	}

	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println("client.Do error", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll error", err)
	}
	fmt.Println(string(body))
}