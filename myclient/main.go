package myclient

import (
	"time"
	"net"
	"net/http"
	"strings"
	"fmt"
"io/ioutil"
)

func main() {
	var N = 100000
	for i := 0; i < N; i++ {
		//httpDo()
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