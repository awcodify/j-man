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
	db          *sql.DB
	Name        string
	Description string
	Options     runner.Options
}

func (r Runner) Run(ctx context.Context) (*models.Round, error) {
	// Create checkpoint to track the test
	m, err := r.createCheckPoint(ctx)
	if err != nil {
		return m, err
	}

	// Run the test
	err = r.triggerJMeter(ctx, m)
	if err != nil {
		return m, err
	}

	return m, nil
}

func (r Runner) createCheckPoint(ctx context.Context) (*models.Round, error) {
	var m models.Round
	m.Status = status.running
	m.Name = r.Name
	m.Description = r.Description
	m.Users = r.Options.Users
	m.RampUp = r.Options.RampUp
	m.Duration = r.Options.Duration

	err := m.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (r Runner) triggerJMeter(ctx context.Context, m *models.Round) error {
	_, err := runner.Run(r.Config.App.JMeter.Path, r.Options)
	if err != nil {

		// If failed to run JMeter, mark the round as failed
		if updateErr := r.updateStatus(ctx, m, status.failed); updateErr != nil {
			return updateErr
		}

		return err
	}
	// If test finised, update the checkpoint status to be FINISHED
	if err = r.updateStatus(ctx, m, status.aggregating); err != nil {
		return err
	}

	return nil
}

func (r Runner) updateStatus(ctx context.Context, m *models.Round, status string) error {
	m.Status = status

	_, err := m.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}
