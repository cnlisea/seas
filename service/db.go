package service

import (
	"errors"
	"github.com/cnlisea/seas/code"
	"strconv"
)

func (s *Service) DB(code *code.Code) error {
	if s.Cfg == nil || s.Cfg.DB == nil {
		return nil
	}

	if code == nil {
		return errors.New("code is nil")
	}

	code.MainWriteString("\t//service db\n")
	// mysql
	if s.Cfg.DB.MySQL != nil {
		code.MainWriteString("\t//service db mysql\n")
		for i := range s.Cfg.DB.MySQL {
			if s.Cfg.DB.MySQL[i] == nil {
				continue
			}
			code.MainWriteString("\tif err = a.DBMySqlRegister(\"", s.Cfg.DB.MySQL[i].Name, "\", ",
				"\"", s.Cfg.DB.MySQL[i].User, "\", ",
				"\"", s.Cfg.DB.MySQL[i].Password, "\", ",
				"\"", s.Cfg.DB.MySQL[i].Addr, "\", ",
				strconv.FormatUint(uint64(s.Cfg.DB.MySQL[i].Port), 10), ", ",
				"\"", s.Cfg.DB.MySQL[i].DBName, "\", ",
				"\"", s.Cfg.DB.MySQL[i].Charset, "\", ",
				"\"", s.Cfg.DB.MySQL[i].Timeout, "\", ",
				strconv.FormatBool(s.Cfg.DB.MySQL[i].ParseTime), ", ",
				"\"", s.Cfg.DB.MySQL[i].Loc, "\", ",
				strconv.Itoa(s.Cfg.DB.MySQL[i].Active), ", ",
				strconv.Itoa(s.Cfg.DB.MySQL[i].Idle), ", ",
				strconv.Itoa(s.Cfg.DB.MySQL[i].IdleTimeout),
				"); err != nil {\n",
				"\t\tpainc(err)\n",
				"\t}\n",
			)
		}
	}

	// redis
	if s.Cfg.DB.Redis != nil {
		code.MainWriteString("\t//service db redis\n")
		for i := range s.Cfg.DB.Redis {
			if s.Cfg.DB.Redis[i] == nil {
				continue
			}
			code.MainWriteString("\tif err = a.DBRedisRegister(\"", s.Cfg.DB.Redis[i].Name, "\", ",
				"\"", s.Cfg.DB.Redis[i].Password, "\", ",
				"\"", s.Cfg.DB.Redis[i].Addr, "\", ",
				strconv.FormatUint(uint64(s.Cfg.DB.Redis[i].Port), 10), ", ",
				strconv.Itoa(s.Cfg.DB.Redis[i].DB), ", ",
				strconv.Itoa(s.Cfg.DB.Redis[i].Active), ", ",
				strconv.Itoa(s.Cfg.DB.Redis[i].Idle), ", ",
				strconv.Itoa(s.Cfg.DB.Redis[i].IdleTimeout),
				"); err != nil {\n",
				"\t\tpainc(err)\n",
				"\t}\n",
			)
		}
	}

	// mongo
	if s.Cfg.DB.MongoDB != nil {
		code.MainWriteString("\t//service db mongo\n")
		for i := range s.Cfg.DB.MongoDB {
			if s.Cfg.DB.MongoDB[i] == nil {
				continue
			}
			code.MainWriteString("\tif err = a.DBMongoRegister(\"", s.Cfg.DB.MongoDB[i].Name, "\", ",
				"\"", s.Cfg.DB.MongoDB[i].User, "\", ",
				"\"", s.Cfg.DB.MongoDB[i].Password, "\", ",
				"[]string{",
			)
			for j := range s.Cfg.DB.MongoDB[i].Addr {
				if j > 0 {
					code.MainWriteString(", ")
				}
				code.MainWriteString("\"", s.Cfg.DB.MongoDB[i].Addr[j], "\"")
			}
			code.MainWriteString("}, ",
				"\"", s.Cfg.DB.MongoDB[i].DBName, "\", ",
				"\"", s.Cfg.DB.MongoDB[i].ReplicaSet, "\", ",
				strconv.Itoa(s.Cfg.DB.MongoDB[i].Timeout), ", ",
				strconv.Itoa(s.Cfg.DB.MongoDB[i].Active), ", ",
				strconv.Itoa(s.Cfg.DB.MongoDB[i].Idle), ", ",
				strconv.Itoa(s.Cfg.DB.MongoDB[i].IdleTimeout),
				"); err != nil {\n",
				"\t\tpainc(err)\n",
				"\t}\n",
			)
		}
	}

	return nil
}
