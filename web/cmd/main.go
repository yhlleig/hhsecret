package main

import (
	"flag"
	"fmt"

	// "github.com/Sirupsen/logrus"
	"github.com/bingbaba/hhsecret/web"
	"github.com/kataras/iris"
)

var (
	cfgfile = flag.String("c", "conf.yml", "the configure file")
	Debug   = flag.Bool("debug", false, "debug the log")
)

func main() {
	flag.Parse()

	// load configure file
	err := web.LoadConfigureByFile(*cfgfile)
	if err != nil {
		panic(err)
	}

	// init logger
	mconf := web.DefaultCfg
	web.InitLogger(mconf.LogFile, *Debug)

	// iris
	app := web.GetIrisApp()

	addr := iris.Addr(fmt.Sprintf(":%d", mconf.Port))
	if err := app.Run(addr); err != nil {
		panic(err)
	}
}
