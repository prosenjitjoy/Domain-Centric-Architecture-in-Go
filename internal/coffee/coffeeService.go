package coffee

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type service struct {
	adapter Adapter
	timeout time.Duration
}

func NewService(adapter Adapter) Service {
	return &service{
		adapter: adapter,
		timeout: 2 * time.Second,
	}
}

func (s *service) GetAllCoffees(c context.Context) ([]Coffee, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()
	return s.adapter.GetAllCoffees(ctx)
}

func (s *service) GetCoffeeByID(c context.Context, id uuid.UUID) (*Coffee, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()
	return s.adapter.GetCoffeeByID(ctx, id)
}

func (s *service) CreateCoffee(c context.Context, coffee *Coffee) (*Coffee, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()
	return s.adapter.CreateCoffee(ctx, coffee)
}

func (s *service) UpdateCoffeeByID(c context.Context, coffee *Coffee) (*Coffee, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()
	return s.adapter.UpdateCoffeeByID(ctx, coffee)
}

func (s *service) DeleteCoffeeByID(c context.Context, id uuid.UUID) (*Coffee, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()
	return s.adapter.DeleteCoffeeByID(ctx, id)
}
