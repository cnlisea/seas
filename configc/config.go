package configc

import (
	"errors"
	"github.com/cnlisea/seas/code"
	"github.com/cnlisea/seas/config"
	"strconv"
)

type ConfigC struct {
	cfg *config.Config
}

func New(cfg *config.Config) *ConfigC {
	return &ConfigC{
		cfg: cfg,
	}
}

func (c *ConfigC) Run(buffer *code.Code) error {
	if c.cfg == nil {
		return nil
	}

	if buffer == nil {
		return errors.New("code is nil")
	}

	if c.cfg.Nacos != nil {
		if c.cfg.Nacos.NamespaceId == "" || len(c.cfg.Nacos.Nodes) == 0 {
			return errors.New("config nacos param invalid")
		}
		buffer.MainWriteString("\t//config center\n")

		buffer.MainWriteString("\tif err = a.ConfigCenter(\"")
		buffer.MainWriteString(c.cfg.Nacos.NamespaceId)
		buffer.MainWriteString("\", []*app.ConfigCenterNote{\n")
		for i := range c.cfg.Nacos.Nodes {
			buffer.MainWriteString("\t\t{\n")
			buffer.MainWriteString("\t\t\tAddr: \"")
			buffer.MainWriteString(c.cfg.Nacos.Nodes[i].Addr)
			buffer.MainWriteString("\",\n")
			buffer.MainWriteString("\t\t\tPort: ")
			buffer.MainWriteString(strconv.FormatUint(uint64(c.cfg.Nacos.Nodes[i].Port), 10))
			buffer.MainWriteString(",\n")
			buffer.MainWriteString("\t\t},\n")
		}
		buffer.MainWriteString("\t}); err != nil {\n")
		buffer.MainWriteString("\t\tpanic(err)\n")
		buffer.MainWriteString("\t}\n")
	}
	return nil
}
