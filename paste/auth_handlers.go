package paste

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

const sessionCookieName = "s"

func handleLoginSelect(ctx *Context, w http.ResponseWriter, r *http.Request) {
	if _, ok := ctx.CurrentUserID(r); ok {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	provider := r.URL.Query().Get("provider")
	if provider != "" {
		if oauth, ok := ctx.app.oauth[provider]; ok {
			url := oauth.AuthCodeURL(NewKey(12))
			http.Redirect(w, r, url, http.StatusFound)
			return
		}
	}

	providers := make([]string, 0)
	for name := range ctx.app.oauth {
		providers = append(providers, name)
	}
	context := map[string]interface{}{
		"Providers": providers,
	}
	render(w, "login", context)
}

func handleLogout(ctx *Context, w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func handleLoginGoogle(ctx *Context, w http.ResponseWriter, r *http.Request) {
	oauth := ctx.app.oauth["google"]

	authcode := r.FormValue("code")
	tok, err := oauth.Exchange(oauth2.NoContext, authcode)
	if err != nil {
		log.Printf("cannot get oauth2 token: %s", err)
		url := oauth.AuthCodeURL(NewKey(12))
		http.Redirect(w, r, url, http.StatusFound)
		return
	}
	client := oauth.Client(oauth2.NoContext, tok)

	resp, err := client.Get(`https://www.googleapis.com/oauth2/v3/userinfo`)
	if err != nil {
		log.Printf("transaction failure: %s", err)
		http.Error(w, "transaction failure", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	var info struct {
		OauthID string `json:"sub"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
		Email   string `json:"email"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		log.Printf("oauth server error: %s", err)
		http.Error(w, "oauth server error", http.StatusInternalServerError)
		return
	}

	user, err := UserByOauth(ctx.app.db, info.OauthID)
	switch err {
	case nil:
	case ErrNotFound:
		user = &User{
			Name:    info.Name,
			Picture: info.Picture,
			Email:   info.Email,
		}
		if err := StoreUser(ctx.app.db, user); err != nil {
			log.Printf("cannot store user: %s", err)
			code := http.StatusInternalServerError
			http.Error(w, http.StatusText(code), code)
			return
		}
		if err := LinkOauthToUser(ctx.app.db, info.OauthID, user.ID); err != nil {
			log.Printf("cannot link user to oauth: %s", err)
			code := http.StatusInternalServerError
			http.Error(w, http.StatusText(code), code)
			return
		}
	default:
		log.Printf("database error: %s", err)
		code := http.StatusInternalServerError
		http.Error(w, http.StatusText(code), code)
		return
	}

	sessionID := NewKey(24)
	ctx.app.cache.Put("session:"+sessionID, user.ID)
	http.SetCookie(w, &http.Cookie{
		Name:    sessionCookieName,
		Value:   sessionID,
		Path:    "/",
		Expires: time.Now().Add(24 * time.Hour * 7),
	})
	http.Redirect(w, r, "/", http.StatusFound)
}
