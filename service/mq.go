package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/cnlisea/seas/code"
)

func (s *Service) MQ(code *code.Code) error {
	if s.Cfg == nil || len(s.Cfg.MQ) == 0 {
		return nil
	}

	if code == nil {
		return errors.New("code is nil")
	}

	code.MainWriteString("\t//service mq\n")

	var (
		i, j, k              int
		nsBuf                strings.Builder
		subImportExist       bool
		subImportPackageName string
	)
	for i = range s.Cfg.MQ {
		nsBuf.Reset()
		nsBuf.WriteString("[]string{")
		for j = range s.Cfg.MQ[i].NameServer {
			if j > 0 {
				nsBuf.WriteString(", ")
			}
			nsBuf.WriteString("\"")
			nsBuf.WriteString(s.Cfg.MQ[i].NameServer[j])
			nsBuf.WriteString("\"")
		}
		nsBuf.WriteString("}")
		// product
		for j = range s.Cfg.MQ[i].Producer {
			code.MainWriteString("\tif err = a.MQProducerRegister(",
				"\"", s.Cfg.MQ[i].Producer[j].Name, "\", ",
				"\"", s.Cfg.MQ[i].AccessKey, "\", ",
				"\"", s.Cfg.MQ[i].SecretKey, "\", ",
				nsBuf.String(), ", ",
				"\"", s.Cfg.MQ[i].Namespace, "\", ",
				"\"", s.Cfg.MQ[i].Producer[j].GroupId, "\"); err != nil {\n",
				"\t\tpanic(err)\n",
				"\t}\n",
			)
		}

		// subscribe
		for j = range s.Cfg.MQ[i].Subscribe {
			if s.Cfg.MQ[i].Subscribe[j].BatchSize <= 0 {
				s.Cfg.MQ[i].Subscribe[j].BatchSize = 1
			}
			if !subImportExist {
				subImportExist = true
				code.ImportWriteString("\t\"github.com/cnlisea/ant/app/proxy\"\n")
			}
			code.MainWriteString("\tif err = a.MQConsumerRegister(",
				"\"", s.Cfg.MQ[i].Subscribe[j].Name, "\", ",
				"\"", s.Cfg.MQ[i].AccessKey, "\", ",
				"\"", s.Cfg.MQ[i].SecretKey, "\", ",
				nsBuf.String(), ", ",
				"\"", s.Cfg.MQ[i].Namespace, "\", ",
				"\"", s.Cfg.MQ[i].Subscribe[j].GroupId, "\", ",
				strconv.FormatBool(s.Cfg.MQ[i].Subscribe[j].Broadcast), ", ",
				strconv.Itoa(s.Cfg.MQ[i].Subscribe[j].BatchSize), ", ",
				"[]*proxy.MQConsumerSubscribe{\n",
			)
			for k = range s.Cfg.MQ[i].Subscribe[j].Config {
				code.MainWriteString("\t\t{\n",
					"\t\t\tTopic: \"", s.Cfg.MQ[i].Subscribe[j].Config[k].Topic, "\",\n",
					"\t\t\tTag: \"", s.Cfg.MQ[i].Subscribe[j].Config[k].Tag, "\",\n",
					"\t\t\tHandler: ",
				)
				subImportPackageName = fmt.Sprintf("mqSubscribe_%d_%d_%d", i, j, k)
				code.ImportWriteString("\t", subImportPackageName, " \"", s.Cfg.MQ[i].Subscribe[j].Config[k].Path, "\"\n")

				code.MainWriteString(subImportPackageName, ".", s.Cfg.MQ[i].Subscribe[j].Config[k].Name, ",\n")
				code.MainWriteString("\t\t},\n")
			}
			code.MainWriteString("\t}); err != nil {",
				"\t\tpanic(err)\n",
				"\t}\n",
			)
		}
	}
	return nil
}
