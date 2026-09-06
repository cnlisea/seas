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
		code.ImportWriteString("\t", cfgPackageName, " \"", s.Cfg.Config[i].Path, "\"\n")
		cfgName = "cfg" + strconv.Itoa(i)
		code.MainWriteString("\tvar ", cfgName, " ", cfgPackageName, ".", s.Cfg.Config[i].Name)
		code.MainWriteString("\tif err = a.ConfigRegister(",
			"\"", s.Cfg.Config[i].Key, "\", ",
		)
		switch s.Cfg.Config[i].Channel {
		case 1:
			code.MainWriteString("\"\", false, &", cfgName, ", &app.RegisterCenter{\n\t\tGroupId: \"", s.Cfg.Config[i].Nacos.GroupId, "\",\n\t\tDataId: \"", s.Cfg.Config[i].Nacos.DataId, "\",\n")
			//update hook
			if len(s.Cfg.Config[i].Nacos.UpdateHook) > 0 {
				code.MainWriteString("\t\tUpdateHook: func(obj interface{}) {\n")
				for j := range s.Cfg.Config[i].Nacos.UpdateHook {
					if s.Cfg.Config[i].Nacos.UpdateHook[j] == nil {
						continue
					}
					cfgPackageName = "cfgUpHookPackage" + strconv.Itoa(j)
					code.ImportWriteString("\t", cfgPackageName, " \"", s.Cfg.Config[i].Nacos.UpdateHook[j].Path, "\"\n")
					code.MainWriteString("\t\t\t", cfgPackageName, ".", s.Cfg.Config[i].Nacos.UpdateHook[j].Name, "()\n")
				}
				code.MainWriteString("\t\t}\n")
			}
			code.MainWriteString("}")
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
