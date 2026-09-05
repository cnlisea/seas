package discovery

import (
	"errors"
	"github.com/cnlisea/seas/code"
	"github.com/cnlisea/seas/config"
	"strconv"
)

type Discovery struct {
	Cfg *config.Discovery
}

func New(cfg *config.Discovery) *Discovery {
	return &Discovery{
		Cfg: cfg,
	}
}

func (d *Discovery) Run(buffer *code.Code) error {
	if d.Cfg == nil {
		return nil
	}

	if buffer == nil {
		return errors.New("builder is nil")
	}

	if d.Cfg.Nacos != nil {
		if d.Cfg.Nacos.NamespaceId == "" || len(d.Cfg.Nacos.Nodes) == 0 {
			return errors.New("discovery nacos param invalid")
		}

		buffer.MainWriteString("\t//discovery\n")

		buffer.MainWriteString("\tif err = a.Discovery(\"")
		buffer.MainWriteString(d.Cfg.Nacos.NamespaceId)
		buffer.MainWriteString("\", []*app.DiscoveryNode{\n")
		for i := range d.Cfg.Nacos.Nodes {
			buffer.MainWriteString("\t\t{\n")
			buffer.MainWriteString("\t\t\tAddr: \"")
			buffer.MainWriteString(d.Cfg.Nacos.Nodes[i].Addr)
			buffer.MainWriteString("\",\n")
			buffer.MainWriteString("\t\t\tPort: ")
			buffer.MainWriteString(strconv.FormatUint(uint64(d.Cfg.Nacos.Nodes[i].Port), 10))
			buffer.MainWriteString(",\n")
			buffer.MainWriteString("\t\t},\n")
		}
		buffer.MainWriteString("\t}); err != nil {\n")
		buffer.MainWriteString("\t\tpainc(err)\n")
		buffer.MainWriteString("\t}\n")
	}
	return nil
}
