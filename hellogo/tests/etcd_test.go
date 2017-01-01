package test

import (
	"testing"
	"log"
        "github.com/coreos/etcd/client"
        "time"
        "golang.org/x/net/context"
)

// TestBeego is a sample to run an endpoint test
func TestEtcd(t *testing.T) {
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
    // set "/foo" key with "bar" value
    log.Print("Setting '/haolipeng' key with 'bar' value")
    resp, err := kapi.Set(context.Background(), "/haolipeng", "bar", nil)
    if err != nil {
        log.Fatal(err)
    } else {
        // print common key info
        log.Printf("Set is done. Metadata is %q\n", resp)
    }
    // get "/foo" key's value
    log.Print("Getting '/haolipeng' key value")
    resp, err = kapi.Get(context.Background(), "/haolipeng", nil)
    if err != nil {
        log.Fatal(err)
    } else {
        // print common key info
        log.Printf("Get is done. Metadata is %q\n", resp)
        // print value
        log.Printf("%q key has %q value\n", resp.Node.Key, resp.Node.Value)
    }
}
