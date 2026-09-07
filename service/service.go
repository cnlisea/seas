package service

import (
	"errors"
	"github.com/cnlisea/seas/code"
	"github.com/cnlisea/seas/config"
)

type Service struct {
	Cfg *config.Service
}

func New(cfg *config.Service) *Service {
	return &Service{
		Cfg: cfg,
	}
}

func (s *Service) Run(code *code.Code) error {
	if s.Cfg == nil {
		return nil
	}

	if code == nil {
		return errors.New("code is nil")
	}
	code.MainWriteString("\t//service\n")

	var err error
	// flag
	if err = s.Flag(code); err != nil {
		return errors.New("service flag fail: " + err.Error())
	}

	// config
	if err = s.Config(code); err != nil {
		return errors.New("service config fail: " + err.Error())
	}

	// log
	if err = s.Log(code); err != nil {
		return errors.New("service log fail: " + err.Error())
	}

	//db
	if err = s.DB(code); err != nil {
		return errors.New("service db fail: " + err.Error())
	}

	//mq
	if err = s.MQ(code); err != nil {
		return errors.New("service mq fail: " + err.Error())
	}

	//rpc
	if err = s.Rpc(code); err != nil {
		return errors.New("service rpc fail: " + err.Error())
	}

	return nil
}
