package main

import (
	"fmt"
	"log"
	"os"

	"github.com/awcodify/j-man/aggregator"
	"github.com/awcodify/j-man/app"
	"github.com/awcodify/j-man/config"
	"github.com/awcodify/j-man/runner"
	"github.com/urfave/cli/v2"
)

var (
	env, jmeterPath, scriptPath, resultPath string
	users, rampUp, duration                 int64
)

func main() {
	cmd := &cli.App{
		Name:  "jman",
		Usage: "Do performance testing with JMeter",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "environment",
				Aliases:     []string{"env", "e"},
				Destination: &env,
				Required:    true,
			},
		},
		Commands: defineCommands(),
	}

	err := cmd.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func defineCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:   "run",
			Usage:  "Non-GUI mode",
			Action: cliAction,
			Flags:  defineNonGUIFlags(),
		},
		{
			Name:  "app",
			Usage: "GUI mode",
			Action: func(c *cli.Context) error {
				return app.Run()
			},
		},
	}
}

func defineNonGUIFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "script",
			Aliases:     []string{"s"},
			Destination: &scriptPath,
			Usage:       "Location of testing script",
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "output",
			Aliases:     []string{"o"},
			Destination: &resultPath,
			Value:       "./results/log.csv",
			Usage:       "Where the result file will be stored",
			Required:    true,
		},
		&cli.Int64Flag{
			Name:        "users",
			Aliases:     []string{"u"},
			Destination: &users,
			Usage:       "How many users needed to test",
			Required:    true,
		},
		&cli.Int64Flag{
			Name:        "rampup",
			Aliases:     []string{"r"},
			Destination: &rampUp,
			Usage:       "Amount of time Jmeter should take to get all the threads sent for the execution",
			Required:    true,
		},
		&cli.Int64Flag{
			Name:        "duration",
			Aliases:     []string{"d"},
			Destination: &duration,
			Usage:       "How the test will be running? (in miliseconds)",
			Required:    true,
		},
	}
}

func cliAction(c *cli.Context) error {
	if c.NumFlags() == 0 {
		cli.ShowAppHelp(c)
		return nil
	}

	os.Setenv("APP_ENV", env)

	cfg, err := config.New()
	if err != nil {
		return err
	}

	if cfg.App.JMeter.Path == "" {
		return fmt.Errorf("JMeter not found")
	}

	options := runner.Options{
		ScriptPath:     scriptPath,
		ResultFilePath: resultPath,
		Users:          users,
		RampUp:         rampUp,
		Duration:       duration,
	}
	resultFilePath, err := runner.Run(cfg.App.JMeter.Path, options)
	if err != nil {
		return err
	}

	result := aggregator.Collect(resultFilePath).ToResult().Aggregate()

	fmt.Println(result)

	return nil

}
