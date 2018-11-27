// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// OpenPitrix Access Management service app.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	"openpitrix.io/iam/pkg/version"
)

func Main() {
	app := cli.NewApp()
	app.Name = "im"
	app.Usage = "im provides im service."
	app.Version = version.GetVersionString()

	app.UsageText = `im [global options] command [options] [args...]

EXAMPLE:
   im gen-config
   im info
   im list
   im ping
   im getv key
   im serve
   im tour`

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config",
			Value:  "im-config.json",
			Usage:  "im config file",
			EnvVar: "OPENPITRIX_AM_CONFIG",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "debug",
			Usage:  "debug app",
			Hidden: true,

			Action: func(c *cli.Context) {
				fmt.Println(nil)
				fmt.Println(version.GetVersion())
				return
			},
		},

		{
			Name:  "serve",
			Usage: "run as drone service",
			Action: func(c *cli.Context) {
				log.Fatal("TODO")
			},
		},

		{
			Name:  "tour",
			Usage: "show more eximples",
			Action: func(c *cli.Context) {
				fmt.Println(tourTopic)
			},
		},
	}

	app.CommandNotFound = func(ctx *cli.Context, command string) {
		fmt.Fprintf(ctx.App.Writer, "not found '%v'!\n", command)
	}

	app.Run(os.Args)
}

const tourTopic = `
im gen-config
`
