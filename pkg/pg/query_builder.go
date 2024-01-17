package pg

import (
	sq "github.com/Masterminds/squirrel"
)

// PSQL предпочтительный формат для postgres
var PSQL = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
