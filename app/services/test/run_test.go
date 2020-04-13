package test

import (
	"context"
	"testing"

	"github.com/awcodify/j-man/app/models"
	"github.com/awcodify/j-man/config"
	"github.com/awcodify/j-man/utils"
	"github.com/stretchr/testify/assert"
)

var (
	cfg, _ = config.New()
	id     = 0
)

// since we are using pointer in models, we need to define ID
func incrementalIdentity() int {
	id = id + 1
	return id
}

func TestRun(t *testing.T) {
	utils.CleanTablesAfter(cfg.DB.DSN, []string{"rounds"}, func() {
		cfg.JMeter.Path = "echo"
		db, _ := cfg.ConnectDB()
		round := models.Round{
			Users:       1,
			RampUp:      1,
			Duration:    1,
			Name:        "Test 1",
			Description: "Test running for first time",
		}
		r := Runner{
			Config:     cfg,
			DB:         db,
			Round:      &round,
			ResultPath: "test.csv",
		}

		t.Run("Run successfully", func(t *testing.T) {
			r.Round.ID = incrementalIdentity()
			round, err := r.Run(context.Background())
			assert.Nil(t, err)
			assert.Equal(t, "Test 1", round.Name)
			assert.Equal(t, status.aggregating, round.Status)
		})

		t.Run("Failed to create checkpoint", func(t *testing.T) {
			r.Round.ID = incrementalIdentity()
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
			rr := Runner{
				Config:     cfg,
				DB:         db,
				ResultPath: "test.failed",
				Round:      &round,
			}

			round, err := rr.Run(context.Background())

			assert.Equal(t,
				"Should use .jmx or .csv file",
				err.Error())
			assert.Equal(t, status.failed, round.Status)
		})

		t.Run("Failed to update status", func(t *testing.T) {
			r.Round.ID = incrementalIdentity()

			// To make checkpoint failed, we create status which have more than 20 chars
			status.aggregating = "It's a sentence with more than twenty chars"

			_, err := r.Run(context.Background())

			assert.Equal(t,
				"models: unable to update rounds row: pq: value too long for type character varying(20)",
				err.Error())

			status.finished = "AGGREGATING"
		})
	})
}
