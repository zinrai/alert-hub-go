package repository

import (
	"context"

	"github.com/zinrai/alert-hub-go/internal/domain"
)

type AlertRepository interface {
	GetAlerts(ctx context.Context) ([]domain.Alert, error)
	CreateAlert(ctx context.Context, alert domain.Alert) error
	GetAlert(ctx context.Context, id int) (domain.Alert, error)
	UpdateAlert(ctx context.Context, alert domain.Alert) error
}
