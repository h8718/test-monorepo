package api

import (
	"log"

	"github.com/direktiv/direktiv/pkg/dlog"
	"github.com/direktiv/direktiv/pkg/refactor/core"
)

func RunApplication(config *core.Config) {
	logger, err := dlog.ApplicationLogger("api")
	if err != nil {
		log.Fatalf("can not get logger: %v", err)
	}

	s, err := NewServer(logger, config)
	if err != nil {
		logger.Errorf("can not create api server: %v", err)
	}

	err = s.Start()
	if err != nil {
		logger.Errorf("can not start api server: %v", err)
		log.Fatal(err.Error())
	}
}
