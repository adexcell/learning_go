package memory

import (
	"context"
	"fmt"
	"sync"
	"time"

	"l2.18/internal/domain"
)

type Storage struct {
	mu     sync.RWMutex
	events map[string]domain.Event
}

func NewStorage() *Storage {
	return &Storage{
		events: make(map[string]domain.Event),
	}
}

func (s *Storage) CreateEvent(ctx context.Context, event *domain.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.events[event.ID]; ok {
		return fmt.Errorf("event already exists: %s", event.ID)
	}

	s.events[event.ID] = *event
	return nil
}

func (s *Storage) UpdateEvent(ctx context.Context, event *domain.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.events[event.ID]; !ok {
		return fmt.Errorf("event not found: %s", event.ID)
	}

	s.events[event.ID] = *event
	return nil
}

func (s *Storage) DeleteEvent(ctx context.Context, eventID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.events[eventID]; !ok {
		return fmt.Errorf("event not found: %s", eventID)
	}

	delete(s.events, eventID)
	return nil
}

func (s *Storage) GetEventsForDay(ctx context.Context, date time.Time) ([]domain.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []domain.Event
	for _, event := range s.events {
		if event.Date.Year() == date.Year() &&
			event.Date.Month() == date.Month() &&
			event.Date.Day() == date.Day() {
			result = append(result, event)
		}
	}
	return result, nil
}

func (s *Storage) GetEventsForWeek(ctx context.Context, date time.Time) ([]domain.Event, error) {
	var result []domain.Event
	for _, event := range s.events {
		eventYear, eventWeek := event.Date.ISOWeek()
		dateYear, dateWeek := date.ISOWeek()
		if eventYear == dateYear &&
			eventWeek == dateWeek {
			result = append(result, event)
		}
	}
	return result, nil
}

func (s *Storage) GetEventsForMonth(ctx context.Context, date time.Time) ([]domain.Event, error) {
	var result []domain.Event
	for _, event := range s.events {
		if event.Date.Year() == date.Year() &&
			event.Date.Month() == date.Month() {
			result = append(result, event)
		}
	}
	return result, nil
}
