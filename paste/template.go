package paste

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/husio/paste/assets"
)

var templates map[string]*template.Template

func init() {
	templates = make(map[string]*template.Template)
	for _, name := range assets.AssetNames() {
		if !strings.HasPrefix(name, "template/") {
			continue
		}
		t, err := template.New(name).Parse(string(assets.MustAsset(name)))
		if err != nil {
			log.Panicf("cannot parse %q template: %s", name, err)
		}
		// cut off template/ prefix and .tmpl suffix
		short := name[9 : len(name)-5]
		templates[short] = t
	}
}

func render(w http.ResponseWriter, name string, context interface{}) {
	t, ok := templates[name]
	if !ok {
		log.Printf("template does not exist: %s", name)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if err := t.Execute(w, context); err != nil {
		log.Printf("cannot execute %q template: %s", name, err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func renderErr(w http.ResponseWriter, code int) {
	context := map[string]interface{}{
		"Text": http.StatusText(code),
		"Code": code,
	}
	w.WriteHeader(code)
	render(w, "error", context)
}
