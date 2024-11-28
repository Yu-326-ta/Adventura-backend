package dao

import (
	"Adventura/internal/domain/model"
	"Adventura/internal/domain/repository"
	"context"
)

type (
	health struct{}
)

func NewHealth() repository.Health {
	return &health{}
}

func (h *health) Status(ctx context.Context) (*model.Health, error) {
	status := model.Health{
		Status: "ok",
	}

	return &status, nil
}
