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
	b, err := db.Get(key("paste:%s", id), nil)
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
	tx.Put(key("paste:%s", p.ID), b)
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
		tx.Delete(key("paste:%s", iter.Value()))
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

// UserByID return from database user with given ID.
func UserByID(db *leveldb.DB, id string) (*User, error) {
	b, err := db.Get(key("user:%s", id), nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("database error: %s", err)
	}

	var user User
	if err := proto.Unmarshal(b, &user); err != nil {
		return nil, fmt.Errorf("cannot unmarshal: %s", err)
	}
	return &user, nil
}

// UserByOauth return from database user with given Oauth ID.
func UserByOauth(db *leveldb.DB, oauthID string) (*User, error) {
	userID, err := db.Get(key("oauth:%s:user", oauthID), nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("database error: %s", err)
	}

	b, err := db.Get(key("user:%s", userID), nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			log.Printf("inconsistend database: oauth %q leads to unknown user %q", oauthID, b)
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("database error: %s", err)
	}

	var user User
	if err := proto.Unmarshal(b, &user); err != nil {
		return nil, fmt.Errorf("cannot unmarshal: %s", err)
	}
	return &user, nil
}

// StoreUser stores given user by ID. It does not create any other relations.
func StoreUser(db *leveldb.DB, u *User) error {
	if u.ID == "" {
		u.ID = NewKey(16)
	}
	if u.CreatedAt == 0 {
		u.CreatedAt = time.Now().UnixNano()
	}

	b, err := proto.Marshal(u)
	if err != nil {
		return fmt.Errorf("cannot serialize: %s", err)
	}
	if err := db.Put(key("user:%s", u.ID), b, nil); err != nil {
		return fmt.Errorf("database error: %s", err)
	}
	return nil
}

// LinkOauthToUser store in database link between oauth ID and user ID. Link
// allows to get user ID for given oauth ID.
func LinkOauthToUser(db *leveldb.DB, oauthID, userID string) error {
	if err := db.Put(key("oauth:%s:user", oauthID), []byte(userID), nil); err != nil {
		return fmt.Errorf("database error: %s", err)
	}
	return nil
}

// BookmarkPaste create relation user-paste by creating entry with user ID and
// current time to given paste ID. Created link allows to get user bookmarks,
// ordered by creation date.
func BookmarkPaste(db *leveldb.DB, userID, pasteID string) error {
	nowUnix := time.Now().UnixNano()
	if err := db.Put(key("user:%s:paste:%d", userID, nowUnix), []byte(pasteID), nil); err != nil {
		return fmt.Errorf("database error: %s", err)
	}
	return nil
}

// ListBookmarkedPastes return list of pastes that user with given ID has
// bookmarked, that creation date is older than gieven time. Pastes are
// ordered descending, sorted by bookmark creation time. Maximum 100 pastes are
// returned.
//
// In case of pagination, to get next 100 pastes, use date of the oldest
// returned previously paste to narrow listing result.
func ListBookmarkedPastes(db *leveldb.DB, userID string, olderThan time.Time) ([]*Paste, error) {
	return nil, errors.New("not implemented")
}
