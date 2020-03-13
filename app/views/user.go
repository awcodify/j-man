package views

import (
	"net/http"

	"github.com/awcodify/j-man/app/modext"
	"github.com/awcodify/j-man/app/services/oauth"
	"github.com/awcodify/j-man/utils"
)

var (
	findUserByEmail = modext.FindUserByEmail
	getUserData     = oauth.GetUserData
)

// HandleSignIn will redirect to google oauth url
func (v View) HandleSignIn(w http.ResponseWriter, r *http.Request) {
	oauthState, oauthCookie := oauth.GenerateStateOauthCookie(v.Config)

	http.SetCookie(w, &oauthCookie)

	oauthConfig := v.Config.GetGoogleOAuthConfig()
	url := oauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Authenticate will verify authentication from google
func (v View) Authenticate(w http.ResponseWriter, r *http.Request) {
	oauthState, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthState.Value {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid oauth google state"))
		return
	}

	data, err := getUserData(r.FormValue("code"), v.Config)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Failed on getting user info from Google"))
		return
	}

	user, err := findUserByEmail(v.Ctx, v.DB, data.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Not registered. Please, contact admin."))
		return
	}

	err = v.Cache.Set(r.FormValue("code"), user.Email, 0).Err()
	utils.DieIf(err)

	cookie := http.Cookie{Name: "session_token", Value: r.FormValue("code"), Path: "/"}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/run", http.StatusSeeOther)
}
