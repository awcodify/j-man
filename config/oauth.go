package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"strings"
)

//OAuth for google sign in
type OAuth struct {
	GoogleClientID     string `yaml:"google_client_id"`
	GoogleClientSecret string `yaml:"google_client_secret"`
	CallbackURL        string `yaml:"callback_url"`
	Scopes             string `yaml:"scope"`
	Expiration         int    `yaml:"expiration"`
	Endpoint           string
}

// GetGoogleOAuthConfig will parse config for google oauth
func (config Config) GetGoogleOAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  config.OAuth.CallbackURL,
		ClientID:     config.OAuth.GoogleClientID,
		ClientSecret: config.OAuth.GoogleClientSecret,
		Scopes:       strings.Split(config.OAuth.Scopes, ","),
		Endpoint:     google.Endpoint,
	}
}
