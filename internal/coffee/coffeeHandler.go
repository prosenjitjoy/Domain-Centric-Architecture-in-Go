package coffee

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var validate = validator.New()

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) GetAllCoffees(w http.ResponseWriter, r *http.Request) {
	coffees, err := h.service.GetAllCoffees(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(coffees)
}

func (h *handler) GetCoffeeByID(w http.ResponseWriter, r *http.Request) {
	coffeeID := chi.URLParam(r, "uuid")
	coffeeUUID, err := uuid.Parse(coffeeID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		return
	}

	coffee, err := h.service.GetCoffeeByID(context.Background(), coffeeUUID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*coffee)
}

func (h *handler) CreateCoffee(w http.ResponseWriter, r *http.Request) {
	var coffee Coffee
	err := json.NewDecoder(r.Body).Decode(&coffee)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		return
	}

	err = validate.Struct(coffee)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		return
	}

	response, err := h.service.CreateCoffee(context.Background(), &coffee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *handler) UpdateCoffeeByID(w http.ResponseWriter, r *http.Request) {
	coffeeID := chi.URLParam(r, "uuid")
	coffeeUUID, err := uuid.Parse(coffeeID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		return
	}

	var coffee Coffee
	err = json.NewDecoder(r.Body).Decode(&coffee)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		return
	}

	err = validate.Struct(coffee)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		return
	}

	coffee.ID = coffeeUUID

	response, err := h.service.UpdateCoffeeByID(context.Background(), &coffee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *handler) DeleteCoffeeByID(w http.ResponseWriter, r *http.Request) {
	coffeeID := chi.URLParam(r, "uuid")
	coffeeUUID, err := uuid.Parse(coffeeID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		return
	}

	coffee, err := h.service.DeleteCoffeeByID(context.Background(), coffeeUUID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*coffee)
}
