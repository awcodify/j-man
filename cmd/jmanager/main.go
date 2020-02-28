package main

import (
	"flag"
	"log"
	"os"

	"github.com/awcodify/j-man/aggregator"
	"github.com/awcodify/j-man/config"
	"github.com/awcodify/j-man/runner"
	"github.com/awcodify/j-man/utils"
)

var (
	jmeterPath, scriptPath, resultPath string
	users, rampUp, duration            int64
)

func init() {
	flag.StringVar(&jmeterPath, "jmeterPath", "bin/jmeter", "Location of executable JMeter")
	flag.StringVar(&scriptPath, "scriptPath", "", "Location of testing script")
	flag.StringVar(&resultPath, "resultPath", "./results/log.csv", "Where the result file will be stored")
	flag.Int64Var(&users, "users", 1, "How many users needed to test")
	flag.Int64Var(&rampUp, "rampUp", 1, "Amount of time Jmeter should take to get all the threads sent for the execution")
	flag.Int64Var(&duration, "duration", 1, "How the test will be running? (in miliseconds)")
}

func main() {
	flag.Parse()

	cfg := config.New()

	if len(os.Args) < 2 {
		flag.PrintDefaults()
	} else {
		options := runner.Options{
			ScriptPath:     scriptPath,
			ResultFilePath: resultPath,
			Users:          users,
			RampUp:         rampUp,
			Duration:       duration,
		}
		resultFilePath, err := runner.Run(cfg.App.JMeter.Path, options)
		utils.DieIf(err)

		result := aggregator.Collect(resultFilePath).ToResult().Aggregate()

		log.Println(result)
	}
}
