package lib

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/fpgeek/gomarathon"
)

type Orchestrator struct {
	ApiHost     string
	Service     string
	DockerImage string
	Scale       int
	Mem         float64
	Cpu         float64
	Client      *gomarathon.Client
	App         *gomarathon.Application
}

func (s *Orchestrator) Check() error {
	println("service check ", s.ApiHost)

	if s.ApiHost == "" {
		return fmt.Errorf("MARATHON_URL env is missing")
	}

	s.Client, _ = gomarathon.NewClient(s.ApiHost, nil)

	if app, err := s.Client.GetApp(s.Service); err != nil {
		return err
	} else {
		s.App = app
	}

	return nil
}

func (s *Orchestrator) Deploy() (deploymentID string, err error) {

	App := &gomarathon.Application{}
	if s.Mem > 0 {
		App.Mem = s.Mem
	}
	if s.Scale > 0 {
		App.Instances = s.Scale
	}
	if s.Cpu > 0 {
		App.CPUs = s.Cpu
	}
	if s.DockerImage != "" {
		App.Container = s.App.Container
		App.Container.Docker.Image = s.DockerImage
	}

	deploymentID, version, err := s.Client.UpdateApp(s.Service, App)
	if err != nil {
		return "", err
	} else {
		log.Debugf("deploymentID: %s\n", deploymentID)
		log.Debugf("version: %s\n", version)
		return deploymentID, err
	}

	return
}
