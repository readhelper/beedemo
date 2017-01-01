package test

import (
	"testing"
	"log"
	"time"
	"crypto/x509"
	"io/ioutil"
	"fmt"
	"crypto/tls"
	"net/http"
	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

// TestBeego is a sample to run an endpoint test
func TestEtcdSSL(t *testing.T) {
	//x509.Certificate.
	pool := x509.NewCertPool()
	caCertPath := "C:/openssl/bin/root.crt"
	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)
	//pool.AddCert(caCrt)

	cliCrt, err := tls.LoadX509KeyPair("C:/openssl/bin/client.crt", "C:/openssl/bin/client.key")
	if err != nil {
		fmt.Println("Loadx509keypair err:", err)
		return
	}

	cfg := client.Config{
		Endpoints:[]string{"https://localhost:2379"},
		Transport:&http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      pool,
				Certificates: []tls.Certificate{cliCrt},
			},
		},
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	insertEtcd(c)
}
func insertEtcd( c client)  {
	kapi := client.NewKeysAPI(c)
	// set "/foo" key with "bar" value
	log.Print("Setting '/foo' key with 'bar' value")
	resp, err := kapi.Set(context.Background(), "/foo", "bar", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		// print common key info
		log.Printf("Set is done. Metadata is %q\n", resp)
	}
	// get "/foo" key's value
	log.Print("Getting '/foo' key value")
	resp, err = kapi.Get(context.Background(), "/foo", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		// print common key info
		log.Printf("Get is done. Metadata is %q\n", resp)
		// print value
		log.Printf("%q key has %q value\n", resp.Node.Key, resp.Node.Value)
	}
}