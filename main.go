package main

import (
	"fmt"
	"os"

	"github.com/rcw5/pipeline-validator/commands"
	"github.com/urfave/cli"
)

var version string = "dev"

func main() {
	app := cli.NewApp()
	app.Name = "pipeline-validator"
	app.Usage = "Validate a Concourse pipeline and its vars"
	app.Version = version

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "pipeline, p",
			Usage: "Pipeline definition",
		},
		cli.StringSliceFlag{
			Name:  "load-vars-from, l",
			Usage: "Vars (secrets) file to load",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.String("pipeline") == "" {
			fmt.Println("ERROR: --pipeline must be specified")
			cli.ShowAppHelpAndExit(c, 1)
		}
		if len(c.StringSlice("load-vars-from")) == 0 {
			fmt.Println("ERROR: --load-vars-from must be specified")
			cli.ShowAppHelpAndExit(c, 1)
		}
		pipelinePath := c.String("pipeline")
		varsPaths := c.StringSlice("load-vars-from")

		err := commands.ValidatePipeline(pipelinePath, varsPaths...)
		if err != nil {
			return cli.NewExitError(err, 1)
		}
		fmt.Println("Looks good!")
		return nil
	}

	app.Run(os.Args)
}
