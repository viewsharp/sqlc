// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
	"iter"
)

const iterValues = `-- name: IterValues :iter
SELECT a, b
FROM myschema.foo
WHERE b = $1
`

func (q *Queries) IterValues(ctx context.Context, b sql.NullInt32) IterValuesRows {
	rows, err := q.db.QueryContext(ctx, iterValues, b)
	if err != nil {
		return IterValuesRows{err: err}
	}
	return IterValuesRows{rows: rows}
}

type IterValuesRows struct {
	rows *sql.Rows
	err  error
}

func (r *IterValuesRows) Iterate() iter.Seq[MyschemaFoo] {
	if r.rows == nil {
		return func(yield func(MyschemaFoo) bool) {}
	}

	return func(yield func(MyschemaFoo) bool) {
		defer r.rows.Close()

		for r.rows.Next() {
			var i MyschemaFoo
			err := r.rows.Scan(&i.A, &i.B)
			if err != nil {
				r.err = err
				return
			}

			if !yield(i) {
				r.err = r.rows.Close()
				return
			}
		}
	}
}

func (r *IterValuesRows) Close() error {
	return r.rows.Close()
}

func (r *IterValuesRows) Err() error {
	if r.err != nil {
		return r.err
	}
	return r.rows.Err()
}
