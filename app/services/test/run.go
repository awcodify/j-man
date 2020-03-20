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
	running  string
	finished string
}

var (
	status = statusMapping{
		running:  "RUNNING",
		finished: "FINISHED",
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
	m, err := createCheckPoint(r, ctx)

	if err != nil {
		return nil, err
	}

	// Run the test
	err = runTest(ctx, r, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func createCheckPoint(r Runner, ctx context.Context) (*models.Round, error) {
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

func runTest(ctx context.Context, r Runner, m *models.Round) error {
	_, err := runner.Run(r.Config.App.JMeter.Path, r.Options)
	if err != nil {
		return err
	}

	// If test finised, update the checkpoint status to be FINISHED
	err = updateStatus(ctx, r, m)
	if err != nil {
		return err
	}

	return nil
}

func updateStatus(ctx context.Context, r Runner, m *models.Round) error {
	round, _ := models.FindRound(ctx, r.db, m.ID)
	round.Status = status.finished

	_, err := round.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}
