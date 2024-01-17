package list

import (
	"context"
	database "example-api/internal/database/dao"
	"example-api/pkg/pg"
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

func (d Dao) List(ctx context.Context, q pg.Queryable, filter ListFilter) ([]List, error) {
	qb := pg.PSQL.
		Select(pg.AllFields).
		From(database.ListTable).
		Where(sq.Eq{
			database.ListIDField: filter.IDs,
		})

	var list []List
	if err := q.Select(ctx, &list, qb); err != nil {
		return nil, fmt.Errorf("list object list from db: %w", err)
	}

	return list, nil
}
