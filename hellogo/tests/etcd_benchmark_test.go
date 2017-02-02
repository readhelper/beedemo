package test

import (
	"testing"
	"log"
	"github.com/coreos/etcd/client"
	"time"
	"golang.org/x/net/context"
	"encoding/base64"
	"crypto/md5"
	"encoding/hex"
	"io"
	"crypto/rand"
)

func BenchmarkLoops(b *testing.B) {
	cfg := client.Config{
		Endpoints:               []string{"http://127.0.0.1:4001"},
		Transport:               client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi := client.NewKeysAPI(c)
	// 必须循环 b.N 次 。 这个数字 b.N 会在运行中调整，以便最终达到合适的时间消耗。方便计算出合理的数据。 （ 免得数据全部是 0 ）
	for i := 0; i < b.N; i++ {
		loopInsertData(kapi, i)
	}
}

// TestBeego is a sample to run an endpoint test
func loopInsertData(kapi client.KeysAPI, cnt int) {
	key := getGuid()
	value := getGuid()
	resp, err := kapi.Set(context.Background(), "/" + key, value, nil)
	if err != nil {
		log.Fatal(err)
	} else if (cnt / 10000 == 1) {
		log.Printf("Set [", cnt, "] is done. Metadata is %q\n", resp)
	}
}

func getGuid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	s := base64.URLEncoding.EncodeToString(b)
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
