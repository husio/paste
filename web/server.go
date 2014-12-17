package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
	"time"
)

type dict map[string]interface{}

type Storage interface {
	Create(string, interface{}, time.Duration) error
	Get(string, interface{}) error
	Del(string) error
}

type Paste struct {
	Content     string
	ContentType string
	ValidTill   *time.Time
	RequireHost string
}

type httphandler struct {
	tmpl   *template.Template
	store  Storage
	pubsub *PubSub
}

func Serve(port int, templatesPath string, staticsPath string, store Storage) error {
	tmpl, err := template.ParseGlob(templatesPath + "/*.html")
	if err != nil {
		return err
	}
	pubsub := NewPubSub(32)
	handler := &httphandler{tmpl: tmpl, store: store, pubsub: pubsub}

	mux := http.NewServeMux()
	mux.Handle("/favicon.ico", http.FileServer(http.Dir(path.Join(staticsPath, "img/"))))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticsPath))))
	mux.Handle("/", handler)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}

func (h *httphandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == "/" && r.Method == "GET":
		h.handleMainPage(w, r)
	case r.URL.Path == "/" && r.Method == "POST":
		h.handlePasteCreate(w, r)
	case r.Method == "GET":
		h.handlePasteGet(w, r)
	default:
		http.NotFound(w, r)
	}

}

func (h *httphandler) render(w http.ResponseWriter, templateName string, context dict) {
	if err := h.tmpl.ExecuteTemplate(w, templateName, context); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (h *httphandler) handleMainPage(w http.ResponseWriter, r *http.Request) {
	ctx := dict{
		"PageTitle": "paste",
	}
	h.render(w, "main-page", ctx)
}

func (h *httphandler) handlePasteCreate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), 400)
	}
	key := genKey(5)
	paste := &Paste{
		Content:     r.FormValue("content"),
		ContentType: "plain/text",
		RequireHost: "",
	}

	if _, ok := r.Form["host-required"]; ok {
		paste.RequireHost = genKey(12)
		http.SetCookie(w, &http.Cookie{
			Name:    key,
			Value:   paste.RequireHost,
			Expires: time.Now().Add(30 * time.Second),
		})
	}

	var expAfter time.Duration
	if expireAfter, err := strconv.Atoi(r.PostFormValue("expire-after")); err != nil {
		http.Error(w, "Invalid \"expire-after\" value", 400)
		return
	} else {
		expAfter = time.Duration(expireAfter) * time.Second
		validTill := time.Now().Add(expAfter)
		paste.ValidTill = &validTill
	}

	// retry several times before giving up
	for i := 0; ; i++ {
		if err := h.store.Create(key, paste, expAfter); err == nil {
			break
		} else if i > 5 {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	http.Redirect(w, r, fmt.Sprintf("/%s", key), 302)
}

func (h *httphandler) handlePasteGet(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	paste := &Paste{}
	if err := h.store.Get(key, &paste); err != nil {
		http.NotFound(w, r)
		return
	}

	if cookie, err := r.Cookie(key); err == nil && cookie.Value == paste.RequireHost {
		http.SetCookie(w, &http.Cookie{
			Name:    key,
			Value:   "",
			Expires: time.Now().Add(-24 * time.Hour),
		})
		h.handleHostPaste(w, r, key, paste)
		return
	}

	if paste.ValidTill != nil && paste.RequireHost == "" {
		validDelta := paste.ValidTill.Sub(time.Now())
		if validDelta > 0 {
			cacheControl := fmt.Sprintf("public, max-age=%d", int(validDelta.Seconds()))
			w.Header().Set("Cache-Control", cacheControl)
			w.Header().Set("Etag", fmt.Sprintf("%d", paste.ValidTill.Unix()))
		}
	}

	if paste.RequireHost != "" {
		h.pubsub.Publish(key, &pastereader{Host: r.Host})
	}

	w.Header().Set("ContentType", paste.ContentType)
	fmt.Fprint(w, paste.Content)
}

func (h *httphandler) handleHostPaste(w http.ResponseWriter, r *http.Request, key string, paste *Paste) {
	defer h.pubsub.UnsubscribeAll(key)
	defer h.store.Del(key)

	flush := w.(http.Flusher).Flush
	connClosed := w.(http.CloseNotifier).CloseNotify()

	sub := h.pubsub.Subscribe(key)

	h.render(w, "paste-host", dict{"PasteKey": key, "Paste": paste})
	flush()
	for {
		select {
		case cli := <-sub.C:
			h.render(w, "paste-client", dict{"Client": cli})
			flush()
		case <-connClosed:
			return
		}
	}
}

type pastereader struct {
	Host string
}
