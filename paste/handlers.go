package paste

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	render(w, "paste_form", http.StatusOK, nil)
}

func handleGetPaste(w http.ResponseWriter, r *http.Request) {
	paste, err := PasteByID(mux.Vars(r)["pasteID"])
	if err != nil {
		if err == ErrNotFound {
			renderErr(w, http.StatusNotFound)
		} else {
			log.Printf("database error: %s", err)
			renderErr(w, http.StatusInternalServerError)
		}
		return
	}
	w.Write(paste.Content)
}

func handleCreatePaste(w http.ResponseWriter, r *http.Request) {
	var expireIn int64
	if raw := r.PostFormValue("expire-in"); raw == "" {
		expireIn = 0
	} else {
		n, err := strconv.Atoi(raw)
		if err != nil || n < 0 {
			renderErr(w, http.StatusBadRequest)
			return
		}
		expireIn = int64(n) * int64(time.Second)
	}

	paste := Paste{
		Content:  []byte(r.PostFormValue("content")),
		ExpireIn: expireIn,
	}
	if err := StorePaste(&paste); err != nil {
		log.Printf("cannot store paste: %s", err)
		renderErr(w, http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/"+paste.ID, http.StatusFound)
}

func handleDeletePaste(w http.ResponseWriter, r *http.Request) {
	if err := DeletePaste(mux.Vars(r)["pasteID"]); err != nil {
		log.Printf("cannot delete paste: %s", err)
		renderErr(w, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func handleLoginSelect(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func handleLoginGoogleOauth2(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}
