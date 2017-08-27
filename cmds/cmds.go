package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	runApp(os.Args)
	return
}

func runApp(args []string) {
	var app = getApp("")
	err := app.Run(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func getApp(apiv string) *cli.App {
	app := cli.NewApp()
	app.Name = "beectl"
	app.Version = "1.0"
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "app version: %v\n", c.App.Version)
	}
	app.Usage = "A simple command line client for beedemo."

	if apiv == "" {
		app.Usage += "\n\n" +
		"WARNING:\n" +
		"   Environment variable ETCDCTL_API is not set; defaults to etcdctl v2.\n"
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "debug", Usage: "output cURL commands which can be used to reproduce the request"},
		cli.StringFlag{Name:"test.v", Usage:"v for test"},
		cli.StringFlag{Name:"test.run", Usage:"run for test"},
		cli.StringFlag{Name:"test.coverprofile", Usage:"coverprofile for test"},
		cli.StringFlag{Name:"test.outputdir", Usage:"outputdir for test"},
	}
	app.Commands = []cli.Command{
		cli.Command{},
	}

	return app
}