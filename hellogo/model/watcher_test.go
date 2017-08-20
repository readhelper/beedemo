package model_test

import (
	"testing"
	"time"
	"math/rand"
	"strconv"
	"hellogo/logger"
	"github.com/astaxie/beego"
	"sync"
	. "hellogo/model"
)

func init() {
	logger.SetLevel(beego.LevelInformational)
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var watch_interval = 2 * time.Second
var waitgroup sync.WaitGroup


func TestWatch(t *testing.T) {
	WEngine.Init(50, 10)
	WEngine.Start()
	defer WEngine.Stop()

	time.Sleep(watch_interval * 2)
	w, err := WEngine.GetWatcher("watcher_test")
	if err != nil {
		t.Error("get watcher failed.err:", err)
	}
	str, err := w.Next("id")
	logger.Warn("w.next. conent:", str, "err:", err)
}

func TestWEngine(t *testing.T) {
	WEngine.Init(50, 10)
	WEngine.Start()
	defer WEngine.Stop()

	time.Sleep(watch_interval * 2)
	for i := 0; i < 10; i++ {
		waitgroup.Add(1)
		go getWatchers(i)
	}

	waitgroup.Wait()
	time.Sleep(watch_interval * 2)
	for num := 1; num > 0; num = WEngine.Count() {
		time.Sleep(time.Second)
	}
}

func getWatchers(index int) {
	defer waitgroup.Done()

	for i := 0; i < 10; i++ {
		_, err := WEngine.GetWatcher(strconv.Itoa(index) + ":" + strconv.Itoa(r.Int()))
		if err != nil {
			logger.Info("get watcher err:", err, "index:", index, "count:", i)
			i--
			time.Sleep(watch_interval)
		}
	}
}