package tests

import (
	"fmt"
	"testing"
)

func TestGolang(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	fmt.Println(arr)

	arr[1] = 11
	fmt.Println(arr)

	change(arr)
	fmt.Println(arr)
}

func change(ints []int) {
	ints[1] = 22
	//fmt.Println(ints)
}

/*
func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")

	fmt.Println(r.Header.Get("Content-Type"))
	fmt.Println(ioutil.ReadAll(r.Body))

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	fmt.Fprintln(w, "upload ok!", handler.Filename)
}

func index(w http.ResponseWriter, r *http.Request) {
	println(getGID())
	var id = Goid()
	w.Write([]byte(strconv.Itoa(id) + ":" + tpl))
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func Goid() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recover:panic info:%v", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

const tpl = `<html>
<head>
<title>上传文件</title>
</head>
<body>
<form enctype="multipart/form-data" action="/upload" method="post">
 <input type="file" name="uploadfile" />
 <input type="hidden" name="token" value="haolipeng"/>
 <input type="submit" value="upload" />
</form>
</body>
</html>`
*/