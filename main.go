package main

import (
	"flag"
	"log"
	"path"
	"runtime"
	"time"

	"github.com/husio/paste/store"
	"github.com/husio/paste/web"
)

var rootDir string

func init() {
	_, filename, _, _ := runtime.Caller(1)
	rootDir = path.Dir(filename)
}

func main() {
	var httpPort *int = flag.Int("port", 8000, "HTTP server port")
	var templatesDir *string = flag.String("templates-dir", path.Join(rootDir, "template"), "Templates directory path")
	var staticsDir *string = flag.String("statics-dir", path.Join(rootDir, "static"), "Static files directory path")
	flag.Parse()

	s := store.NewMemoryStore(1 * time.Minute)
	defer s.Close()
	log.Fatal(web.Serve(*httpPort, *templatesDir, *staticsDir, s))
}
