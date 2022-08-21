package dbr

import "github.com/gocraft/dbr/v2"

func SprintSelect(stmt *dbr.SelectStmt) string {
	q := dbr.NewBuffer()
	err := stmt.Build(stmt.Dialect, q)
	if err != nil {
		return ""
	}

	whereVals := getValues(stmt.WhereCond, stmt.Dialect)

	s, err := dbr.InterpolateForDialect(q.String(), whereVals, stmt.Dialect)
	if err != nil {
		return ""
	}

	return s
}

func getValues(whereCond []dbr.Builder, dialect dbr.Dialect) []interface{} {
	vals := make([]interface{}, 0)

	for _, v := range whereCond {
		b := dbr.NewBuffer()
		err := v.Build(dialect, b)
		if err != nil {
			return vals
		}

		vals = append(vals, b.Value()...)
	}

	return vals
}
