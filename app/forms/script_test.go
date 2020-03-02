package forms

import (
	"context"
	"testing"
	"time"

	"github.com/awcodify/j-man/app/models"
	"github.com/awcodify/j-man/config"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null"
	"gopkg.in/khaiql/dbcleaner.v2"
	"gopkg.in/khaiql/dbcleaner.v2/engine"
)

var Cleaner = dbcleaner.New()

type ScriptSuite struct {
	suite.Suite
}

func (suite *ScriptSuite) SetupSuite() {
	// Init and set mysql cleanup engine
	cfg := config.New()
	mysql := engine.NewPostgresEngine(cfg.DB.DSN)
	Cleaner.SetEngine(mysql)
}

func (suite *ScriptSuite) SetupTest() {
	Cleaner.Acquire("scripts")
}

func (suite *ScriptSuite) TearDownTest() {
	Cleaner.Clean("scripts")
}

func (suite *ScriptSuite) TestCreate() {
	cfg := config.New()
	category := Category{
		Name: "hello",
	}
	var scp models.Script
	scp.Name = "Hello"
	scp.Content = "World"
	scp.CreatedAt = null.NewTime(time.Now(), true)
	scp.UpdatedAt = null.NewTime(time.Now(), true)
	script := Script{
		Category: category,
		Script:   scp,
	}

	actual, err := script.Create(context.Background(), cfg)

	suite.Nil(err)
	suite.Equal(script, actual)
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(ScriptSuite))
}
