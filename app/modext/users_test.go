package modext

import (
	"context"
	"testing"

	"github.com/awcodify/j-man/app/factories"
	"github.com/awcodify/j-man/config"
	"github.com/stretchr/testify/assert"
)

func TestFindByEmail(t *testing.T) {
	email := "findbyemail@email.com"
	config := config.New()
	ctx := context.Background()
	db, _ := config.ConnectDB()

	userFactory := factories.User{}
	userFactory.Email = email

	userFactory.Create(func() {
		user, _ := FindUserByEmail(ctx, db, email)
		assert.Equal(t, email, user.Email)

		_, err := FindUserByEmail(ctx, db, "email")
		assert.NotNil(t, err)
		assert.Equal(t, "sql: no rows in result set", err.Error())
	})
}
