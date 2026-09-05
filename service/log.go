package service

import (
	"errors"
	"github.com/cnlisea/seas/code"
	"strconv"
)

func (s *Service) Log(code *code.Code) error {
	if s.Cfg == nil || s.Cfg.Log == nil {
		return nil
	}

	if code == nil {
		return errors.New("code is nil")
	}
	code.MainWriteString("\t//service log\n")
	code.MainWriteString("\tif err = a.Logger(\"", s.Cfg.Log.Path, "\", ",
		strconv.FormatUint(uint64(s.Cfg.Log.Level), 10), ", ",
		"true",
		"0",
		"); err != nil {\n",
		"\t\tpainc(err)\n",
		"\t}\n",
	)
	return nil
}
