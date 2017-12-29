package main

import (
	"flag"
	"fmt"

	"github.com/dynamicgo/config"
	"github.com/dynamicgo/slf4go"
	"github.com/go-xorm/xorm"
	"github.com/inwecrypto/mailbox"
	_ "github.com/lib/pq"
)

var logger = slf4go.Get("ethdb")
var configpath = flag.String("conf", "../conf/mailbox.json", "mailbox database config file path")

func main() {

	flag.Parse()

	conf, err := config.NewFromFile(*configpath)

	if err != nil {
		logger.ErrorF("load eth indexer config err , %s", err)
		return
	}

	username := conf.GetString("mailbox.db.username", "xxx")
	password := conf.GetString("mailbox.db.password", "xxx")
	port := conf.GetString("mailbox.db.port", "6543")
	host := conf.GetString("mailbox.db.host", "localhost")
	scheme := conf.GetString("mailbox.db.schema", "postgres")

	engine, err := xorm.NewEngine("postgres", fmt.Sprintf("user=%v password=%v host=%v dbname=%v port=%v sslmode=disable", username, password, host, scheme, port))

	if err != nil {
		logger.ErrorF("create postgres orm engine err , %s", err)
		return
	}

	if err := engine.Sync2(new(mailbox.Status), new(mailbox.Mail), new(mailbox.User)); err != nil {
		logger.ErrorF("sync table schema error , %s", err)
		return
	}

}
