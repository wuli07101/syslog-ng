package main

import (
	// "fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/bitly/go-simplejson"
	"gopkg.in/mcuadros/go-syslog.v2"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	server := syslog.NewServer()
	server.SetFormat(syslog.Automatic)
	server.SetHandler(handler)
	server.ListenTCP("0.0.0.0:5121")

	server.Boot()

	go dealMessage(channel)

	server.Wait()
}

func dealMessage(channel syslog.LogPartsChannel) {
	for logParts := range channel {
		var content interface{} = logParts["content"]
		b := content.(string)
		c := []byte(b)

		syslogContent, err := simplejson.NewJson(c)
		if err == nil {
			httpUrl := syslogContent.Get("url").MustString()
			httpMsg := syslogContent.Get("data").MustString()
			sendToHttp(httpUrl, httpMsg)
		}
	}
}

func sendToHttp(httpUrl, httpMsg string) {
	req := httplib.Post(httpUrl).SetTimeout(2*time.Second, 2*time.Second)
	req.Header("Content-Type", "application/x-www-form-urlencoded")
	req.Body(httpMsg)
	req.Bytes()
}
