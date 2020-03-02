package forms

import (
	"context"

	"github.com/awcodify/j-man/app/models"
	"github.com/awcodify/j-man/config"
	"github.com/volatiletech/sqlboiler/boil"
)

// Category will seperate script per category
type Category struct {
	Name string
}

// Script is testing script
type Script struct {
	Script   models.Script
	Category Category
}

// Create will store script to database
func (script Script) Create(ctx context.Context, cfg config.Config) (Script, error) {
	db, err := cfg.ConnectDB()

	err = script.Script.Insert(ctx, db, boil.Infer())

	return script, err
}
