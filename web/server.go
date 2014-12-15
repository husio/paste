package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type dict map[string]interface{}

type Storage interface {
	Create(string, interface{}, time.Duration) error
	Get(string, interface{}) error
}

type Paste struct {
	Content string
}

type httphandler struct {
	tmpl  *template.Template
	store Storage
}

func Serve(port int, templatesPath string, staticsPath string, store Storage) error {
	tmpl, err := template.ParseGlob(templatesPath + "/*.html")
	if err != nil {
		return err
	}
	handler := &httphandler{tmpl: tmpl, store: store}

	mux := http.NewServeMux()
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
		Content: r.FormValue("content"),
	}

	expireAfter, err := strconv.Atoi(r.PostFormValue("expire-after"))
	if err != nil {
		http.Error(w, "Invalid \"expire-after\" value", 400)
		return
	}

	// retry several times before giving up
	for i := 0; ; i++ {
		if err := h.store.Create(key, paste, time.Duration(expireAfter)*time.Second); err == nil {
			break
		} else if i > 5 {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	http.Redirect(w, r, fmt.Sprintf("/%s", key), 302)
}

func (h *httphandler) handlePasteGet(w http.ResponseWriter, r *http.Request) {
	paste := &Paste{}
	if err := h.store.Get(r.URL.Path[1:], &paste); err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, paste.Content)
}
