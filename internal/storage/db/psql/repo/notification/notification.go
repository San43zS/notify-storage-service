package notification

import (
	notification2 "Notify-storage-service/internal/model/notification"
	"Notify-storage-service/internal/storage/api/notification"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) notification.Notification {
	return repository{
		db: db,
	}
}

// ?????????????????????????????????????????????????????
func (r repository) GetById(ctx context.Context, id int) (notification2.Notification, error) {
	return notification2.Notification{}, nil
}

func (r repository) Add(ctx context.Context, notification notification2.Notification) error {
	query := "INSERT INTO notify (user_id, notification, created_at) VALUES ($1, $2, $3)"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare query: %w", err)
	}

	_, err = stmt.ExecContext(ctx, notification.UserId, notification.Data, notification.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to exec query: %w", err)
	}
	return nil
}

func (r repository) GetOld(ctx context.Context, userID int) ([]notification2.Notification, error) {
	query := `SELECT user_id, notification, created_at FROM notify WHERE user_id = $1 AND status = $2`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return []notification2.Notification{}, fmt.Errorf("failed to prepare query: %w", err)
	}

	rows, err := stmt.QueryContext(ctx, userID, OldStatus)
	if err != nil {
		return []notification2.Notification{}, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var notifications []notification2.Notification
	for rows.Next() {
		var existing notification2.Notification
		err := rows.Scan(
			&existing.UserId,
			&existing.Data,
			&existing.CreatedAt,
		)
		if err != nil {
			return []notification2.Notification{}, fmt.Errorf("failed to scan row: %w", err)
		}
		notifications = append(notifications, existing)
	}

	return notifications, nil
}

func (r repository) GetCurrent(ctx context.Context, userID int) ([]notification2.Notification, error) {
	query := `SELECT user_id, notification, created_at, expired_at FROM notify WHERE user_id = $1 AND status = $2`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return []notification2.Notification{}, fmt.Errorf("failed to prepare query: %w", err)
	}

	rows, err := stmt.QueryContext(ctx, userID, CurrentStatus)
	if err != nil {
		return []notification2.Notification{}, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var notifications []notification2.Notification
	for rows.Next() {
		var existing notification2.Notification
		err := rows.Scan(
			&existing.UserId,
			&existing.Data,
			&existing.CreatedAt,
			&existing.ExpiredAt,
		)
		if err != nil {
			return []notification2.Notification{}, fmt.Errorf("failed to scan row: %w", err)
		}
		notifications = append(notifications, existing)
	}

	return notifications, nil
}

func (r repository) Delete(ctx context.Context, userID int, ids []int) error {
	query := `DELETE FROM notify WHERE user_id = $1 AND id IN (SELECT unnest($2))`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare query: %w", err)
	}

	_, err = stmt.ExecContext(ctx, userID, ids)
	if err != nil {
		return fmt.Errorf("failed to exec query: %w", err)
	}
	return nil
}
