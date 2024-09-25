package usecase

import (
	"context"

	"github.com/zinrai/alert-hub-go/internal/domain"
	"github.com/zinrai/alert-hub-go/internal/repository"
)

type AlertUsecase interface {
	GetAlerts(ctx context.Context) ([]domain.Alert, error)
	CreateAlert(ctx context.Context, alert domain.Alert) error
	GetAlert(ctx context.Context, id int) (domain.Alert, error)
	UpdateAlert(ctx context.Context, alert domain.Alert) error
}

type alertUsecase struct {
	repo repository.AlertRepository
}

func NewAlertUsecase(repo repository.AlertRepository) AlertUsecase {
	return &alertUsecase{repo}
}

func (u *alertUsecase) GetAlerts(ctx context.Context) ([]domain.Alert, error) {
	return u.repo.GetAlerts(ctx)
}

func (u *alertUsecase) CreateAlert(ctx context.Context, alert domain.Alert) error {
	return u.repo.CreateAlert(ctx, alert)
}

func (u *alertUsecase) GetAlert(ctx context.Context, id int) (domain.Alert, error) {
	return u.repo.GetAlert(ctx, id)
}

func (u *alertUsecase) UpdateAlert(ctx context.Context, alert domain.Alert) error {
	return u.repo.UpdateAlert(ctx, alert)
}
