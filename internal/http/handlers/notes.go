package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"example.com/goprac11-borisovda/internal/core"
	"example.com/goprac11-borisovda/internal/repo"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Repo repo.NoteRepository
}

func writeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(v)
}

func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var n core.Note
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	id, err := h.Repo.Create(r.Context(), n)
	if err != nil {
		http.Error(w, "Failed to create note", http.StatusInternalServerError)
		return
	}
	n.ID = id
	writeJSON(w, http.StatusCreated, n)
}

func (h *Handler) GetAllNotes(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	limit := 10
	offset := 0
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}
	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
			offset = o
		}
	}
	notes, err := h.Repo.GetAll(r.Context(), limit, offset)
	if err != nil {
		http.Error(w, "Failed to fetch notes", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, notes)
}

func (h *Handler) GetNotesKeyset(w http.ResponseWriter, r *http.Request) {
	lastCreatedAt := r.URL.Query().Get("last_created_at")
	lastIDStr := r.URL.Query().Get("last_id")
	limitStr := r.URL.Query().Get("limit")

	limit := 10
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	var lastID int64 = 0
	if lastIDStr != "" {
		if id, err := strconv.ParseInt(lastIDStr, 10, 64); err == nil {
			lastID = id
		}
	}

	if lastCreatedAt == "" {
		lastCreatedAt = time.Now().Format(time.RFC3339)
	}

	notes, err := h.Repo.GetAllKeyset(r.Context(), lastCreatedAt, lastID, limit)
	if err != nil {
		http.Error(w, "Failed to fetch notes", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, notes)
}

func (h *Handler) GetNotesBatch(w http.ResponseWriter, r *http.Request) {
	var ids []int64
	if err := json.NewDecoder(r.Body).Decode(&ids); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if len(ids) == 0 {
		writeJSON(w, http.StatusOK, []core.Note{})
		return
	}

	notes, err := h.Repo.GetBatch(r.Context(), ids)
	if err != nil {
		http.Error(w, "Failed to fetch notes", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, notes)
}

func (h *Handler) GetNote(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	n, err := h.Repo.Get(r.Context(), id)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, n)
}

func (h *Handler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	var upd core.Note
	if err := json.NewDecoder(r.Body).Decode(&upd); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	n, err := h.Repo.Update(r.Context(), id, upd)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, n)
}

func (h *Handler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err := h.Repo.Delete(r.Context(), id); err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
