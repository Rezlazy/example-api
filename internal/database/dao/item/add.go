package item

import (
	"context"
	database "example-api/internal/database/dao"
	"example-api/pkg/pg"
	"fmt"
)

func (d Dao) Add(ctx context.Context, q pg.Queryable, createInputs ...CreateInput) ([]Item, error) {
	sqlizers := make([]pg.Sqlizer, 0, len(createInputs))
	for _, createInput := range createInputs {
		sqlizer := pg.PSQL.
			Insert(database.ItemTable).
			Columns(database.ItemTitleField, database.ItemDescriptionField, database.ItemListIDField).
			Values(createInput.Title, createInput.Description, createInput.ListID).
			SuffixExpr(pg.Returning())
		sqlizers = append(sqlizers, sqlizer)
	}

	var items []Item
	if err := q.QueryBatch(ctx, &items, sqlizers); err != nil {
		return nil, fmt.Errorf("query batch: %w", err)
	}

	return items, nil
}
