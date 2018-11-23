package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"strconv"
)

var logger = logrus.New()
var config *jsonConfig
var wechatWork *WechatWork
var req = Request{}
var version = "0.1"

func main() {
	var router = httprouter.New()
	router.GET("/", index)
	router.POST("/push",push)
	addr := config.Host + ":" + strconv.Itoa((config.Port))
	logger.Infof("listening at %s",addr)
	log.Fatal(http.ListenAndServe(addr, router))
	//wechatWork.SendMessage("213")

}

func init() {
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	params := parseCmdParams()
	if params.isPrintVersion {
		fmt.Printf("version: %s", version)
		os.Exit(0)
	}
	var err error
	config, err = loadConfig(params.configPath)
	wechatWork = &WechatWork{config, ""}
	if err != nil {
		print(err)
	}
}