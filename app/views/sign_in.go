package views

import (
	"fmt"
	"net/http"

	"github.com/awcodify/j-man/app/oauth"
	"github.com/awcodify/j-man/config"
)

// Handler for request handler
type Handler struct {
	Config config.Config
}

// HandleSignIn will redirect to google oauth url
func (h Handler) HandleSignIn(w http.ResponseWriter, r *http.Request) {
	oauthState, oauthCookie := oauth.GenerateStateOauthCookie(h.Config)

	http.SetCookie(w, &oauthCookie)

	oauthConfig := h.Config.GetGoogleOAuthConfig()
	url := oauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Authenticate will verify authentication from google
func (h Handler) Authenticate(w http.ResponseWriter, r *http.Request) {
	oauthState, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthState.Value {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid oauth google state"))
		return
	}

	data, err := oauth.GetUserData(r.FormValue("code"), h.Config)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Fprintf(w, "UserInfo: %s\n", data)
}
