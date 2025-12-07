package business

import (
	"context"
	"testing"
	"time"

	"l2.18/internal/domain"
	"l2.18/internal/storage/memory"
)

func TestCalendarService_CreateEvent(t *testing.T) {
	repo := memory.NewStorage()
	service := NewCalendarService(repo)

	tests := []struct {
		name    string
		event   *domain.Event
		wantErr bool
	}{
		{
			name: "Success - Normal Event",
			event: &domain.Event{
				ID:    "1",
				Title: "Meeting",
				Date:  time.Now(),
			},
			wantErr: false,
		},
		{
			name: "Fail - Duplicate ID",
			event: &domain.Event{
				ID:     "1",
				Title: "Party",
				Date:   time.Now(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.CreateEvent(context.Background(), tt.event)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
