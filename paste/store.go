package paste

import (
	"crypto/rand"
	"encoding/base32"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

var ErrNotFound = errors.New("not found")

// PasteByID return paste with given id or ErrNotFound if not exist.
func PasteByID(db *leveldb.DB, id string) (*Paste, error) {
	b, err := db.Get(key("paste:%s:id", id), nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("database error: %s", err)
	}

	var paste Paste
	if err := proto.Unmarshal(b, &paste); err != nil {
		return nil, fmt.Errorf("cannot unmarshal: %s", err)
	}

	if paste.ExpireIn > 0 {
		now := time.Now().UnixNano()
		if paste.CreatedAt+paste.ExpireIn < now {
			// paste is expired, delete it manually, but ignore errors - they
			// are not relevant
			if err := DeletePaste(db, id); err != nil {
				log.Printf("cannot delete expired paste: %s", err)
			}
			return nil, ErrNotFound
		}
	}

	return &paste, nil
}

func StorePaste(db *leveldb.DB, p *Paste) error {
	if p.ID == "" {
		p.ID = NewKey(16)
	}
	if p.CreatedAt == 0 {
		p.CreatedAt = time.Now().UnixNano()
	}

	b, err := proto.Marshal(p)
	if err != nil {
		return fmt.Errorf("cannot serialize: %s", err)
	}
	tx := &leveldb.Batch{}
	tx.Put(key("paste:%s:id", p.ID), b)
	if p.ExpireIn > 0 {
		expire := p.CreatedAt + p.ExpireIn
		tx.Put(key("paste:%d:expire", expire), []byte(p.ID))
	}
	if err := db.Write(tx, nil); err != nil {
		return fmt.Errorf("db batch error: %s", err)
	}
	return nil
}

func DeletePaste(db *leveldb.DB, id string) error {
	if err := db.Delete(key("paste:%s:id", id), nil); err != nil {
		return fmt.Errorf("database error: %s", err)
	}
	return nil
}

// DeleteExpiredPastes removes all paste that "CreatedAt + ExpireIn" value is
// less than current UNIX time in nanoseconds.
//
// Deleted keys are
// * paste:<paste-id>:id
// * paste:<expire-at>:expire
func DeleteExpiredPastes(db *leveldb.DB) error {
	now := time.Now()
	iter := db.NewIterator(&util.Range{
		Start: []byte("paste:0:expire"),
		Limit: key("paste:%d:expire", now.UnixNano()),
	}, nil)
	tx := &leveldb.Batch{}
	for iter.Next() {
		tx.Delete(iter.Key())
		tx.Delete(key("paste:%s:id", iter.Value()))
	}
	iter.Release()
	if err := iter.Error(); err != nil {
		return fmt.Errorf("iter error: %s", err)
	}
	if err := db.Write(tx, nil); err != nil {
		return fmt.Errorf("db batch error: %s", err)
	}
	if tx.Len() > 0 {
		work := time.Now().Sub(now)
		log.Printf("delete expired pastes: %d removed in %s", tx.Len()/2, work)
	}
	return nil
}

func NewKey(size int) string {
	b := make([]byte, size)
	if _, err := rand.Read(b); err != nil {
		panic("cannot read random data")
	}
	s := base32.HexEncoding.EncodeToString(b)
	return strings.TrimRight(s, "=")
}

func key(format string, args ...interface{}) []byte {
	s := fmt.Sprintf(format, args...)
	return []byte(s)
}
