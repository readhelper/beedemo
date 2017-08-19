package model

import (
	"time"
	"sync"
	"hellogo/proxys"
	"hellogo/logger"
	"strings"
	"errors"
	"sync/atomic"
	"strconv"
)

const (
	watch_interval = 2 * time.Second
)

type Watcher interface {
	Next(uid string) (string, error)
}

type watcher struct {
	key     string
	uid     string
	status  int
	content string
	peers   int32
	killed  bool
	start   time.Time
	end     time.Time
}

func (this *watcher) addPeer() {
	atomic.AddInt32(&this.peers, 1)
}
func (this *watcher) delPeer() {
	atomic.AddInt32(&this.peers, -1)
}

func (this *watcher) Next(uid string) (string, error) {
	this.addPeer()
	defer this.delPeer()

	for i := 0; i < watchEngine.expire; i = i + 2 {
		if len(this.content) > 0 && (strings.Compare(this.uid, uid) != 0 || this.killed) {
			return this.content, nil
		}
		if len(this.content) == 0 && this.killed {
			return "", errors.New("watch is killed by " + this.end.String())
		}
		time.Sleep(watch_interval)
	}

	if len(this.content) > 0 {
		return this.content, nil
	}
	return "", errors.New("watch timeout")
}

type WatchEngine struct {
	sync.RWMutex
	capacity    int
	watchers    map[string]*watcher
	processor   chan *watcher
	checkTicker *time.Ticker
	clearTicker *time.Ticker
	count       int32
	expire      int
}

func (this *WatchEngine) Init(capacity int, expire int) {
	watchEngine.capacity = capacity
	watchEngine.expire = expire
	watchEngine.processor = make(chan *watcher, capacity)
}

func (this *WatchEngine) Start() {
	this.startWatching()
	go this.check()  //lock
	go this.process() //unlock
}

func (this *WatchEngine) Stop() {
	this.checkTicker.Stop()
	this.clearTicker.Stop()
	close(this.processor)
	this.stopWatching()
}

func (this *WatchEngine) check() {
	for {
		select {
		case <-this.checkTicker.C:
			{
				this.checkWatchers()
			}
		case <-this.clearTicker.C:
			{
				this.clearWatchers()
			}
		}
	}
}
func (this *WatchEngine) process() {
	for {
		select {
		case w := <-this.processor:
			{
				this.counter()
				this.processWatcher(w)
			}
		}
	}
}
func (this *WatchEngine) counter() {
	num := atomic.AddInt32(&this.count, 1)
	if num > 100000000 {
		atomic.SwapInt32(&this.count, 0)
	}
}

func (this *WatchEngine) processWatcher(w *watcher) {
	if w == nil {
		logger.Warn("process watcher is nil,processor:", this.processor)
		return
	}
	sc, ret, err := proxys.GetRemoteObject(w.key)
	if err != nil {
		logger.Warn("getRemoteObject err:", err, "key:", w.key)
	}
	w.uid = strconv.Itoa(sc)
	w.status = sc
	w.content = ret
}
func (this *WatchEngine) startWatching() {
	logger.Debug("start watching ...watchers num:", len(this.watchers), "capacity:", this.capacity, "queue:", len(this.processor), "count:", this.count)
}

func (this *WatchEngine) stopWatching() {
	logger.Debug("stop watching ...watchers num:", len(this.watchers), "capacity:", this.capacity, "queue:", len(this.processor), "count:", this.count)
}

func (this *WatchEngine) checkWatchers() {
	if len(this.watchers) == 0 {
		time.Sleep(watch_interval)
		logger.Debug("watchers is empty.so sleep ...")
	}

	this.Lock()
	defer this.Unlock()

	logger.Info("watchers num:", len(this.watchers), "capacity:", this.capacity, "queue:", len(this.processor), "count:", this.count)
	var isFull bool = false
	for _, val := range this.watchers {
		queue := len(this.processor)
		if queue < this.capacity - 1 {
			this.processor <- val
		}else if !isFull {
			isFull = true
			logger.Info("process queue is full.capacity:", this.capacity, "queue:", queue)
		}
	}
}

func (this *WatchEngine) GetKeys() []string {
	var keys = []string{}
	for key, _ := range this.watchers {
		keys = append(keys, key)
	}
	return keys
}

func (this *WatchEngine) GetWatcher(key string) (Watcher, error) {
	size := len(this.watchers)
	if size > this.capacity - 1 {
		return nil, errors.New("watcher is full. capacity:" + strconv.Itoa(this.capacity))
	}

	this.Lock()
	defer this.Unlock()

	w, ok := this.watchers[key]
	if !ok {
		w = &watcher{key:key}
		this.addWatcher(key, &watcher{key:key})
	}
	return w, nil
}
func (this *WatchEngine) addWatcher(key string, w *watcher) {
	w.start = time.Now()
	this.watchers[key] = w
	logger.Debug("add watcher,key=", key, "time:", w.start.String())
}
func (this *WatchEngine) delWatcher(w *watcher) {
	w.killed = true
	w.end = time.Now()
	delete(this.watchers, w.key)
	logger.Debug("delete watcher.key:", w.key, "time:", w.start.String())
}

//proctecd watchengine
func (this *WatchEngine) clearWatchers() (error) {
	var before = len(this.watchers)
	if before < this.capacity / 2 {
		return nil
	}

	this.Lock()
	defer this.Unlock()

	num := len(this.processor)
	logger.Warn("begin clearing wachers. capacity:", this.capacity, "queue:", num, "count:", this.count, "before:", before)
	for _, val := range this.watchers {
		p := atomic.LoadInt32(&val.peers)
		if p <= 0 || time.Now().After(val.start.Add(time.Hour)) {
			this.delWatcher(val)
		}

	}

	after := len(this.watchers)
	logger.Warn("end clearing wachers. capacity:", this.capacity, "queue:", num, "count:", this.count, "after:", after)
	return nil
}

var watchEngine = WatchEngine{
	capacity:1000,
	watchers:map[string]*watcher{},
	checkTicker: time.NewTicker(watch_interval),
	clearTicker: time.NewTicker(5 * watch_interval),
	processor: make(chan *watcher, 1000),
	expire:60 * 60,
}