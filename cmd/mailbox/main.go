package main

import (
	"flag"

	"github.com/dynamicgo/config"
	"github.com/dynamicgo/slf4go"
	"github.com/inwecrypto/mailbox"
	_ "github.com/lib/pq"
)

var logger = slf4go.Get("mailbox")
var configpath = flag.String("conf", "../conf/mailbox.json", "mailbox config path")

func main() {

	flag.Parse()

	neocnf, err := config.NewFromFile(*configpath)

	if err != nil {
		logger.ErrorF("load mailbox config err , %s", err)
		return
	}

	// factory, err := aliyunlog.NewAliyunBackend(neocnf)

	// if err != nil {
	// 	logger.ErrorF("create aliyun log backend err , %s", err)
	// 	return
	// }

	// slf4go.Backend(factory)

	monitor, err := mailbox.NewAPIServer(neocnf)

	if err != nil {
		logger.ErrorF("create neo monitor err , %s", err)
		return
	}

	monitor.Run()

}
