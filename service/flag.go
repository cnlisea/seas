package service

import (
	"errors"
	"github.com/cnlisea/seas/code"
	"strconv"
)

func (s *Service) Flag(code *code.Code) error {
	if s.Cfg == nil || len(s.Cfg.Flag) == 0 {
		return nil
	}

	if code == nil {
		return errors.New("code is nil")
	}

	code.MainWriteString("\t//service flag\n")
	code.ImportWriteString("\t\"gopkg.in/alecthomas/kingpin.v2\"\n")

	var (
		num    = len(s.Cfg.Flag)
		cfgMap = make(map[string]string, num)
		cmdVar string
		i      int
	)
	for i = 0; i < num; i++ {
		cmdVar = "cmdVar" + strconv.Itoa(i)
		code.MainWriteString("\tvar ", cmdVar, " = kingpin.Flag(\"", s.Cfg.Flag[i].Name, "\", \"", s.Cfg.Flag[i].Help, "\")")
		if s.Cfg.Flag[i].Short != 0 {
			code.MainWriteString(".Short('", string(s.Cfg.Flag[i].Short), "')")
		}
		if s.Cfg.Flag[i].Require {
			code.MainWriteString(".Required()")
		}
		if s.Cfg.Flag[i].Default != nil {
			code.MainWriteString(".Default(\"", *s.Cfg.Flag[i].Default, "\")")
		}
		if s.Cfg.Flag[i].Env != "" {
			code.MainWriteString(".Envar(\"", s.Cfg.Flag[i].Env, "\")")
		}
		code.MainWriteString(".String()")
		cfgMap[s.Cfg.Flag[i].Key] = cmdVar
	}
	if s.Cfg.Version != "" {
		code.MainWriteString("\tkingpin.Version(\"", s.Cfg.Version, "\")\n",
			"\tkingpin.HelpFlag.Short('I')\n")
	}
	code.MainWriteString("kingpin.Parse()")

	const cmdValCfgKey = "cmdVarCfg "
	code.MainWriteString("\tvar ", cmdValCfgKey, " = map[string]string{\n")
	for k, v := range cfgMap {
		code.MainWriteString("\t\t\"", k, "\": \"", v, "\",\n")
	}
	code.MainWriteString("\t}")
	code.MainWriteString("\ta.ProxyConfig().SetCfg(\"flag\", ", cmdValCfgKey, ")\n")
	return nil
}
