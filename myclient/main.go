package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"net"
)

func main() {
	var N = 100000
	for i := 0; i < N; i++ {
		httpDo()
		time.Sleep(time.Millisecond * 10)
	}
	time.Sleep(time.Second * 1000)
}

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, time.Second * 10)
}

func httpDo() {
	client := &http.Client{}
	client.Transport = &http.Transport{
		Dial:              dialTimeout,
	}
	req, err := http.NewRequest("GET", "http://localhost:12345", strings.NewReader(""))
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

