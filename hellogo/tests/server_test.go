package tests

import (
	"testing"
	"net/http"
	"fmt"
)

func TestServer(t *testing.T)  {
	var server = http.Server{}
	fmt.Println(server.IdleTimeout)
	fmt.Println(server.ReadHeaderTimeout)
	fmt.Println(server.WriteTimeout)
}
