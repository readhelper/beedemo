package proxys

import (
	"net/http"
	"io/ioutil"
	"strings"
	"time"
)

func GetRemoteObject(key string) (sc int, ret string, err error) {
	time.Sleep(10 * time.Millisecond)
	return 200, "body", nil
	//return httpDo("http://baidu.com/?t=" + key)
}

func httpDo(url string) (sc int, ret string, err error) {
	client := &http.Client{}
	client.Transport = &http.Transport{
		IdleConnTimeout: 10 * time.Second,
	}
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return resp.StatusCode, string(body), nil
}
