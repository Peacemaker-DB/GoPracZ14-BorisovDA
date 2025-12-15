package httpx

import (
	"example.com/goprac11-borisovda/internal/http/handlers"
	"github.com/go-chi/chi/v5"
)

func NewRouter(h *handlers.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/api/v1/notes", h.GetAllNotes)
	r.Get("/api/v1/notes/keyset", h.GetNotesKeyset)
	r.Post("/api/v1/notes/batch", h.GetNotesBatch)
	r.Post("/api/v1/notes", h.CreateNote)
	r.Get("/api/v1/notes/{id}", h.GetNote)
	r.Patch("/api/v1/notes/{id}", h.UpdateNote)
	r.Delete("/api/v1/notes/{id}", h.DeleteNote)

	return r
}
