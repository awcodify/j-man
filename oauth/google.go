package oauth

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/awcodify/j-man/config"
)

const oauthGoogleURLAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

// User returned from Google API
type User struct {
	Email string `json:"email"`
}

// GenerateStateOauthCookie for getting the url of sign in page in google
func GenerateStateOauthCookie(cfg config.Config) (string, http.Cookie) {
	expiration := time.Now().Add(time.Duration(cfg.OAuth.Expiration) * 24 * time.Hour)
	buffer := make([]byte, 16)
	rand.Read(buffer)
	state := base64.URLEncoding.EncodeToString(buffer)
	cookie := http.Cookie{
		Name:    "oauthstate",
		Value:   state,
		Expires: expiration,
	}
	return state, cookie
}

// GetUserData will parse google response to user data
func GetUserData(code string, cfg config.Config) (*User, error) {
	oauthConfig := cfg.GetGoogleOAuthConfig()

	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}

	var user User
	err = getJSON(oauthGoogleURLAPI+token.AccessToken, &user)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	return &user, nil
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
