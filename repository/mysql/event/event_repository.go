package event

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/ryanbaskara/learning-go/entity"
)

type EventRepository struct {
	db *sqlx.DB
}

type Event struct {
	ID        int64     `db:"id"`
	UserID    string    `db:"user_id"`
	Type      string    `db:"type"`
	Metadata  []byte    `db:"metadata"`
	CreatedAt time.Time `db:"created_at"`
}

func NewEventRepository(db *sqlx.DB) *EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (r *EventRepository) ListEventByUserID(ctx context.Context, userID string) ([]*entity.Event, error) {
	q := "SELECT * FROM events WHERE user_id = ? ORDER BY id DESC LIMIT 100"
	q = r.db.Rebind(q)

	var events []*Event
	if err := r.db.SelectContext(ctx, &events, q, userID); err != nil {
		return nil, err
	}

	eventsEntity := make([]*entity.Event, len(events))
	for i, e := range events {
		var metadata map[string]interface{}
		json.Unmarshal(e.Metadata, &metadata)

		eventsEntity[i] = &entity.Event{
			ID:        e.ID,
			UserID:    e.UserID,
			Type:      e.Type,
			MetaData:  metadata,
			CreatedAt: e.CreatedAt,
		}
	}

	return eventsEntity, nil
}

func (r *EventRepository) CreateEvent(ctx context.Context, event *entity.Event) error {
	q := "INSERT INTO events (user_id, type, metadata, created_at) "
	q += "VALUES (:user_id, :type, :metadata, :created_at)"

	jsonBytes, err := json.Marshal(event.MetaData)
	if err != nil {
		return err
	}

	createEvent := &Event{
		UserID:    event.UserID,
		Type:      event.Type,
		Metadata:  jsonBytes,
		CreatedAt: event.CreatedAt,
	}

	res, err := r.db.NamedExecContext(ctx, q, createEvent)
	if err != nil {
		return err
	}

	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return err
	}

	event.ID = lastInsertId

	return nil
}
