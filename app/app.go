package app

import (
	"errors"
	"fmt"
	"github.com/cnlisea/seas/code"
	"github.com/cnlisea/seas/config"
	"github.com/cnlisea/seas/configc"
	"github.com/cnlisea/seas/discovery"
	"github.com/cnlisea/seas/file"
	"github.com/cnlisea/seas/service"
)

type App struct {
	Cfg *config.App
}

func New(cfg *config.App) *App {
	return &App{
		Cfg: cfg,
	}
}

func (a *App) Init() error {
	if a.Cfg == nil {
		return errors.New("cfg is nil")
	}

	serviceNum := len(a.Cfg.Services)
	if serviceNum == 0 {
		return errors.New("service config not settings")
	}

	var (
		codeBuffer = code.New()
		err        error
	)
	// code
	if err = codeBuffer.Init(); err != nil {
		return errors.New("code init fail: " + err.Error())
	}

	// config center
	if a.Cfg.Config != nil {
		cc := configc.New(a.Cfg.Config)
		if err = cc.Run(codeBuffer); err != nil {
			return errors.New("config center run fail: " + err.Error())
		}
	}

	// discovery
	if a.Cfg.Discovery != nil {
		d := discovery.New(a.Cfg.Discovery)
		if err = d.Run(codeBuffer); err != nil {
			return errors.New("discovery run fail: " + err.Error())
		}
	}

	var (
		s     *service.Service
		clone *code.Code
		f     *file.File
	)
	for i := 0; i < serviceNum; i++ {
		clone = codeBuffer.Clone()
		s = service.New(a.Cfg.Services[i])
		if err = s.Run(clone); err != nil {
			return errors.New("service run fail: " + err.Error())
		}
		// write file
		f = file.New("./main.go")
		if err = f.Init(); err != nil {
			return errors.New("file init fail: " + err.Error())
		}
		if _, err = f.WriteString(clone.Code()); err != nil {
			return errors.New("file write string fail: " + err.Error())
		}
		f.Close()
		fmt.Printf("num:%d, code:\n%s\n", i, clone.Code())
	}

	return nil
}

func (a *App) Run() error {

	return nil
}
