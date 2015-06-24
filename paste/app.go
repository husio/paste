package paste

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
	"golang.org/x/oauth2"
)

type App struct {
	cache  *memoryCache
	oauth  map[string]*oauth2.Config
	db     *leveldb.DB
	routes []route
}

type route struct {
	method  string
	handler handler
	params  []string
	rx      *regexp.Regexp
}

type handler func(*Context, http.ResponseWriter, *http.Request)

func NewApp() *App {
	app := &App{
		cache: newMemoryCache(),
		oauth: make(map[string]*oauth2.Config),
	}

	app.handle("GET", "/", handleHello)
	app.handle("POST", "/", handleCreatePaste)
	app.handle("GET", "/login", handleLoginSelect)
	app.handle("GET", "/login/google", handleLoginGoogle)
	app.handle("GET", "/logout", handleLogout)
	app.handle("GET", "/:pasteID", handleGetPaste)
	app.handle("DELETE", "/:pasteID", handleDeletePaste)

	// TODO(husio) debug
	app.handle("GET", "/_/db", handleInspectDB)
	app.handle("GET", "/_/cache", handleInspectCache)

	return app
}

func (app *App) OauthCredentials(provider string, conf oauth2.Config) {
	app.oauth[provider] = &conf
}

func (app *App) ResetCache() {
	app.cache.Reset()
}

func (app *App) UseDatabase(path string) error {
	conn, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return fmt.Errorf("cannot open database: %s", err)
	}
	app.db = conn
	return nil
}

func (app *App) CloseDatabase() {
	if err := app.db.Close(); err != nil {
		log.Printf("cannot close database: %s", err)
	}
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if len(path) > 1 && path[len(path)-1] == '/' {
		path = strings.TrimRight(path, "/")
		http.Redirect(w, r, path, http.StatusMovedPermanently)
		return
	}
	for _, route := range app.routes {
		if route.method != r.Method {
			continue
		}
		var ctx *Context

		if len(route.params) == 0 {
			if !route.rx.MatchString(r.URL.Path) {
				continue
			}
			ctx = &Context{app: app}
		} else {
			match := route.rx.FindAllStringSubmatch(r.URL.Path, 1)
			if len(match) == 0 {
				continue
			}
			ctx = &Context{app: app}
			ctx.params = make(map[string]string)
			for i, name := range route.params {
				ctx.params[name] = match[0][i+1]
			}
		}

		if ctx != nil {
			route.handler(ctx, w, r)
			return
		}
	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func (app *App) handle(method string, path string, h handler) {
	params := make([]string, 0)
	chunks := strings.Split(path, "/")
	for i, chunk := range chunks {
		if strings.HasPrefix(chunk, ":") {
			params = append(params, chunk[1:])
			chunks[i] = "([^/]+)"
		}
	}
	rx := regexp.MustCompile("^" + strings.Join(chunks, "/") + "$")
	app.routes = append(app.routes, route{
		method:  method,
		handler: h,
		params:  params,
		rx:      rx,
	})
}

type Context struct {
	params map[string]string
	app    *App
}

// CurrentUserID return ID of user authenticated within given HTTP request.
// Returned ID is not validated for existance in database.
func (ctx *Context) CurrentUserID(r *http.Request) (string, bool) {
	sessionID, err := r.Cookie(sessionCookieName)
	if err != nil || sessionID.Value == "" {
		return "", false
	}
	userID, ok := ctx.app.cache.GetString("session:" + sessionID.Value)
	if !ok {
		return "", false
	}
	return userID, true
}

// CurrentUser return user authenticated within given HTTP request.
func (ctx *Context) CurrentUser(r *http.Request) (*User, bool) {
	userID, ok := ctx.CurrentUserID(r)
	if !ok {
		return nil, false
	}
	user, err := UserByID(ctx.app.db, userID)
	if err != nil {
		if err != ErrNotFound {
			log.Printf("database error: %s", err)
		}
		return nil, false
	}
	return user, true
}
