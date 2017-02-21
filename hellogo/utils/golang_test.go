package utils

import (
	"testing"
	"fmt"
)

func TestGolang(t *testing.T){
	arr := []int{1, 2, 3, 4}
	fmt.Println(arr)

	arr[1] = 11
	fmt.Println(arr)

	change(arr)
	fmt.Println(arr)
}

func change(ints []int)  {
	ints[1] = 22
	//fmt.Println(ints)
}

