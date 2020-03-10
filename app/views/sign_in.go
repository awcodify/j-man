package views

import (
	"fmt"
	"net/http"

	"github.com/awcodify/j-man/app/oauth"
)

// HandleSignIn will redirect to google oauth url
func (cfg Config) HandleSignIn(w http.ResponseWriter, r *http.Request) {
	oauthState, oauthCookie := oauth.GenerateStateOauthCookie(cfg.Config)

	http.SetCookie(w, &oauthCookie)

	oauthConfig := cfg.GetGoogleOAuthConfig()
	url := oauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Authenticate will verify authentication from google
func (cfg Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	oauthState, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthState.Value {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid oauth google state"))
		return
	}

	data, err := oauth.GetUserData(r.FormValue("code"), cfg.Config)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Fprintf(w, "UserInfo: %s\n", data)
}
