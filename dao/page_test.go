package dao_test

import (
	"fmt"
	"golang.org/x/net/context"
	"github.com/readhelper/beedemo/dao"
	"strings"
	"testing"
	"github.com/readhelper/beedemo/assert"
)

/*
func TestFindNextSubDir(t *testing.T) {
	var dirMap = map[string]int{}
	dirMap["003"] = 0
	dirMap["002"] = 0
	dirMap["001"] = 0

	var nextDir string = ""
	nextDir = findNextSubDir(dirMap, nextDir)
	assert.Equal(t, "003", nextDir)

	nextDir = findNextSubDir(dirMap, nextDir)
	assert.Equal(t, "002", nextDir)

	nextDir = findNextSubDir(dirMap, nextDir)
	assert.Equal(t, "001", nextDir)

	nextDir = findNextSubDir(dirMap, nextDir)
	assert.Equal(t, "", nextDir)
}

func TestReverseInts(t *testing.T) {
	oldInts := []string{"a1", "a2", "a3" }
	newInts := reverseInts(oldInts)

	assert.Equal(t, oldInts[0],newInts[2])
	assert.Equal(t, oldInts[1],newInts[1])
	assert.Equal(t, oldInts[2],newInts[0])
}
*/
func TestStrCompare(t *testing.T) {
	assert.Equal(t, strings.Compare("0", "1"), -1)
	assert.Equal(t, strings.Compare("1", "1"), 0)
	assert.Equal(t, strings.Compare("11", "1"), 1)
	assert.Equal(t, strings.Compare("01", "00"), 1)

	assert.Equal(t, strings.Compare("/huawei/0/0", "/huawei/0/1"), -1)

}

func TestPaginate(t *testing.T) {
	initMock()
	kapi := dao.GetKeysAPI()
	mutiInsert(kapi, 3, 101)

	data := []struct {
		Root     string
		PageSize int
		LastDir  string
		LastKey  string
		Count    int
	}{
		{Root: "/huawei", LastDir: "", LastKey: "", PageSize: 13, Count: 13},
		{Root: "/huawei", LastDir: "", LastKey: "", PageSize: 1000, Count: 303},
		{Root: "/huawei", LastDir: "d000", LastKey: "", PageSize: 100, Count: 100},
		{Root: "/huawei", LastDir: "d000", LastKey: "/huawei/d000/k001", PageSize: 100, Count: 1},
		{Root: "/huawei", LastDir: "d001", LastKey: "/huawei/d001/k009", PageSize: 12, Count: 12},
	}

	for _, d := range data {
		opts := &dao.PageOptions{
			Root:     d.Root,
			LastDir:  d.LastDir,
			LastKey:  d.LastKey,
			PageSize: d.PageSize,
		}

		resp, err := dao.Paginate(context.Background(), kapi, opts)
		assert.Equal(t, nil, err, "should not be nil")
		assert.Equal(t, d.Count, len(resp.Node.Nodes))

		fmt.Println("resp.Node:", resp.Node)
		//for _, n := range resp.Node.Nodes {
		//	fmt.Println("node.key:", n.Key)
		//}
	}
}
