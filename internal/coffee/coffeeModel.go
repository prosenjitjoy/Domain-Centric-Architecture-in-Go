package coffee

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Coffee struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Roast     string    `json:"roast" validate:"required"`
	Region    string    `json:"region" validate:"required"`
	Image     string    `json:"image" validate:"required"`
	Price     float32   `json:"price" validate:"required"`
	GrindUnit int32     `json:"grind_unit" db:"grind_unit" validate:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Adapter interface {
	GetAllCoffees(ctx context.Context) ([]Coffee, error)
	GetCoffeeByID(ctx context.Context, id uuid.UUID) (*Coffee, error)
	CreateCoffee(ctx context.Context, coffee *Coffee) (*Coffee, error)
	UpdateCoffeeByID(ctx context.Context, coffee *Coffee) (*Coffee, error)
	DeleteCoffeeByID(ctx context.Context, id uuid.UUID) (*Coffee, error)
}

type Service interface {
	GetAllCoffees(c context.Context) ([]Coffee, error)
	GetCoffeeByID(c context.Context, id uuid.UUID) (*Coffee, error)
	CreateCoffee(c context.Context, coffee *Coffee) (*Coffee, error)
	UpdateCoffeeByID(c context.Context, coffee *Coffee) (*Coffee, error)
	DeleteCoffeeByID(c context.Context, id uuid.UUID) (*Coffee, error)
}
