package list

import (
	"context"
	database "example-api/internal/database/dao"
	"example-api/pkg/pg"
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

func (d Dao) CheckExists(ctx context.Context, q pg.Queryable, id int64) (bool, error) {
	qb := pg.PSQL.
		Select(pg.MockField).
		From(database.ListTable).
		Where(sq.Eq{
			database.ListIDField: id,
		})

	existsQuery := pg.SelectExists(qb)

	var isExists bool
	if err := q.Get(ctx, &isExists, existsQuery); err != nil {
		return false, fmt.Errorf("get is exists item: %w", err)
	}

	return isExists, nil
}
