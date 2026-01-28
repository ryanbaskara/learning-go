package event

import (
	"context"
	"time"

	"github.com/ryanbaskara/learning-go/entity"
)

type EventRepository interface {
	ListEventByUserID(ctx context.Context, userID string) ([]*entity.Event, error)
	CreateEvent(ctx context.Context, event *entity.Event) error
}

type EventUsecase struct {
	eventRepository EventRepository
}

func NewEventUsecase(
	eventRepository EventRepository,
) *EventUsecase {
	return &EventUsecase{
		eventRepository: eventRepository,
	}
}

func (u *EventUsecase) CreateEvent(ctx context.Context, req *entity.CreateEventRequest) (*entity.Event, error) {

	event := &entity.Event{
		UserID:    req.UserID,
		Type:      req.Type,
		MetaData:  req.Metadata,
		CreatedAt: time.Now(),
	}

	err := u.eventRepository.CreateEvent(ctx, event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (u *EventUsecase) ListEventByUserID(ctx context.Context, req *entity.ListEventRequest) ([]*entity.Event, error) {
	events, err := u.eventRepository.ListEventByUserID(ctx, req.UserID)

	return events, err
}
