package modext

import (
	"context"
	"database/sql"

	"github.com/awcodify/j-man/app/models"
)

func FindUserByEmail(ctx context.Context, db *sql.DB, email string) (*models.User, error) {
	return models.Users(models.UserWhere.Email.EQ(email)).One(ctx, db)
}
