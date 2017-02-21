package chans

import (
	"fmt"
	"math/rand"
	"time"
	"testing"
)

var go_on bool = true

func productor(channel chan <- string) {
	for ; go_on; {
		message :=fmt.Sprintf("%v", rand.Float64())
		fmt.Println("productor:",message)

		channel <- message
		time.Sleep(time.Second * time.Duration(1))
	}
}

func customer(channel <-chan string) {
	for ; go_on; {
		message := <-channel // 此处会阻塞, 如果信道中没有数据的话
		fmt.Println("customer:",message)
	}
}

func TestChannels(t *testing.T) {
	channel := make(chan string, 5) // 定义带有5个缓冲区的信道(当然可以是其他数字)
	go productor(channel) // 将 productor 函数交给协程处理, 产生的结果传入信道中
	go customer(channel) // 主线程从信道中取数据


	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		i++
	}
	go_on = false
}
