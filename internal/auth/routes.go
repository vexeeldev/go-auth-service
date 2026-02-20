package auth

import "github.com/go-chi/chi/v5"

func Routes(r chi.Router, h *Handler) {
	r.Get("/users", h.GetUsers)
	r.Post("/users", h.CreateUser)
	r.Get("/users/{id}", h.GetUser)
	r.Delete("/users/{id}", h.DeleteUser)
	r.Put("/users/{id}", h.UpdateUser)
}
