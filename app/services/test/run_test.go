package test

import (
	"context"
	"testing"

	"github.com/awcodify/j-man/config"
	"github.com/awcodify/j-man/runner"
	"github.com/awcodify/j-man/utils"
	"github.com/stretchr/testify/assert"
)

var cfg, _ = config.New()

func TestRun(t *testing.T) {
	utils.CleanTablesAfter(cfg.DB.DSN, []string{"rounds"}, func() {
		cfg.App.JMeter.Path = "echo"
		db, _ := cfg.ConnectDB()
		options := runner.Options{
			Users:          1,
			RampUp:         1,
			Duration:       1,
			ResultFilePath: "test.csv",
		}
		r := Runner{
			Config:      cfg,
			db:          db,
			Name:        "Test 1",
			Description: "Test running for first time",
			Options:     options,
		}

		t.Run("Run successfully", func(t *testing.T) {
			round, err := r.Run(context.Background())
			assert.Nil(t, err)
			assert.Equal(t, "Test 1", round.Name)
		})

		t.Run("Failed to create checkpoint", func(t *testing.T) {
			// To make checkpoint failed, we create status which have more than 20 chars
			status.running = "It's a sentence with more than twenty chars"

			_, err := r.Run(context.Background())

			assert.Equal(t,
				"models: unable to insert into rounds: pq: value too long for type character varying(20)",
				err.Error())

			// Bring back the correct status
			status.running = "RUNNING"
		})

		t.Run("Runner failed to run", func(t *testing.T) {
			o := runner.Options{
				ResultFilePath: "test.failed",
			}
			rr := Runner{
				Config:  cfg,
				db:      db,
				Name:    "Test 1",
				Options: o,
			}

			_, err := rr.Run(context.Background())

			assert.Equal(t,
				"Should use .jmx or .csv file",
				err.Error())
		})

		t.Run("Failed to update status", func(t *testing.T) {

			// To make checkpoint failed, we create status which have more than 20 chars
			status.finished = "It's a sentence with more than twenty chars"

			_, err := r.Run(context.Background())

			assert.Equal(t,
				"models: unable to update rounds row: pq: value too long for type character varying(20)",
				err.Error())

			status.finished = "FINISHED"
		})
	})
}
