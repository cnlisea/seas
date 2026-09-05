package main

import (
	"fmt"

	"github.com/alecthomas/kingpin/v2"
	antApp "github.com/cnlisea/ant/app"
	"github.com/cnlisea/seas/app"
	"github.com/cnlisea/seas/config"
)

func main() {
	var (
		path = kingpin.Flag(
			"path",
			"path on which to config path.",
		).Short('p').Default("./").String()
		name = kingpin.Flag(
			"name",
			"name on which to config name",
		).Short('n').Required().String()
	)
	kingpin.Version("1.0")
	kingpin.HelpFlag.Short('I')
	kingpin.Parse()

	fmt.Println("path:", *path, " name:", *name)
	var (
		cfg = new(config.App)
		err error
	)
	err = antApp.New().ConfigRegister("", *name, *path, true, cfg, nil)
	if err != nil {
		panic(*path + " config get fail: " + err.Error())
	}

	a := app.New(cfg)
	if err = a.Init(); err != nil {
		panic("app init fail:" + err.Error())
	}

	if err = a.Run(); err != nil {
		panic("app run fail:" + err.Error())
	}

	fmt.Println("app run successfully!")
}
