package service

import (
	"errors"
	"github.com/cnlisea/seas/code"
	"strconv"
)

func (s *Service) Rpc(code *code.Code) error {
	if s.Cfg == nil {
		return nil
	}

	if code == nil {
		return errors.New("builder is nil")
	}

	code.MainWriteString("\t//service rpc\n")
	// rpc
	var (
		i              int
		offlineExist   bool
		rpcPackageName string
	)
	for i = range s.Cfg.Rpc {
		rpcPackageName = "rpc" + strconv.Itoa(i)
		code.ImportWriteString("\t", rpcPackageName, " \"", s.Cfg.Rpc[i].Path, "\"\n")
		if s.Cfg.Rpc[i].Offline != nil && !offlineExist {
			offlineExist = true
			code.MainWriteString("\tvar rpcState bool\n")
		}
		if s.Cfg.Rpc[i].Offline != nil {
			code.MainWriteString("\trpcState = ", strconv.FormatBool(*s.Cfg.Rpc[i].Offline), "\n")
		}
		code.MainWriteString("\tif err = a.NetRpcRegister(\"",
			s.Cfg.Name,
			"\", \"",
			s.Cfg.Listen.Ip,
			"\", ",
			strconv.FormatUint(uint64(s.Cfg.Listen.Port), 10),
			", \"",
			s.Cfg.GroupName,
			"\", ")
		// rpc offlice
		if s.Cfg.Rpc[i].Offline != nil {
			code.MainWriteString("&rpcState, ")
		} else {
			code.MainWriteString("nil, ")
		}
		code.MainWriteString(rpcPackageName, ".", s.Cfg.Rpc[i].Name, "(")
		if len(s.Cfg.Rpc[i].Param) > 0 {
			for j := range s.Cfg.Rpc[i].Param {
				if j > 0 {
					code.MainWriteString(", ")
				}
				switch s.Cfg.Rpc[i].Param[j] {
				case "db_mysql":
					code.MainWriteString("a.ProxyDBMySQL()")
				case "db_redis":
					code.MainWriteString("a.ProxyDBRedis()")
				case "db_mongo":
					code.MainWriteString("a.ProxyDBMongo()")
				case "rpc_client":
					code.MainWriteString("a.ProxyRpcClient()")
				case "cfg":
					code.MainWriteString("a.ProxyConfig()")
				case "mq":
					code.MainWriteString("a.ProxyMQ()")
				case "discovery":
					code.MainWriteString("a.ProxyDiscovery()")
				default:
					return errors.New(s.Cfg.Name + " invalid rpc param " + s.Cfg.Rpc[i].Param[j])
				}
			}
		}
		code.MainWriteString(")); err != nil {\n",
			"\t\tpanic(err)\n",
			"\t}\n")
	}
	return nil
}
