package dao_test

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/coreos/etcd/client"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"github.com/readhelper/beedemo/dao"
	"testing"
	"time"
)

func initMock() {
	dao.SetKeysAPI(dao.GetMockEtcd())
	kapi := dao.GetKeysAPI()
	_, err := kapi.Delete(context.Background(), "/huawei/", &client.DeleteOptions{Dir: true, Recursive: true})
	if err != nil {
		beego.Warn("init etcd delete old data error", err)
	}
}

func mutiInsert(kapi client.KeysAPI, dirCnt int, keyCnt int) {
	for i := 0; i < dirCnt; i++ {
		for j := 0; j < keyCnt; j++ {
			kapi.Set(context.Background(), "/huawei/d"+fmt.Sprintf("%03d", i)+"/k"+fmt.Sprintf("%03d", j), time.Now().String(), &client.SetOptions{})
		}
	}
}
func TestMultiKey(t *testing.T) {
	initMock()
	kapi := dao.GetKeysAPI()

	data := []struct {
		dirCnt int
		keyCnt int
	}{
		{dirCnt: 3, keyCnt: 10},
	}

	for _, d := range data {
		mutiInsert(kapi, d.dirCnt, d.keyCnt)

		resp, err := kapi.Get(context.Background(), "/huawei", &client.GetOptions{Sort: true})
		assert.NoError(t, err, "should be nil")
		assert.Equal(t, d.dirCnt, len(resp.Node.Nodes))

		resp, err = kapi.Get(context.Background(), "/huawei/d000", &client.GetOptions{Sort: true})
		assert.NoError(t, err, "should be nil")
		assert.Equal(t, d.keyCnt, len(resp.Node.Nodes))
	}
}
