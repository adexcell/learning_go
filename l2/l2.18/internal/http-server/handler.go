package httpserver

import (
	"encoding/json"
	"net/http"
	"time"

	"l2.18/internal/business"
	"l2.18/internal/domain"
)

type Handler struct {
	service *business.CalendarService
}

func NewHandler(service *business.CalendarService) *Handler {
	return &Handler{service: service}
}

type CreateEventRequest struct {
	UserID string `json:"user_id"`
	Date   string `json:"date"`
	Title  string `json:"title"`
}

type UpdateEventRequest struct {
	ID string `json:"id"`
	CreateEventRequest
}

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	// 1. Check Method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2. Decode JSON Body
	var req CreateEventRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// 3. Validate & Parse Date
	parsedDate, err := time.Parse(time.DateOnly, req.Date)
	if err != nil {
		http.Error(w, "Invalid date format. Use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	// 4. Map DTO to Domain Model
	event := &domain.Event{
		ID:     req.UserID + "_" + req.Title,
		UserID: req.UserID,
		Title:  req.Title,
		Date:   parsedDate,
	}

	// 5. Call Business Logic
	if err := h.service.CreateEvent(r.Context(), event); err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	// 6. Send Success Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Write a simple JSON success message
	w.Write([]byte(`{"result": "event created"}`))
}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req UpdateEventRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	parsedDate, err := time.Parse(time.DateOnly, req.Date)
	if err != nil {
		http.Error(w, "Invalid date format. Use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	event := &domain.Event{
		ID:     req.ID,
		UserID: req.UserID,
		Title:  req.Title,
		Date:   parsedDate,
	}

	if err := h.service.UpdateEvent(r.Context(), event); err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(`{"result":"event updated"}`))
}

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	eventID := r.URL.Query().Get("event_id")
	
	if eventID == "" {
		http.Error(w, "missing event_id", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteEvent(r.Context(), eventID); err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("user_id")
	dateStr := r.URL.Query().Get("date")

	if userID == "" || dateStr == "" {
		http.Error(w, "missing user_id or date", http.StatusBadRequest)
		return
	}

	parsedDate, err := time.Parse(time.DateOnly, dateStr)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	events, err := h.service.GetEventsForDay(r.Context(), parsedDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"result": events,
	}); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *Handler) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("user_id")
	dateStr := r.URL.Query().Get("date")

	if userID == "" || dateStr == "" {
		http.Error(w, "missing user_id or date", http.StatusBadRequest)
		return
	}

	parsedDate, err := time.Parse(time.DateOnly, dateStr)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	events, err := h.service.GetEventsForWeek(r.Context(), parsedDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"result": events,
	}); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *Handler) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("user_id")
	dateStr := r.URL.Query().Get("date")

	if userID == "" || dateStr == "" {
		http.Error(w, "missing user_id or date", http.StatusBadRequest)
		return
	}

	parsedDate, err := time.Parse(time.DateOnly, dateStr)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	events, err := h.service.GetEventsForMonth(r.Context(), parsedDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"result": events,
	}); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
