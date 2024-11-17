package notification

import (
	message "Notify-storage-service/internal/handler/model/msg"
	notify "Notify-storage-service/internal/model/notification"
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
func (r repository) GetById(ctx context.Context, id int) (notify.Notification, error) {
	return notify.Notification{}, nil
}

func (r repository) GetOld(ctx context.Context, userID int) ([]message.Notify, error) {
	query := `SELECT user_id, notification, created_at FROM notify WHERE user_id = $1 AND status = $2`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return []message.Notify{}, fmt.Errorf("failed to prepare query: %w", err)
	}

	rows, err := stmt.QueryContext(ctx, userID, OldStatus)
	if err != nil {
		return []message.Notify{}, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var notifications []message.Notify
	for rows.Next() {
		var existing message.Notify
		err := rows.Scan(
			&existing.UserId,
			&existing.Data,
			&existing.CreatedAt,
		)
		if err != nil {
			return []message.Notify{}, fmt.Errorf("failed to scan row: %w", err)
		}
		notifications = append(notifications, existing)
	}

	return notifications, nil
}

func (r repository) GetCurrent(ctx context.Context, userID int) ([]message.Notify, error) {
	query := `SELECT user_id, notification, created_at, expired_at FROM notify WHERE user_id = $1 AND status = $2`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return []message.Notify{}, fmt.Errorf("failed to prepare query: %w", err)
	}

	rows, err := stmt.QueryContext(ctx, userID, CurrentStatus)
	if err != nil {
		return []message.Notify{}, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var notifications []message.Notify
	for rows.Next() {
		var existing message.Notify
		err := rows.Scan(
			&existing.UserId,
			&existing.Data,
			&existing.CreatedAt,
			&existing.ExpiredAt,
		)
		if err != nil {
			return []message.Notify{}, fmt.Errorf("failed to scan row: %w", err)
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
