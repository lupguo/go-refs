package dotenv

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestDotEnv(t *testing.T) {
	getwd, _ := os.Getwd()
	envfile := getwd + "/../../.env"
	if err := godotenv.Load(envfile); err != nil {
		log.Fatal("Error load dot env")
	}

	s3Bucket := os.Getenv("S3_BUCKET")
	s3SecrectKey := os.Getenv("S3_SECRET_KEY")

	t.Logf("s3Bucket=>%s, s3SecerctKey=>%s\n", s3Bucket, s3SecrectKey)
}

func TestDotEnvLoad(t *testing.T) {
	// 支持YML读取
	getwd, _ := os.Getwd()
	dotfile := getwd + "/.env.yml"
	if err := godotenv.Load(dotfile); err != nil {
		log.Fatal("Error load dot env")
	}
	foo, bar, baz := os.Getenv("FOO"), os.Getenv("BAR"), os.Getenv("BAZ")
	t.Logf("foo=>%s, bar=>%s, baz=>%s\n", foo, bar, baz)

	// 区分大小写
	t.Logf("http_proxy=>%s, https_proxy_1=>%s, https_proxy_2=>%s \n", os.Getenv("http_proxy"), os.Getenv("HTTPS_PROXY"), os.Getenv("https_proxy"))
}

func TestDotEnvLoadRemote(t *testing.T) {
	// go test http
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		jdt, _ := json.Marshal(map[string]string{
			"BUCKET_ID":  "id3302",
			"BUCKET_KEY": "secret110myEnvmyEnv",
		})
		if _, err := w.Write(jdt); err != nil {
			return
		}
	}
	// new an test server
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	// request test server
	res, err := ts.Client().Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	myEnv, err := godotenv.Parse(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", myEnv)
}

func TestRespRecorder(t *testing.T) {
	// define handler
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	}
	// define request
	req := httptest.NewRequest("GET", "http://httpbin.org/ip", nil)
	t.Logf("%+v", req)
	// new an recorder impl http.ResponseWriter
	w := httptest.NewRecorder()
	// handle
	handler(w, req)
	// get fake http response
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Logf("%+v", resp)
	t.Log(string(body))
}

func TestUnstartedHTTP2(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s", r.Proto)
	}
	// 创建一个server但并不启动他
	ts := httptest.NewUnstartedServer(http.HandlerFunc(handler))
	// 设置
	ts.EnableHTTP2 = true
	// 启动
	ts.StartTLS()
	t.Logf("http server addr: %s", ts.URL)
	defer ts.Close()

	// 请求测试地址
	res, err := ts.Client().Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
	t.Logf("%s", greeting)
}

func TestNormalHTTPServer(t *testing.T) {
	// 创建一个普通HTTP1.1测试服务
	handler := func(w http.ResponseWriter, r *http.Request) {
		// json header
		w.Header().Set("Content-Type", "application/json")
		// json回包
		m := map[string]interface{}{
			"id":    100,
			"name":  "clark",
			"likes": []string{"play game", "watch tv", "read book"},
		}
		jdt, err := json.Marshal(m)
		if err != nil {
			t.Fatal(err)
		}
		w.Write(jdt)
		// fmt.Fprintf(w, "%s", jdt)
	}
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	// 请求
	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("%s", res.Header)
	t.Logf("%s", res.Request.URL)
	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("%s", greeting)
}

func TestTLSHTTP(t *testing.T) {
	// 直接创建一个TLS测试Server
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	// 请求测试地址
	res, err := ts.Client().Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	t.Logf("%s", greeting)
}
