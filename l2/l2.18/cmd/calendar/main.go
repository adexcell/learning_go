package main

import (
	"log"
	"net/http"
	"time"

	"l2.18/internal/business"
	"l2.18/internal/config"
	httpserver "l2.18/internal/http-server"
	"l2.18/internal/storage/memory"
)

func main() {
	cfg := config.Load()

	storage := memory.NewStorage()
	calendar := business.NewCalendarService(storage)
	handler := httpserver.NewHandler(calendar)

	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", handler.CreateEvent)
	mux.HandleFunc("/update_event", handler.UpdateEvent)
	mux.HandleFunc("/delete_event", handler.DeleteEvent)
	mux.HandleFunc("/events_for_day", handler.GetEventsForDay)
	mux.HandleFunc("/events_for_week", handler.GetEventsForWeek)
	mux.HandleFunc("/events_for_month", handler.GetEventsForMonth)

	wrappedMux := LoggerMidleware(mux)

	log.Printf("Calendar app running on port %s...", cfg.Port)
	if err := http.ListenAndServe(cfg.Port, wrappedMux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func LoggerMidleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		log.Printf("[%s] %s - %v", r.Method, r.URL.Path, duration)
	})
}
