package handler

import (
	"context"
	"example-api/internal/api/http/server"
	item_dao "example-api/internal/database/dao/item"
	list_dao "example-api/internal/database/dao/list"
	"example-api/pkg/convert"
	"fmt"
)

func (h Handler) V1TodoListPost(ctx context.Context, req *server.ListCreateInput) (*server.List, error) {
	tx, err := h.pg.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	listCreateInputDao := list_dao.CreateInput{
		Name: req.Name,
	}

	listFromDao, err := h.listDao.Add(ctx, tx, listCreateInputDao)
	if err != nil {
		return nil, fmt.Errorf("add list: %w", err)
	}

	var items []server.Item
	if len(req.Items) > 0 {
		itemsCreateInputDao := convertItemsCreateInputToDao(listFromDao.ID, req.Items)
		itemsFromDao, err := h.itemDao.Add(ctx, tx, itemsCreateInputDao...)
		if err != nil {
			return nil, fmt.Errorf("add items: %w", err)
		}
		items = convert.Slice(itemsFromDao, convertItemFromDao)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("commit tx: %w", err)
	}

	list := &server.List{
		ID:    listFromDao.ID,
		Name:  listFromDao.Name,
		Items: items,
	}

	return list, nil
}

func convertItemsCreateInputToDao(listID int64, itemsCommonData []server.ItemCommonData) []item_dao.CreateInput {
	itemsCreateInput := make([]item_dao.CreateInput, 0, len(itemsCommonData))
	for _, itemCommonData := range itemsCommonData {
		itemCreateInput := convertItemCreateInputToDao(listID, itemCommonData.Title, itemCommonData.Description)
		itemsCreateInput = append(itemsCreateInput, itemCreateInput)
	}
	return itemsCreateInput
}
