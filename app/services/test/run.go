package test

import (
	"context"
	"database/sql"

	"github.com/awcodify/j-man/app/models"
	"github.com/awcodify/j-man/config"
	"github.com/awcodify/j-man/runner"
	"github.com/volatiletech/sqlboiler/boil"
)

type statusMapping struct {
	running     string
	cancelled   string
	aggregating string
	finished    string
	failed      string
}

var (
	status = statusMapping{
		running:     "RUNNING",
		cancelled:   "CANCELLED",
		aggregating: "AGGREGATING",
		finished:    "FINISHED",
		failed:      "FAILED",
	}
)

type Runner struct {
	*config.Config
	*sql.DB
	*models.Round
	ResultPath string
}

func (r Runner) Run(ctx context.Context) (*models.Round, error) {
	// Create checkpoint to track the test
	err := r.createCheckPoint(ctx)
	if err != nil {
		return r.Round, err
	}

	// Run the test
	err = r.triggerJMeter(ctx)
	if err != nil {
		return r.Round, err
	}

	return r.Round, nil
}

func (r Runner) createCheckPoint(ctx context.Context) error {
	r.Round.Status = status.running
	err := r.Round.Insert(ctx, r.DB, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func (r Runner) triggerJMeter(ctx context.Context) error {
	options := runner.Options{
		Users:          r.Round.Users,
		RampUp:         r.Round.RampUp,
		Duration:       r.Round.Duration,
		ResultFilePath: r.ResultPath,
	}
	_, err := runner.Run(r.Config.JMeter.Path, options)
	if err != nil {

		// If failed to run JMeter, mark the round as failed
		if updateErr := r.updateStatus(ctx, r.Round, status.failed); updateErr != nil {
			return updateErr
		}

		return err
	}
	// If test finised, update the checkpoint status to be FINISHED
	if err = r.updateStatus(ctx, r.Round, status.aggregating); err != nil {
		return err
	}

	return nil
}

func (r Runner) updateStatus(ctx context.Context, m *models.Round, status string) error {
	m.Status = status

	_, err := m.Update(ctx, r.DB, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}
