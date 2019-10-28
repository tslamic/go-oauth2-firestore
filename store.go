package fstorage

import (
	"cloud.google.com/go/firestore"
	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/models"
	"time"
)

const (
	keyCode    = "Code"
	keyAccess  = "Access"
	keyRefresh = "Refresh"

	timeout = 30 * time.Second
)

// New returns a new Firestore token store.
func New(c *firestore.Client, collection string) oauth2.TokenStore {
	return NewWithTimeout(c, collection, timeout)
}

// NewWithTimeout returns a new Firestore token store.
// Any Firestore operation will be cancelled if it surpasses the provided timeout.
func NewWithTimeout(c *firestore.Client, collection string, timeout time.Duration) oauth2.TokenStore {
	fs := &store{c: c, n: collection, t: timeout}
	return &client{c: fs}
}

type client struct {
	c *store
}

func (f *client) Create(info oauth2.TokenInfo) error {
	return f.c.Put(token(info))
}

func (f *client) RemoveByCode(code string) error {
	return f.c.Del(keyCode, code)
}

func (f *client) RemoveByAccess(access string) error {
	return f.c.Del(keyAccess, access)
}

func (f *client) RemoveByRefresh(refresh string) error {
	return f.c.Del(keyRefresh, refresh)
}

func (f *client) GetByCode(code string) (oauth2.TokenInfo, error) {
	return f.c.Get(keyCode, code)
}

func (f *client) GetByAccess(access string) (oauth2.TokenInfo, error) {
	return f.c.Get(keyAccess, access)
}

func (f *client) GetByRefresh(refresh string) (oauth2.TokenInfo, error) {
	return f.c.Get(keyRefresh, refresh)
}

func token(info oauth2.TokenInfo) *models.Token {
	return &models.Token{
		ClientID:         info.GetClientID(),
		UserID:           info.GetUserID(),
		RedirectURI:      info.GetRedirectURI(),
		Scope:            info.GetScope(),
		Code:             info.GetCode(),
		CodeCreateAt:     info.GetCodeCreateAt(),
		CodeExpiresIn:    info.GetCodeExpiresIn(),
		Access:           info.GetAccess(),
		AccessCreateAt:   info.GetAccessCreateAt(),
		AccessExpiresIn:  info.GetAccessExpiresIn(),
		Refresh:          info.GetRefresh(),
		RefreshCreateAt:  info.GetRefreshCreateAt(),
		RefreshExpiresIn: info.GetRefreshExpiresIn(),
	}
}
