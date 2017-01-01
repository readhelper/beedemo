package test

import (
	"fmt"
	"net/http"
	"crypto/tls"
	"testing"
	"crypto/x509"
	"io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of https service in golang!")
}

func TestHttps(t *testing.T) {
	//x509.Certificate.
	pool := x509.NewCertPool()
	//caCertPath := "etcdcerts/ca.crt"
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

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}
	client := &http.Client{Transport: tr}
	//resp, err := client.Get("https://localhost:8081")
	resp, err := client.Get("https://localhost:2379/v2/keys/")
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
