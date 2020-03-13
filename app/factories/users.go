package factories

import (
	"context"

	"github.com/awcodify/j-man/app/models"
	"github.com/awcodify/j-man/config"
	"github.com/awcodify/j-man/utils"
	"github.com/khaiql/dbcleaner"
	"github.com/khaiql/dbcleaner/engine"
	"github.com/volatiletech/sqlboiler/boil"
)

var (
	cfg, _  = config.New()
	db, _   = cfg.ConnectDB()
	Cleaner = dbcleaner.New()
)

type User struct {
	models.User
}

func (u User) Create(fn func()) {
	psql := engine.NewPostgresEngine(cfg.DB.DSN)
	Cleaner.SetEngine(psql)
	Cleaner.Acquire("users")
	Cleaner.Clean("users")

	if u.Email == "" {
		u.Email = "example@email.com"
	}

	if u.PasswordDigest == "" {
		u.PasswordDigest = "passwordHash"
	}

	err := u.Insert(context.Background(), db, boil.Infer())
	utils.DieIf(err)

	fn()

	Cleaner.Clean("users")
}
