package lib

import (
	"fmt"
	"time"
	log "github.com/Sirupsen/logrus"
	"github.com/vixns/gomarathon"
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
	Sync        bool
}

func (s *Orchestrator) Check() error {
	log.Debugf("service check %s\n", s.ApiHost)

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

func (s *Orchestrator) HasDeploymentID (deploymentID string) (ret bool, err error) {
	if deployments, err := s.Client.GetDeployments(); err != nil {
		return false, err
	} else {
		log.Debugf("deployments:  %#v\n", deployments)
		for d := range deployments {
			if  deployments[d].ID == deploymentID {
				return true, nil
			}
		}
	}
	return false, nil
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
		if s.Sync {
			time.Sleep(2 * time.Second)
			c := time.Tick(time.Second * 2)
			for i := range c {
				log.Debugf("tick: %s", i)
				isDeploying, err := s.HasDeploymentID(deploymentID);
				if err != nil {
					return "", err
				}
				if !isDeploying {
					time.Sleep(2 * time.Second)
					break
				}
			}
		} else {
			println("deploymentID:", deploymentID)
			log.Debugf("version: %s", version)
		}
		return deploymentID, err
	}
}
