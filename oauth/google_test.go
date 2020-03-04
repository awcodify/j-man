package oauth

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/awcodify/j-man/config"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2/google"
)

var cfg = config.Config{
	OAuth: config.OAuth{
		Expiration: 1,
	},
}

func TestGenerateStateOauthCookie(t *testing.T) {
	actualState, actualCookie := GenerateStateOauthCookie(cfg)

	assert.Equal(t, actualState, actualCookie.Value)
}

func TestGetUserData(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	token := "token"
	tokenGoogleResponse := fmt.Sprintf(`{"access_token":"%s", "token_type": "jwt", "refresh_token": "refresh_token"}`, token)

	t.Run("Positive scenario", func(t *testing.T) {

		httpmock.RegisterResponder("POST", "=~"+google.Endpoint.TokenURL+"(.*)",
			httpmock.NewStringResponder(200, tokenGoogleResponse))

		httpmock.RegisterResponder("GET", oauthGoogleURLAPI+token,
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewJsonResponse(200, map[string]interface{}{
					"email":         "example@test.com",
					"id":            1,
					"verfied_email": true,
				})
			},
		)

		actualData, actualError := GetUserData(token, cfg)

		expectedData := User{Email: "example@test.com"}
		assert.Equal(t, &expectedData, actualData)
		assert.Nil(t, actualError)

	})

	t.Run("Wrong code exchange", func(t *testing.T) {
		httpmock.RegisterResponder("POST", "=~"+google.Endpoint.TokenURL+"(.*)",
			httpmock.NewStringResponder(400, tokenGoogleResponse))

		_, actualError := GetUserData(token, cfg)
		expectedError := "code exchange wrong: oauth2: cannot fetch token: 400\nResponse: " + tokenGoogleResponse
		assert.Equal(t, expectedError, actualError.Error())
	})

	t.Run("Failed getting user info", func(t *testing.T) {
		httpmock.RegisterResponder("POST", "=~"+google.Endpoint.TokenURL+"(.*)",
			httpmock.NewStringResponder(200, tokenGoogleResponse))

		httpmock.RegisterResponder("GET", oauthGoogleURLAPI+token, func(req *http.Request) (*http.Response, error) {
			return nil, fmt.Errorf("Client closed")
		})

		_, actualError := GetUserData(token, cfg)

		assert.NotNil(t, actualError)
		assert.Equal(t, "failed getting user info: Get https://www.googleapis.com/oauth2/v2/userinfo?access_token=token: Client closed", actualError.Error())
	})
}
