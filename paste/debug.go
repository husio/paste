package paste

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func handleInspectDB(ctx *Context, w http.ResponseWriter, r *http.Request) {
	entities := make([]interface{}, 0, 100)

	var irange util.Range
	if start := r.URL.Query().Get("start"); start != "" {
		irange.Start = []byte(start)
	}
	if end := r.URL.Query().Get("end"); end != "" {
		irange.Limit = []byte(end)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	enc := json.NewEncoder(w)

	limit := 50
	if raw := r.URL.Query().Get("limit"); raw != "" {
		l, err := strconv.Atoi(raw)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			enc.Encode("invalid limit value")
			return
		}
		limit = l
	}

	offset := 0
	if raw := r.URL.Query().Get("offset"); raw != "" {
		o, err := strconv.Atoi(raw)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			enc.Encode("invalid offset value")
			return
		}
		offset = o
	}

	start := time.Now()
	iter := ctx.app.db.NewIterator(&irange, nil)
	count := 0
	for ; iter.Next(); count++ {
		if count < offset || count-offset > limit {
			continue
		}
		var value interface{}
		chunks := strings.SplitN(string(iter.Key()), ":", 2)
		switch chunks[0] {
		case "user":
			var user User
			if err := proto.Unmarshal(iter.Value(), &user); err != nil {
				value = map[string]string{
					"key":   string(iter.Key()),
					"value": string(iter.Value()),
				}
			} else {
				value = user
			}
		case "paste":
			var paste Paste
			if err := proto.Unmarshal(iter.Value(), &paste); err != nil {
				value = map[string]string{
					"key":   string(iter.Key()),
					"value": string(iter.Value()),
				}
			} else {
				value = paste
			}
		default:
			value = map[string]string{
				"key":   string(iter.Key()),
				"value": string(iter.Value()),
			}
		}
		entities = append(entities, value)
	}
	iter.Release()

	content := map[string]interface{}{
		"count":    count,
		"limit":    limit,
		"offset":   offset,
		"entities": entities,
		"duration": time.Now().Sub(start).String(),
	}

	if err := iter.Error(); err != nil {
		content["error"] = err.Error()
	}

	if err := enc.Encode(content); err != nil {
		log.Printf("cannot encode response: %s", err)
	}
}

func handleInspectCache(ctx *Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(ctx.app.cache.mem); err != nil {
		log.Printf("cannot encode cache: %s", err)
	}
}
