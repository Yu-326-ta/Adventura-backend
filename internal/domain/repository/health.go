package repository

import (
	"Adventura/internal/domain/model"
	"context"
)

type Health interface {
	Status(ctx context.Context) (*model.Health, error)
}
