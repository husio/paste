package main

import (
	"flag"
	"log"
	"time"

	"github.com/husio/paste/store"
	"github.com/husio/paste/web"
)

func main() {
	var httpPort *int = flag.Int("port", 8000, "HTTP server port")
	var keyLen *int = flag.Int("key-length", 5, "Paste key length")
	flag.Parse()

	s := store.NewMemoryStore(1 * time.Minute)
	defer s.Close()
	log.Fatal(web.Serve(*httpPort, s, *keyLen))
}
