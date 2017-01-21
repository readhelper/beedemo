package main

import (
	_ "hellogo/routers"
	"github.com/astaxie/beego"
	"fmt"
)

type UnName struct {

}

func unDelete() {
	fmt.Print(".....")
}

func main() {
	beego.Run()
}

