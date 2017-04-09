# 当GOLang遇到大量ESTABLISHED的分析及解决方案
## 现象描述
当Go程序(两个GO程序互为服务端、客户端)运行一段时间后，使用netstat -a查看，就出现大量ESTABLISHED的连接，且久久保持不断开。
目前这个问题涉及PSM,PM，郭文丹负责的项目等go应用。对大量请求第三方服务的go应用影响较大，PSM出现过一次端口被用光的情况。
![cmd-markdown-logo](img/net1.jpg)
客户端连接代码如下：
```
client := &http.Client{}
	client.Transport = tr
	req, err := http.NewRequest(metod, url, strings.NewReader(""))
	if err != nil {
		fmt.Println("http.NewRequest error", err)
		return
	}

	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println("client.Do error", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll error", err)
	}
```
## 问题定位
一般来说就是在进行HTTP(TCP)调用时没有断开或者说关闭连接造成的。<br>
首先排查服务端,分别访问baidu,beego官网，同样会出现ESTABLISHED问题。
```
func TestRequetToBeego(t *testing.T) {
	var N = 100000
	for i := 0; i < N; i++ {
		transport := &http.Transport{}
		httpDo("GET", "https://beego.me/", transport)
		time.Sleep(time.Millisecond * 10)
	}
	time.Sleep(time.Second * 60)
}
func TestRequetToBaidu(t *testing.T) {
	var N = 100000
	for i := 0; i < N; i++ {
		transport := &http.Transport{}
		httpDo("GET", "http://www.baidu.com/", transport)
		time.Sleep(time.Millisecond * 10)
	}
	time.Sleep(time.Second * 60)
}
```

重点排查客户端的go程序，查阅client源代码,发现线索，如果Transport没有设置，那么会使用DefaultTransport。
```
type Client struct {
	// Transport specifies the mechanism by which individual
	// HTTP requests are made.
	// If nil, DefaultTransport is used.
	Transport RoundTripper
	... 
}	
```
将测试用例改为使用DefaultTransport，发现不会出现ESTABLISHED问题，进一步分析DefaultTransport。
```
var DefaultTransport RoundTripper = &Transport{
	Proxy: ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}).DialContext,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}
```
再次分析，定位为IdleConnTimeout
```
// IdleConnTimeout is the maximum amount of time an idle
	// (keep-alive) connection will remain idle before closing
	// itself.
	// Zero means no limit.
	IdleConnTimeout time.Duration
```
参考该字段含义，设置自定义的Transport.IdleConnTimeou=2 * time.Second,
ESTABLISHED问题解决。

## 解决方案
```
//方案一：设置IdleConnTimeout，超时断开
func TestRequetOk1(t *testing.T) {
	var N = 10
	for i := 0; i < N; i++ {
		transport := &http.Transport{
			IdleConnTimeout:       2 * time.Second,
		}
		httpDo("GET", "http://192.168.0.4:12345", transport)
		time.Sleep(time.Millisecond * 100)
	}
	time.Sleep(time.Second * 60)
}
//方案二：设置DisableKeepAlives，不重用
func TestRequetOk2(t *testing.T) {
	var N = 10
	for i := 0; i < N; i++ {
		transport := &http.Transport{
			DisableKeepAlives:true,
		}
		httpDo("GET", "http://192.168.0.4:12345", transport)
		time.Sleep(time.Millisecond * 100)
	}
	time.Sleep(time.Second * 60)
}
//方案三：共享全局Transport
func TestRequetOk3(t *testing.T) {
	var N = 10
	transport := &http.Transport{}
	for i := 0; i < N; i++ {
		httpDo("GET", "http://192.168.0.4:12345", transport)
		time.Sleep(time.Second * 3)
	}
	time.Sleep(time.Second * 60)
}
//方案四：设置req.Header.Add("connection", "close")
func TestRequetOk4(t *testing.T) {
	var N = 100
	for i := 0; i < N; i++ {
		transport := &http.Transport{}
		//req.Header.Add("connection", "close")
		httpDoWithClose("GET", "http://www.baidu.com", transport)
		time.Sleep(time.Second * 3)
	}
	time.Sleep(time.Second * 60)
}          
```
## 扩展阅读       
* IdleConnTimeout 是golang1.7引入的，1.6标记为todo，golang团队应该意识到这个设计缺陷。
```
// TODO: tunable on global max cached connections
// TODO: tunable on timeout on cached connections
// TODO: tunable on max per-host TCP dials in flight (Issue 13957)
```
* 目前我们go应用基本上用的是beego框架，服务端长时间不关闭连接，也可能造成自身崩溃
```
GO 可以通过 net.TCPConn 的 SetKeepAlive 来启用 TCP keepalive。在 OS X 和 Linux 系统上，当一个连接空间了2个小时时，会以75秒的间隔发送8个TCP keepalive探测包。换句话说， 在两小时10分钟后(7200+8*75)Read将会返回一个 io.EOF 错误.

对于你的应用，这个超时间隔可能太长了。在这种情况下你可以调用SetKeepAlivePeriod方法。但这个方法在不同的操作系统上会有不同的表现。在OSX上它会更改发送探测包前连接的空间时间。在Linux上它会更改连接的空间时间与探测包的发送间隔。所以以30秒的参数调用 SetKeepAlivePeriod在OSX系统上会导致共10分30秒(30+8*75)的超时时间，但在linux上却是4分30秒(30+8*30).
```
