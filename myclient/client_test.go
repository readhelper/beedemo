package myclient

import (
	"net/http"
	"time"
	"net"
	"testing"
)

func testRequetWithDial(t *testing.T) {
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
	time.Sleep(time.Second * 60)
}

func testRequetWithEmpty(t *testing.T) {
	var N = 100000
	for i := 0; i < N; i++ {
		transport := &http.Transport{}
		httpDo("POST", "http://localhost:12345", transport)
		time.Sleep(time.Millisecond * 10)
	}
	time.Sleep(time.Second * 60)
}
func testRequetWithDefault(t *testing.T) {
	var N = 100000
	for i := 0; i < N; i++ {
		transport := http.DefaultTransport
		httpDo("POST", "http://localhost:12345", transport)
		time.Sleep(time.Millisecond * 10)
	}
	time.Sleep(time.Second * 60)
}
func testRequetToBeego(t *testing.T) {
	var N = 100000
	for i := 0; i < N; i++ {
		transport := &http.Transport{}
		httpDo("GET", "https://beego.me/", transport)
		time.Sleep(time.Millisecond * 10)
	}
	time.Sleep(time.Second * 60)
}
func testRequetToBaidu(t *testing.T) {
	var N = 100000
	for i := 0; i < N; i++ {
		transport := &http.Transport{}
		httpDo("GET", "http://www.baidu.com/", transport)
		time.Sleep(time.Millisecond * 10)
	}
	time.Sleep(time.Second * 60)
}
