package handler

import (
	"context"
	"crypto/rsa"
	item_dao "example-api/internal/database/dao/item"
	list_dao "example-api/internal/database/dao/list"
	"example-api/pkg/pg"
)

type listDao interface {
	Add(ctx context.Context, q pg.Queryable, createInput list_dao.CreateInput) (list_dao.List, error)
	CheckExists(ctx context.Context, q pg.Queryable, id int64) (bool, error)
	List(ctx context.Context, q pg.Queryable, filter list_dao.ListFilter) ([]list_dao.List, error)
}

type itemDao interface {
	Add(ctx context.Context, q pg.Queryable, createInputs ...item_dao.CreateInput) ([]item_dao.Item, error)
	List(ctx context.Context, q pg.Queryable, filter item_dao.ListFilter) ([]item_dao.Item, error)
}

type Handler struct {
	email      string
	password   string
	privateKey *rsa.PrivateKey

	pg pg.PG

	listDao listDao
	itemDao itemDao
}

func New(email, password string, privateKey *rsa.PrivateKey, pg pg.PG, listDao listDao, itemDao itemDao) Handler {
	return Handler{
		email:      email,
		password:   password,
		privateKey: privateKey,
		pg:         pg,
		listDao:    listDao,
		itemDao:    itemDao,
	}
}
