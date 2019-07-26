package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"os"
)

var (
	url     string
	port    int
	chat_id string
	h       bool
)

func init() {
	flag.IntVar(&port, "port", 8888, "listen port")
	flag.StringVar(&url, "url", "", "telegram sendMessage url")
	flag.StringVar(&chat_id, "char_id", "", "telegram chat id")
	flag.BoolVar(&h, "h", false, "print help")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `telegram prometheus bot: 
Usage: bot [-port 6666] [-url http://xxxxxx] -h help

Options:	
`)
	flag.PrintDefaults()
}

func httpHandle(ctx *fasthttp.RequestCtx) {

	var (
		body      string
		json_body Message
	)

	req := &ctx.Request

	body = string(req.Body())

	json.Unmarshal([]byte(body), &json_body)

	list := json_body.FormatBody(json_body)

	SendMessage(list, url, chat_id)

}

func main() {

	flag.Parse()

	if h {
		flag.Usage()
		return
	}

	if url == "" || chat_id == "" {
		flag.Usage()
		return
	}

	router := fasthttprouter.New()
	router.POST("/alert", httpHandle)
	listenHost := fmt.Sprintf("0.0.0.0:%d", port)

	fmt.Println(listenHost)
	if err := fasthttp.ListenAndServe(listenHost, router.Handler); err != nil {
		logs.Error("start fasthttp fail:", err.Error())
	}
}
