# j-man [![codecov](https://codecov.io/gh/awcodify/j-man/branch/master/graph/badge.svg)](https://codecov.io/gh/awcodify/j-man) ![Go](https://github.com/awcodify/j-man/workflows/Go/badge.svg)
All in One JMeter Manager

## Development
 ```bash 
 ❯ cp config.yaml.example config.development.yaml
 ❯ cp air.conf.example air.conf
 ❯ air -c air.conf
```
## Guide
### Database & Migrations
* To create migration file, we are using [sql-migrate](https://github.com/rubenv/sql-migrate).
* We are a "database-first" oriented. Use [sqlboiler](https://github.com/volatiletech/sqlboiler) to generate models (inside `app/models`).
  * make sure you are on the root folder of project
  * `sqlboiler psql --wipe -o app/models`
* For extending the models to create helper or custom function, please put it inside `app/modext`

### Usage
```Go
package main

import (
	"flag"
	"log"

	"github.com/awcodify/j-man/aggregator"
	"github.com/awcodify/j-man/runner"
	"github.com/awcodify/j-man/utils"
)

var (
	jmeterPath, scriptPath, resultPath string
	users, rampUp, duration            int64
)

func init() {
	flag.StringVar(&jmeterPath, "jmeterPath", "bin/jmeter", "Location of executable JMeter")
	flag.StringVar(&scriptPath, "scriptPath", "./scripts/google.jmx", "Location of testing script")
	flag.StringVar(&resultPath, "resultPath", "./results/log.csv", "Where the result file will be stored")
	flag.Int64Var(&users, "users", 1, "How many users needed to test")
	flag.Int64Var(&rampUp, "rampUp", 1, "Amount of time Jmeter should take to get all the threads sent for the execution")
	flag.Int64Var(&duration, "duration", 1, "How the test will be running? (in miliseconds)")
}

func main() {
	flag.Parse()

	options := runner.Options{
		JMeterPath:     jmeterPath,
		ScriptPath:     scriptPath,
		ResultFilePath: resultPath,
		Users:          users,
		RampUp:         rampUp,
		Duration:       duration,
	}
	resultFilePath, err := runner.Run(options)
	utils.DieIf(err)

	result := aggregator.Collect(resultFilePath).ToResult().Aggregate()

	log.Println(result)
}

```

It will return something like:
```bash
~/Documents/playground/j-man
❯ go run cmd/jmanager/main.go
Creating summariser <summary>
Created the tree successfully using ./scripts/google.jmx
Starting the test @ Mon Feb 24 15:20:09 WIB 2020 (1582532409595)
Waiting for possible Shutdown/StopTestNow/Heapdump message on port 4445
summary =      1 in 00:00:05 =    0.2/s Avg:  4667 Min:  4667 Max:  4667 Err:     0 (0.00%)
Tidying up ...    @ Mon Feb 24 15:20:14 WIB 2020 (1582532414711)
... end of run
2020/02/24 15:20:15 {{1514.8333333333333 3341}}
```
