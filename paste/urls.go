package paste

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var Router *http.ServeMux

func init() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", handleHello).Methods("GET")
	r.HandleFunc("/", handleCreatePaste).Methods("POST")
	r.HandleFunc("/{pasteID}", handleGetPaste).Methods("GET")
	r.HandleFunc("/{pasteID}", handleDeletePaste).Methods("DELETE")

	r.HandleFunc("/login", handleLoginSelect).Methods("GET")
	r.HandleFunc("/login/google-oatuh2", handleLoginGoogleOauth2).Methods("GET")
	r.HandleFunc("/logout", handleLogout).Methods("GET")

	router := http.NewServeMux()
	router.Handle("/", timeHandler(r))
	Router = router
}

type timedResponseWriter struct {
	start time.Time
	set   bool
	w     http.ResponseWriter
}

func (tw *timedResponseWriter) Header() http.Header {
	return tw.w.Header()
}

func (tw *timedResponseWriter) Write(b []byte) (int, error) {
	tw.ensureHeader()
	return tw.w.Write(b)
}

func (tw *timedResponseWriter) WriteHeader(code int) {
	tw.ensureHeader()
	tw.w.WriteHeader(code)
}

func (tw *timedResponseWriter) ensureHeader() {
	work := time.Now().Sub(tw.start)
	tw.w.Header().Set("Server-Work-Time", work.String())
}

func timeHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tw := &timedResponseWriter{
			w:     w,
			start: time.Now(),
		}
		h.ServeHTTP(tw, r)
	})
}
