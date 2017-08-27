package myclient

import (
	"testing"
	"net/http"
	"time"
)

func TestRequetOk1(t *testing.T) {
	var N = 10
	for i := 0; i < N; i++ {
		transport := &http.Transport{
			IdleConnTimeout:       2 * time.Second,
		}
		httpDo("GET", "http://192.168.0.4:12345", transport)
		time.Sleep(time.Millisecond * 100)
	}
	time.Sleep(time.Second * 60)
}

func TestRequetOk2(t *testing.T) {
	var N = 10
	for i := 0; i < N; i++ {
		transport := &http.Transport{
			DisableKeepAlives:true,
		}
		httpDo("GET", "http://192.168.0.4:12345", transport)
		time.Sleep(time.Millisecond * 100)
	}
	time.Sleep(time.Second * 60)
}
func TestRequetOk3(t *testing.T) {
	var N = 10
	transport := &http.Transport{}
	for i := 0; i < N; i++ {
		httpDo("GET", "http://192.168.0.4:12345", transport)
		time.Sleep(time.Second * 3)
	}
	time.Sleep(time.Second * 60)
}
