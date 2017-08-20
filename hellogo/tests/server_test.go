package tests

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	var server = http.Server{}
	fmt.Println(server.IdleTimeout)
	fmt.Println(server.ReadHeaderTimeout)
	fmt.Println(server.WriteTimeout)
}
