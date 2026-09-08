package app

import (
	"container/list"
	"errors"
	"os/exec"

	"github.com/cnlisea/seas/code"
	"github.com/cnlisea/seas/config"
	"github.com/cnlisea/seas/configc"
	"github.com/cnlisea/seas/discovery"
	"github.com/cnlisea/seas/file"
	"github.com/cnlisea/seas/service"
)

type App struct {
	Cfg      *config.App
	Services *list.List
}

func New(cfg *config.App) *App {
	return &App{
		Cfg:      cfg,
		Services: list.New(),
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
	)
	for i := 0; i < serviceNum; i++ {
		clone = codeBuffer.Clone()
		s = service.New(a.Cfg.Services[i])
		if err = s.Run(clone); err != nil {
			return errors.New("service run fail: " + err.Error())
		}
		clone.Name = a.Cfg.Services[i].Name
		clone.Version = a.Cfg.Services[i].Version
		a.Services.PushBack(clone)
	}
	return nil
}

func (a *App) Run() error {
	if a.Services.Len() == 0 {
		return nil
	}
	var (
		e   *list.Element
		c   *code.Code
		f   *file.File
		cmd *exec.Cmd
		err error
	)
	for e = a.Services.Front(); e != nil; e = e.Next() {
		c = e.Value.(*code.Code)

		// gen main.go file
		f = file.New("cmd/main.go")
		if err = f.Init(); err != nil {
			return errors.New("file init fail: " + err.Error())
		}
		if _, err = f.WriteString(c.Code()); err != nil {
			return errors.New("file write string fail: " + err.Error())
		}
		f.Close()

		// copy upx
		cmd = exec.Command("cp",
			`$GOPATH/src/github.com/cnlisea/seas/build/tool/upx`,
			`$GOPATH/src/github.com/cnlisea/seas/build/docker/Dockerfile`,
			"cmd/")
		_, err = cmd.Output()
		if err != nil {
			return errors.New("exec cp upx and Dockerfile fail: " + err.Error())
		}

		// docker build
		cmd = exec.Command("docker", "build", "-t", c.Name+":"+c.Version, "-f", "cmd/Dockerfile")
		_, err = cmd.Output()
		if err != nil {
			return errors.New("exec docker fail: " + err.Error())
		}
	}
	return nil
}
