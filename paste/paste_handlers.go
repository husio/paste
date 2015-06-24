package paste

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func handleHello(ctx *Context, w http.ResponseWriter, r *http.Request) {
	user, _ := ctx.CurrentUser(r)
	render(w, "paste_form", user)
}

func handleGetPaste(ctx *Context, w http.ResponseWriter, r *http.Request) {
	paste, err := PasteByID(ctx.app.db, ctx.params["pasteID"])
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

func handleCreatePaste(ctx *Context, w http.ResponseWriter, r *http.Request) {
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
	if err := StorePaste(ctx.app.db, &paste); err != nil {
		log.Printf("cannot store paste: %s", err)
		renderErr(w, http.StatusInternalServerError)
		return
	}

	// if user is logged, bookmark created paste for him automatically.
	if userID, ok := ctx.CurrentUserID(r); ok {
		if err := BookmarkPaste(ctx.app.db, userID, paste.ID); err != nil {
			// database error is not critical for paste creation
			log.Printf("database error: %s", err)
		}
	}

	http.Redirect(w, r, "/"+paste.ID, http.StatusFound)
}

func handleDeletePaste(ctx *Context, w http.ResponseWriter, r *http.Request) {
	if err := DeletePaste(ctx.app.db, ctx.params["pasteID"]); err != nil {
		log.Printf("cannot delete paste: %s", err)
		renderErr(w, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
