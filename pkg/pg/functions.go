package pg

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"strings"
)

const (
	AllFields = "*"
	MockField = "1"
)

func SelectExists(qb sq.SelectBuilder) Sqlizer {
	return qb.Prefix("SELECT EXISTS (").Suffix(")")
}

func Returning(fields ...string) Sqlizer {
	fieldsExpr := "*"
	if len(fields) > 0 {
		fieldsExpr = strings.Join(fields, ", ")
	}

	return sq.Expr(fmt.Sprintf("RETURNING %s", fieldsExpr))
}
