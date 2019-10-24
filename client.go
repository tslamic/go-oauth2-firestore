package fstorage

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/tslamic/go-oauth2-firestore/token"
	"time"
)

type fstore struct {
	c *firestore.Client
	n string
	t time.Duration
}

func (f *fstore) Put(token *token.Info) error {
	ctx, cancel := context.WithTimeout(context.Background(), f.t)
	defer cancel()
	_, _, err := f.c.Collection(f.n).Add(ctx, token)
	return err
}

func (f *fstore) Get(key string, val interface{}) (*token.Info, error) {
	ctx, cancel := context.WithTimeout(context.Background(), f.t)
	defer cancel()
	iter := f.c.Collection(f.n).Where(key, "==", val).Limit(1).Documents(ctx)
	doc, err := iter.Next()
	if err != nil {
		return nil, err
	}
	info := &token.Info{}
	err = doc.DataTo(info)
	return info, err
}

func (f *fstore) Del(key string, val interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), f.t)
	defer cancel()
	iter := f.c.Collection(f.n).Where(key, "==", val).Limit(1).Documents(ctx)
	doc, err := iter.Next()
	if err != nil {
		return err
	}
	_, err = f.c.Collection(f.n).Doc(doc.Ref.ID).Delete(ctx)
	return err
}

func (f *fstore) Close() error {
	return f.c.Close()
}
