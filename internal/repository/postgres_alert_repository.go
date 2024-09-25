package repository

import (
	"context"
	"database/sql"
	"os"

	"github.com/zinrai/alert-hub-go/internal/domain"
)

type postgresAlertRepository struct {
	db  *sql.DB
	key string
}

func NewPostgresAlertRepository(db *sql.DB) (AlertRepository, error) {
	key := os.Getenv("ENCRYPTION_KEY")
	return &postgresAlertRepository{db: db, key: key}, nil
}

func (r *postgresAlertRepository) GetAlerts(ctx context.Context) ([]domain.Alert, error) {
	query := `SELECT id, subject, decrypt_body(body, $1) as body, identifier, urgency, status, created_at, updated_at FROM alerts`

	rows, err := r.db.QueryContext(ctx, query, r.key)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alerts []domain.Alert
	for rows.Next() {
		var a domain.Alert
		err := rows.Scan(&a.ID, &a.Subject, &a.Body, &a.Identifier, &a.Urgency, &a.Status, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			return nil, err
		}
		alerts = append(alerts, a)
	}

	return alerts, nil
}

func (r *postgresAlertRepository) CreateAlert(ctx context.Context, alert domain.Alert) error {
	query := `INSERT INTO alerts (subject, body, identifier, urgency, status) VALUES ($1, encrypt_body($2, $3), $4, $5, $6)`
	_, err := r.db.ExecContext(ctx, query, alert.Subject, alert.Body, r.key, alert.Identifier, alert.Urgency, alert.Status)
	return err
}

func (r *postgresAlertRepository) GetAlert(ctx context.Context, id int) (domain.Alert, error) {
	query := `SELECT id, subject, decrypt_body(body, $1) as body, identifier, urgency, status, created_at, updated_at FROM alerts WHERE id = $2`

	var alert domain.Alert
	err := r.db.QueryRowContext(ctx, query, r.key, id).Scan(
		&alert.ID, &alert.Subject, &alert.Body, &alert.Identifier, &alert.Urgency, &alert.Status,
		&alert.CreatedAt, &alert.UpdatedAt)

	return alert, err
}

func (r *postgresAlertRepository) UpdateAlert(ctx context.Context, alert domain.Alert) error {
	query := `
		UPDATE alerts
		SET
			subject = COALESCE($1, subject),
			body = COALESCE(encrypt_body($2, $3), body),
			identifier = COALESCE($4, identifier),
			urgency = COALESCE($5, urgency),
			status = COALESCE($6, status),
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $7
	`

	_, err := r.db.ExecContext(ctx, query,
		alert.Subject,
		alert.Body,
		r.key,
		alert.Identifier,
		alert.Urgency,
		alert.Status,
		alert.ID)

	return err
}
