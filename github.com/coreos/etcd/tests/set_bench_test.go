package tests

import (
	"testing"
	"time"
	"math/rand"
	"github.com/coreos/etcd/Godeps/_workspace/src/github.com/coreos/go-etcd/etcd"
	"fmt"
)

func BenchmarkSet(b *testing.B) {
	client := etcd.NewClient([]string{"http://127.0.0.1:4001"})

	for i := 0; i < b.N; i++ {
		key := "/hao/key" + string(krand(20, 0))
		value := "val" + string(krand(20, 0))
		_, err := client.Set(key, value, 1000)
		if (err != nil) {
			fmt.Println("set error:", err)
		}
		if i % 5000 == 1 {
			println("[", i, "]", "key=", key, ",value=", value)
		}
	}
}
// 随机字符串
func krand(size int, kind int) []byte {
	rand.Seed(time.Now().UnixNano())
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	for i := 0; i < size; i++ {
		if is_all {
			// random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}