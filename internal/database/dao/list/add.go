package list

import (
	"context"
	database "example-api/internal/database/dao"
	"example-api/pkg/pg"
	"fmt"
)

func (d Dao) Add(ctx context.Context, q pg.Queryable, createInput CreateInput) (List, error) {
	qb := pg.PSQL.
		Insert(database.ListTable).
		Columns(database.ListNameField).
		Values(createInput.Name).
		SuffixExpr(pg.Returning())

	var list List
	if err := q.Get(ctx, &list, qb); err != nil {
		return List{}, fmt.Errorf("add list to db: %w", err)
	}

	return list, nil
}
