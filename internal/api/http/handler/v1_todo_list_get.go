package handler

import (
	"context"
	"example-api/internal/api/http/server"
	item_dao "example-api/internal/database/dao/item"
	list_dao "example-api/internal/database/dao/list"
	"example-api/pkg/convert"
	"example-api/pkg/logger"
	"fmt"
)

func (h Handler) V1TodoListGet(ctx context.Context, params server.V1TodoListGetParams) (server.ListListResponseSchema, error) {
	listListFilterDao := list_dao.ListFilter{
		IDs: params.Ids,
	}

	listsDao, err := h.listDao.List(ctx, h.pg, listListFilterDao)
	if err != nil {
		return nil, fmt.Errorf("list object list from dao: %w", err)
	}

	listIDs := make([]int64, 0, len(listsDao))
	listMap := make(map[int64]server.List, len(listsDao))
	for _, listDao := range listsDao {
		id := listDao.ID
		listIDs = append(listIDs, id)
		list := convertListFromDao(listDao)
		listMap[id] = list
	}

	itemListFilterDao := item_dao.ListFilter{
		ListIDs: listIDs,
	}
	itemsDao, err := h.itemDao.List(ctx, h.pg, itemListFilterDao)
	if err != nil {
		return nil, fmt.Errorf("list items from dao: %w", err)
	}

	for _, itemDao := range itemsDao {
		list, isExistList := listMap[itemDao.ListID]
		if !isExistList {
			logger.Errorf(ctx, "not found list %d for item %d", list.ID, itemDao.ID)
			continue
		}

		item := convertItemFromDao(itemDao)
		list.Items = append(list.Items, item)
		listMap[itemDao.ListID] = list
	}

	lists := convert.MapValuesToSlice(listMap)
	return lists, nil
}

func convertListFromDao(listDao list_dao.List) server.List {
	return server.List{
		ID:    listDao.ID,
		Name:  listDao.Name,
		Items: nil,
	}
}
