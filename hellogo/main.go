package main

import (
	_ "hellogo/routers"
	"github.com/astaxie/beego"
	"fmt"
	_ "net/http/pprof"
	"time"
	"net/http"
)

type UnName struct {

}

func unDelete() {
	fmt.Print(".....")
}

func main() {
	/*
	f, _ := os.Create("profile_file")
	err := pprof.StartCPUProfile(f)  // 开始cpu profile，结果写到文件f中
	if (err != nil) {
		fmt.Println("err:", err)
	}
	defer pprof.StopCPUProfile()  // 结束profile

	go func() {
		ee := http.ListenAndServe("localhost:6060", nil)
		if (ee != nil) {
			fmt.Println("err:", ee)
		}
	}()

	getStats()
	*/
	beego.Run()
}

func getStats() {
	for i := 0; i < 10; i++ {
		printData()
		time.Sleep(10000)
	}
	//panic("")
}
func printData() {
	fmt.Println(time.Now().String())
}
