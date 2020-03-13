package views

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/awcodify/j-man/app/models"
	"github.com/awcodify/j-man/app/services/oauth"
	"github.com/awcodify/j-man/config"
	"github.com/awcodify/j-man/utils"
	"github.com/elliotchance/redismock"
	"github.com/stretchr/testify/assert"
)

func TestHandleSignIn(t *testing.T) {
	cfg := config.New()
	v := View{Config: cfg}

	req, err := http.NewRequest("GET", "/sign_in", nil)
	utils.DieIf(err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(v.HandleSignIn)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusTemporaryRedirect, rr.Code)
}

func TestAuthenticate(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	s.Set("session_token", "expected")

	cfg := config.New()
	cfg.Redis.Host = s.Addr()
	cache := cfg.ConnectRedis()

	fakeCache := redismock.NewNiceMock(cache)

	v := View{Config: cfg, Cache: fakeCache}

	params := url.Values{
		"state":    {"oauth"},
		"code":     {"token"},
		"authuser": {"1"},
		"scope":    {"email+https://www.googleapis.com/auth/userinfo.email+openid&"},
	}

	t.Run("Authentication success", func(t *testing.T) {
		rr := httptest.NewRecorder()
		http.SetCookie(rr, &http.Cookie{Name: "oauthstate", Value: "oauth"})
		req := &http.Request{URL: &url.URL{RawQuery: params.Encode()}, Header: http.Header{"Cookie": rr.HeaderMap["Set-Cookie"]}}

		// mock oauth.GetUserData
		getUserData = func(code string, cfg config.Config) (*oauth.User, error) {
			return &oauth.User{Email: "example@test.com"}, nil
		}

		// mock modext.FindUserByEmail
		findUserByEmail = func(ctx context.Context, db *sql.DB, email string) (*models.User, error) {
			return &models.User{Email: "example@test.com"}, nil
		}

		v.Authenticate(rr, req)

		assert.Equal(t, http.StatusSeeOther, rr.Code)
	})

	t.Run("oauthState not match", func(t *testing.T) {
		rr := httptest.NewRecorder()
		http.SetCookie(rr, &http.Cookie{Name: "oauthstate", Value: "not-match"})
		req := &http.Request{URL: &url.URL{RawQuery: params.Encode()}, Header: http.Header{"Cookie": rr.HeaderMap["Set-Cookie"]}}

		v.Authenticate(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
		assert.Equal(t, "Invalid oauth google state", rr.Body.String())
	})

	t.Run("Failed to get user data from exchange", func(t *testing.T) {
		rr := httptest.NewRecorder()
		http.SetCookie(rr, &http.Cookie{Name: "oauthstate", Value: "oauth"})
		req := &http.Request{URL: &url.URL{RawQuery: params.Encode()}, Header: http.Header{"Cookie": rr.HeaderMap["Set-Cookie"]}}

		// mock oauth.GetUserData
		getUserData = func(code string, cfg config.Config) (*oauth.User, error) {
			return nil, fmt.Errorf("")
		}

		v.Authenticate(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
		assert.Equal(t, "Failed on getting user info from Google", rr.Body.String())
	})

	t.Run("User not registered yet", func(t *testing.T) {
		rr := httptest.NewRecorder()
		http.SetCookie(rr, &http.Cookie{Name: "oauthstate", Value: "oauth"})
		req := &http.Request{URL: &url.URL{RawQuery: params.Encode()}, Header: http.Header{"Cookie": rr.HeaderMap["Set-Cookie"]}}

		// mock oauth.GetUserData
		getUserData = func(code string, cfg config.Config) (*oauth.User, error) {
			return &oauth.User{Email: "example@test.com"}, nil
		}

		// mock modext.FindUserByEmail
		findUserByEmail = func(ctx context.Context, db *sql.DB, email string) (*models.User, error) {
			return nil, fmt.Errorf("")
		}

		v.Authenticate(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
		assert.Equal(t, "Not registered. Please, contact admin.", rr.Body.String())
	})

}
