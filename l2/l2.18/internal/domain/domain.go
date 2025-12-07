package domain

import (
	"context"
	"time"
)

type Event struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Date        time.Time `json:"date"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

type Storage interface {
	CreateEvent(ctx context.Context, event *Event) error
	UpdateEvent(ctx context.Context, event *Event) error
	DeleteEvent(ctx context.Context, eventID string) error
	GetEventsForDay(ctx context.Context, date time.Time) ([]Event, error)
	GetEventsForWeek(ctx context.Context, date time.Time) ([]Event, error)
	GetEventsForMonth(ctx context.Context, date time.Time) ([]Event, error)
}
