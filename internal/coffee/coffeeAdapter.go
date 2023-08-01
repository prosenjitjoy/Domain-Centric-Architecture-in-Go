package coffee

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type adapter struct {
	db *pgx.Conn
}

func NewAdapter(db *pgx.Conn) Adapter {
	return &adapter{db: db}
}

func (a *adapter) GetAllCoffees(ctx context.Context) ([]Coffee, error) {
	rows, err := a.db.Query(ctx, "SELECT * FROM coffees")
	if err != nil {
		return nil, err
	}

	coffees, err := pgx.CollectRows(rows, pgx.RowToStructByName[Coffee])
	if err != nil {
		fmt.Println("here")
		return nil, err
	}

	return coffees, nil
}

func (a *adapter) GetCoffeeByID(ctx context.Context, id uuid.UUID) (*Coffee, error) {
	rows, err := a.db.Query(ctx, "SELECT * FROM coffees WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	coffee, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Coffee])
	if err != nil {
		return nil, err
	}

	return &coffee, nil
}

func (a *adapter) CreateCoffee(ctx context.Context, coffee *Coffee) (*Coffee, error) {
	query := "INSERT INTO coffees (name, roast, region, image, price, grind_unit) values($1, $2, $3, $4, $5, $6) returning id"

	err := a.db.QueryRow(ctx, query, coffee.Name, coffee.Roast, coffee.Region, coffee.Image, coffee.Price, coffee.GrindUnit).Scan(&coffee.ID)
	if err != nil {
		return nil, err
	}

	return coffee, nil
}

func (a *adapter) UpdateCoffeeByID(ctx context.Context, coffee *Coffee) (*Coffee, error) {
	query := "UPDATE coffees SET name=$1, roast=$2, region=$3, image=$4, price=$5, grind_unit=$6, updated_at=$7 WHERE id=$8 returning *"

	rows, err := a.db.Query(ctx, query, coffee.Name, coffee.Roast, coffee.Region, coffee.Image, coffee.Price, coffee.GrindUnit, time.Now(), coffee.ID)
	if err != nil {
		return nil, err
	}

	response, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Coffee])
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (a *adapter) DeleteCoffeeByID(ctx context.Context, id uuid.UUID) (*Coffee, error) {
	rows, err := a.db.Query(ctx, "DELETE FROM coffees WHERE id=$1 returning *", id)
	if err != nil {
		return nil, err
	}

	coffee, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Coffee])
	if err != nil {
		return nil, err
	}

	return &coffee, nil
}
