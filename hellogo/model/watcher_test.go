package model

import (
	"testing"
	"time"
	"math/rand"
	"strconv"
	"hellogo/logger"
	"github.com/astaxie/beego"
	"sync"
)

func init() {
	logger.SetLevel(beego.LevelInformational)
	//watchEngine.capacity = 50
	//watchEngine.expire = 10
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func TestLogger(t *testing.T) {
	logger.Debug("debug")
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
	watchEngine.Init(50, 20)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 51; i++ {
		w, err := watchEngine.GetWatcher(strconv.Itoa(i) + ":" + strconv.Itoa(r.Int()))
		logger.Info("GetWatcher", i, w, err)
	}
	watchEngine.clearWatchers()
}

var waitgroup sync.WaitGroup

func TestWatcher(t *testing.T) {
	watchEngine.Init(50, 20)
	watchEngine.Start()

	for i := 0; i < 10; i++ {
		waitgroup.Add(1)
		go getWatchers(i)
	}

	waitgroup.Wait()
	time.Sleep(watch_interval * 3)
	for num := 1; num > 0; num = len(watchEngine.processor) {
		time.Sleep(time.Second)
	}
	watchEngine.Stop()
}

func getWatchers(index int) {
	defer waitgroup.Done()

	for i := 0; i < 100; i++ {
		_, err := watchEngine.GetWatcher(strconv.Itoa(index) + ":" + strconv.Itoa(r.Int()))
		if err != nil {
			logger.Info("get watcher err:", err, "index:", index, "count:", i)
			i--
			time.Sleep(watch_interval)
		}
	}
}