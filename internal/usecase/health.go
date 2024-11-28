package usecase

import (
	"Adventura/internal/domain/repository"
	"context"
)

type Health interface {
	Health(ctx context.Context) (*HealthResponse, error)
}

type health struct {
	HealthRepo repository.Health
}

type HealthResponse struct {
	Status string
}

func NewHealth(healthtRepo repository.Health) *health {
	return &health{
		HealthRepo: healthtRepo,
	}
}

func (h *health) Health(ctx context.Context) (*HealthResponse, error) {
	healthStatus, err := h.HealthRepo.Status(ctx)
	if err != nil {
		return nil, err
	}

	return &HealthResponse{
		Status: healthStatus.Status,
	}, nil
}
