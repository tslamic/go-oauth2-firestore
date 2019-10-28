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

func New(c *firestore.Client, collection string) *Client {
	return NewWithTimeout(c, collection, timeout)
}

func NewWithTimeout(c *firestore.Client, collection string, timeout time.Duration) *Client {
	fs := &fstore{c: c, n: collection, t: timeout}
	return &Client{c: fs}
}

type Client struct {
	c *fstore
}

func (f *Client) Create(info oauth2.TokenInfo) error {
	return f.c.Put(token(info))
}

func (f *Client) RemoveByCode(code string) error {
	return f.c.Del(keyCode, code)
}

func (f *Client) RemoveByAccess(access string) error {
	return f.c.Del(keyAccess, access)
}

func (f *Client) RemoveByRefresh(refresh string) error {
	return f.c.Del(keyRefresh, refresh)
}

func (f *Client) GetByCode(code string) (oauth2.TokenInfo, error) {
	return f.c.Get(keyCode, code)
}

func (f *Client) GetByAccess(access string) (oauth2.TokenInfo, error) {
	return f.c.Get(keyAccess, access)
}

func (f *Client) GetByRefresh(refresh string) (oauth2.TokenInfo, error) {
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
