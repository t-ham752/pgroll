package migrations

import (
	"context"
	"database/sql"
	"fmt"

	"pg-roll/pkg/schema"
)

type OpDropIndex struct {
	Name string `json:"name"`
}

var _ Operation = (*OpDropIndex)(nil)

func (o *OpDropIndex) Start(ctx context.Context, conn *sql.DB, schemaName string, stateSchema string, s *schema.Schema) error {
	// no-op
	return nil
}

func (o *OpDropIndex) Complete(ctx context.Context, conn *sql.DB) error {
	// drop the index concurrently
	_, err := conn.ExecContext(ctx, fmt.Sprintf("DROP INDEX CONCURRENTLY IF EXISTS %s", o.Name))

	return err
}

func (o *OpDropIndex) Rollback(ctx context.Context, conn *sql.DB) error {
	// no-op
	return nil
}

func (o *OpDropIndex) Validate(ctx context.Context, s *schema.Schema) error {
	for _, table := range s.Tables {
		_, ok := table.Indexes[o.Name]
		if ok {
			return nil
		}
	}
	return IndexDoesNotExistError{Name: o.Name}
}
