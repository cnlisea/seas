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
		short  []rune
		cmdVar string
		i      int
	)
	for i = 0; i < num; i++ {
		cmdVar = "cmdVar" + strconv.Itoa(i)
		code.MainWriteString("\tvar ", cmdVar, " = kingpin.Flag(\"", s.Cfg.Flag[i].Name, "\", \"", s.Cfg.Flag[i].Help, "\")")
		if s.Cfg.Flag[i].Short != "" {
			short = []rune(s.Cfg.Flag[i].Short)
			if len(short) > 0 {
				code.MainWriteString(".Short('", string(short[0]), "')")
			}
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
		code.MainWriteString(".String()\n")
		cfgMap[s.Cfg.Flag[i].Key] = cmdVar
	}
	if s.Cfg.Version != "" {
		code.MainWriteString("\tkingpin.Version(\"", s.Cfg.Version, "\")\n",
			"\tkingpin.HelpFlag.Short('I')\n")
	}
	code.MainWriteString("\tkingpin.Parse()\n")

	code.MainWriteString("\ta.ProxyConfig().SetCfg(\"flag\", ", "map[string]string{\n")
	for k, v := range cfgMap {
		code.MainWriteString("\t\t\"", k, "\": *", v, ",\n")
	}
	code.MainWriteString("\t})\n")
	return nil
}
