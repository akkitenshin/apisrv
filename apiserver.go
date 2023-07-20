package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIserver) Start() error {
	if err := s.configureLogger(); err != nil{
		return err
	}

	s.logger.Info("starting api server")

	return nil
}

func (s *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.Loglevel)
	if err != nil{
		return err
	}
}
