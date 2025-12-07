package business

import (
	"context"
	"errors"
	"fmt"
	"time"

	"l2.18/internal/domain"
)

var ErrEventTitleRequired = errors.New("event title can't be empty")

type CalendarService struct {
	storage domain.Storage
}

func NewCalendarService(storage domain.Storage) *CalendarService {
	return &CalendarService{storage: storage}
}

func (s *CalendarService) CreateEvent(ctx context.Context, event *domain.Event) error {
	if event.Title == "" {
		return ErrEventTitleRequired
	}
	return s.storage.CreateEvent(ctx, event)
}

func (s *CalendarService) UpdateEvent(ctx context.Context, event *domain.Event) error {
	if event.Title == "" {
		return ErrEventTitleRequired
	}

	if event.ID == "" {
		return fmt.Errorf("event ID is required for update")
	}

	return s.storage.UpdateEvent(ctx, event)
}

func (s *CalendarService) DeleteEvent(ctx context.Context, eventID string) error {
	if eventID == "" {
		return fmt.Errorf("event ID is required for delete")
	}

	return s.storage.DeleteEvent(ctx, eventID)
}

func (s *CalendarService) GetEventsForDay(ctx context.Context, date time.Time) ([]domain.Event, error) {
	result, err := s.storage.GetEventsForDay(ctx, date)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *CalendarService) GetEventsForWeek(ctx context.Context, date time.Time) ([]domain.Event, error) {
	return s.storage.GetEventsForWeek(ctx, date)
}

func (s *CalendarService) GetEventsForMonth(ctx context.Context, date time.Time) ([]domain.Event, error) {
	return s.storage.GetEventsForMonth(ctx, date)
}
