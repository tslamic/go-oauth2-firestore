package fstorage

import (
	"cloud.google.com/go/firestore"
	"context"
	"gopkg.in/oauth2.v3/models"
	"time"
)

type fstore struct {
	c *firestore.Client
	n string
	t time.Duration
}

func (f *fstore) Put(token *models.Token) error {
	ctx, cancel := context.WithTimeout(context.Background(), f.t)
	defer cancel()
	_, _, err := f.c.Collection(f.n).Add(ctx, token)
	return err
}

func (f *fstore) Get(key string, val interface{}) (*models.Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), f.t)
	defer cancel()
	iter := f.c.Collection(f.n).Where(key, "==", val).Limit(1).Documents(ctx)
	doc, err := iter.Next()
	if err != nil {
		return nil, err
	}
	info := &models.Token{}
	err = doc.DataTo(info)
	return info, err
}

func (f *fstore) Del(key string, val interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), f.t)
	defer cancel()
	return f.c.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		query := f.c.Collection(f.n).Where(key, "==", val).Limit(1)
		iter := tx.Documents(query)
		doc, err := iter.Next()
		if err != nil {
			return err
		}
		return tx.Delete(doc.Ref)
	})
}

func (f *fstore) Close() error {
	return f.c.Close()
}
