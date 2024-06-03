package service

import (
	"context"

	"github.com/Penomatikus/onionarch/internal/domain/model"
)

type PlayerSerivce interface {
	Create(ctx context.Context, player model.Player) error
	Login(ctx context.Context) (b []byte, err error)
}
