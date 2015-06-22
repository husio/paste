package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/husio/paste/assets"
	"github.com/husio/paste/paste"
)

//go:generate go-bindata -pkg assets -ignore=\\..* -prefix assets -o assets/assets.go assets/template/... assets/static/...
//go:generate protoc paste/entities.proto --go_out=.

func main() {
	log.SetFlags(log.Lshortfile)
	log.SetOutput(os.Stdout)

	var (
		staticsFl = flag.String("static", "", "Path to statics")
	)
	flag.Parse()

	if err := paste.OpenDB("/tmp/notepad.leveldb"); err != nil {
		log.Fatalf("cannot open database: %s", err)
	}
	defer paste.CloseDB()

	if *staticsFl != "" {
		log.Printf("using statics directly from directory: %s", *staticsFl)
		fs := http.FileServer(http.Dir(*staticsFl))
		http.Handle("/static/", http.StripPrefix("/static", fs))
	} else {
		statics := assetfs.AssetFS{
			Asset:    assets.Asset,
			AssetDir: assets.AssetDir,
			Prefix:   "/static",
		}
		http.Handle("/static/", http.StripPrefix("/static", http.FileServer(&statics)))
	}

	go func() {
		for {
			paste.DeleteExpiredPastes()
			time.Sleep(time.Minute)
		}
	}()

	http.Handle("/", paste.Router)
	if err := http.ListenAndServe("localhost:8000", nil); err != nil {
		log.Fatalf("HTTP server error: %s", err)
	}
}
