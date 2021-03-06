package dao

import (
	"fmt"
	"github.com/coreos/etcd/client"
	"time"
)

var kapi client.KeysAPI

func getconn() {
	cfg := client.Config{
		Endpoints:               []string{"http://127.0.0.1:4001"},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		fmt.Println("etcd cfg error:", err)
	}
	kapi = client.NewKeysAPI(c)

}

func GetKeysAPI() client.KeysAPI {
	if kapi == nil {
		getconn()
	}
	return kapi
}

func SetKeysAPI(api client.KeysAPI) {
	kapi = api
}
