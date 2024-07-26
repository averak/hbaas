package repository

import (
	"context"

	"github.com/averak/hbaas/app/domain/model"
	"github.com/averak/hbaas/app/domain/repository/transaction"
)

type EchoRepository interface {
	Save(ctx context.Context, tx transaction.Transaction, echos ...model.Echo) error
}
