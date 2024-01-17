package item

import (
	"context"
	database "example-api/internal/database/dao"
	"example-api/pkg/pg"
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

func (d Dao) List(ctx context.Context, q pg.Queryable, filter ListFilter) ([]Item, error) {
	qb := pg.PSQL.
		Select(pg.AllFields).
		From(database.ItemTable).
		Where(sq.Eq{
			database.ItemListIDField: filter.ListIDs,
		})

	var items []Item
	if err := q.Select(ctx, &items, qb); err != nil {
		return nil, fmt.Errorf("list items from db: %w", err)
	}

	return items, nil
}
