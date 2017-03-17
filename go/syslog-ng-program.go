package main

import (
	"bufio"
	"github.com/astaxie/beego/httplib"
	"github.com/bitly/go-simplejson"
	"os"
	"runtime"
	"sync"
	"time"
)

var (
	sysLogMsgList chan []byte
	wg            sync.WaitGroup
)

func init() {
	sysLogMsgList = make(chan []byte, 1000000)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg.Add(2)
	go readMsgFromStdin()
	go dealMsg()
	wg.Wait()
}

func readMsgFromStdin() {
	reader := bufio.NewReader(os.Stdin)
	for {
		data, _, _ := reader.ReadLine()
		sysLogMsgList <- data
	}
	wg.Done()
}

func dealMsg() {

	for {
		select {
		case syslogMsg := <-sysLogMsgList:
			syslogContent, err := simplejson.NewJson(syslogMsg)
			if err == nil {
				jsonUrl, isJsonUrlOk := syslogContent.CheckGet("url")
				if isJsonUrlOk == false {
					break
				}
				httpUrl := jsonUrl.MustString()

				jsonMsg, isJsonMsgOk := syslogContent.CheckGet("data")
				if isJsonMsgOk == false {
					break
				}
				httpMsg := jsonMsg.MustString()

				go sendToHttp(httpUrl, httpMsg)
			}
		}
	}
	wg.Done()
}

func sendToHttp(httpUrl, httpMsg string) {
	req := httplib.Post(httpUrl).SetTimeout(2*time.Second, 2*time.Second)
	req.Header("Content-Type", "application/x-www-form-urlencoded")
	req.Body(httpMsg)
	req.Bytes()
}
