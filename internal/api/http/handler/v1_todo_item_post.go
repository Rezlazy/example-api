package handler

import (
	"context"
	"example-api/internal/api/http/server"
	item_dao "example-api/internal/database/dao/item"
	"fmt"
)

func (h Handler) V1TodoItemPost(ctx context.Context, req *server.ItemCreateInput) (*server.Item, error) {
	itemCreateInputDao := convertItemCreateInputToDao(req.ListID, req.Title, req.Description)

	listIsExists, err := h.listDao.CheckExists(ctx, h.pg, req.ListID)
	if err != nil {
		return nil, fmt.Errorf("check exists list: %w", err)
	}

	if !listIsExists {
		return nil, fmt.Errorf("list %d not exist", req.ListID)
	}

	itemsFromDao, err := h.itemDao.Add(ctx, h.pg, itemCreateInputDao)
	if err != nil {
		return nil, fmt.Errorf("add item: %w", err)
	}
	if len(itemsFromDao) != 1 {
		return nil, fmt.Errorf("length items after add is equal %d", len(itemsFromDao))
	}
	itemFromDao := itemsFromDao[0]

	item := convertItemFromDao(itemFromDao)
	return &item, nil
}

func convertItemCreateInputToDao(listID int64, title string, description server.OptString) item_dao.CreateInput {
	var descriptionDao *string
	if description.IsSet() {
		descriptionDao = &description.Value
	}

	return item_dao.CreateInput{
		Title:       title,
		Description: descriptionDao,
		ListID:      listID,
	}
}

func convertItemFromDao(itemFromDao item_dao.Item) server.Item {
	var descriptionResponse server.OptString
	if itemFromDao.Description != nil {
		descriptionResponse = server.NewOptString(*itemFromDao.Description)
	}

	return server.Item{
		ID:          itemFromDao.ID,
		Title:       itemFromDao.Title,
		Description: descriptionResponse,
		ListID:      itemFromDao.ListID,
	}
}
