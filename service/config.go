package service

import (
	"errors"
	"github.com/cnlisea/seas/code"
	"strconv"
)

func (s *Service) Config(code *code.Code) error {
	if s.Cfg == nil || s.Cfg.Config == nil {
		return nil
	}

	if code == nil {
		return errors.New("code is nil")
	}

	code.MainWriteString("\t//service config\n")
	var (
		cfgPackageName string
		cfgName        string
		i              int
	)
	for i = range s.Cfg.Config {
		if s.Cfg.Config[i] == nil {
			continue
		}
		cfgPackageName = "cfgPackage" + strconv.Itoa(i)
		code.ImportWriteString("\t", cfgPackageName, " \"", s.Cfg.Config[i].Path, "\"")
		code.MainWriteString("\tvar ", cfgName, cfgPackageName, ".", s.Cfg.Config[i].Name)
		cfgName = "cfg" + strconv.Itoa(i)
		code.MainWriteString("\tif err = a.ConfigRegister(",
			"\"", s.Cfg.Config[i].Key, "\", ",
		)
		switch s.Cfg.Config[i].Channel {
		case 1:
			code.MainWriteString("\"\", false, &", cfgName, ", &app.RegisterCenter{GroupId: \"", s.Cfg.Config[i].Nacos.GroupId, "\", DataId:  \"", s.Cfg.Config[i].Nacos.DataId, "\"}")

		default:
			code.MainWriteString("\"", s.Cfg.Config[i].Local.Name, "\", true, &", cfgName, ", nil")
		}
		code.MainWriteString("); err != nil {\n",
			"\t\tpanic(err)\n",
			"\t}\n",
		)

	}
	return nil
}
