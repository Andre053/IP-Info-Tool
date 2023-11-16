package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	run()
}

func run() {
	app := &cli.App{
		Name:     "IP Address Tool",
		Version:  "v2",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Andre053",
				Email: "andre.61122@proton.me",
			},
		},
		Usage: "Tools for gathering information about IP addresses",
		Commands: []*cli.Command{
			{
				Name:    "IP Info",
				Aliases: []string{"info", "ipi"},
				Usage:   "Gather public data of an IP address",
				Action: func(*cli.Context) error {
					return ipInfo()
				},
			},
			{
				Name:    "Distance",
				Aliases: []string{"d", "dist"},
				Usage:   "Calculate the kilometers between two IP address locations",
				Action: func(*cli.Context) error {
					return ipDistance()
				},
			},
			{
				Name:    "Ping",
				Aliases: []string{"p"},
				Usage:   "Ping IP addresses or hostnames (unprivileged, requires enablement in OS)",
				Action: func(*cli.Context) error {
					return ping(false)
				},
			},
			{
				Name:    "Ping (privileged)",
				Aliases: []string{"pp", "ppr", "ppriv"},
				Usage:   "Ping IP addresses or hostnames (requires root or setcap)",
				Action: func(*cli.Context) error {
					return ping(true)
				},
			},
		}}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
