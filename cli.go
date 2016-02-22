package main

import (
	"github.com/codegangsta/cli"
	"github.com/yuzu-co/doploy/lib"
	"log"
	"os"
	"strconv"
)

func main() {
	app := cli.NewApp()
	app.Name = "doploy"
	app.Version = "1.0.1"
	app.Usage = "Cli to update Marathon apps"
	app.Commands = []cli.Command{
		{
			Name:    "update",
			Aliases: []string{"up"},
			Usage:   "Update an app",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "scale",
					Usage: "Nb of instances",
				},
				cli.StringFlag{
					Name:  "mem",
					Usage: "Total memory allowed",
				},
				cli.StringFlag{
					Name:  "cpu",
					Usage: "Total cpu allowed",
				},
				cli.StringFlag{
					Name:  "image",
					Usage: "Set docker image",
				},
				cli.StringFlag{
					Name:  "sync",
					Usage: "Wait for deployment end",
				},
			},
			Action: func(c *cli.Context) {
				var service string = c.Args().First()

				o := lib.Orchestrator{
					ApiHost: os.Getenv("MARATHON_URL"),
					Service: service}

				err := o.Check()
				if err != nil {
					log.Fatal(err)
				}

				if c.String("scale") != "" {
					o.Scale, _ = strconv.Atoi(c.String("scale"))
				}

				if c.String("image") != "" {
					o.DockerImage = c.String("image")
				}

				if c.String("mem") != "" {
					o.Mem, _ = strconv.ParseFloat(c.String("mem"), 64)
				}

				if c.String("cpu") != "" {
					o.Cpu, _ = strconv.ParseFloat(c.String("cpu"), 64)
				}

				if c.String("sync") != "" {
					o.Sync = true;
				}

				o.Deploy()

			},
		},
	}

	app.Run(os.Args)
}
