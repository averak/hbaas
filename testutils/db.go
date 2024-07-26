package testutils

import (
	"context"
	"reflect"
	"testing"

	"github.com/averak/hbaas/app/adapter/dao"
	"github.com/averak/hbaas/app/domain/repository/transaction"
	"github.com/averak/hbaas/app/infrastructure/db"
	"github.com/huandu/go-sqlbuilder"
)

var tableNames []string

func init() {
	val := reflect.ValueOf(dao.TableNames)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		tableNames = append(tableNames, field.Interface().(string)) //nolint:forcetypeassert
	}
}

func MustDBConn(t *testing.T) transaction.Connection {
	t.Helper()

	conn, err := db.NewConnection()
	if err != nil {
		t.Fatal(err)
	}
	return conn
}

func Teardown(t *testing.T) {
	t.Helper()

	conn := MustDBConn(t)
	err := conn.BeginRwTransaction(context.Background(), func(ctx context.Context, tx transaction.Transaction) error {
		for _, table := range tableNames {
			query, args := sqlbuilder.NewDeleteBuilder().DeleteFrom(table).Build()
			_, err := tx.ExecContext(ctx, query, args...)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("failed to teardown: %v", err)
	}
}
