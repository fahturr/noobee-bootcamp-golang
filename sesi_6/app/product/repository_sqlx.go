package product

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type PostgresSQLXRepository struct {
	db *sqlx.DB
}

// DeleteById implements Repository.
func (p PostgresSQLXRepository) DeleteById(ctx context.Context, id int) (err error) {
	query := `
		DELETE FROM products
		WHERE id = :id
	`

	stmt, err := p.db.PrepareNamed(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	return
}

// UpdateById implements Repository.
func (p PostgresSQLXRepository) UpdateById(ctx context.Context, req Product) (err error) {
	query := `
		UPDATE products
		SET name = :name, category = :category, price = :price, stock = :stock
		WHERE id = :id
	`

	stmt, err := p.db.PrepareNamed(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(req)

	return
}

// GetAll implements Repository.
func (p PostgresSQLXRepository) GetAll(ctx context.Context) ([]Product, error) {
	query := `
		SELECT id, name, category, price, stock
		FROM products
	`

	var products []Product

	err := p.db.Select(&products, query)
	if err != nil {
		return products, err
	}

	return products, nil
}

func (p PostgresSQLXRepository) Create(ctx context.Context, req Product) (err error) {
	query := `
		INSERT INTO products (
			name, category, price, stock
		) VALUES (
			:name, :category, :price, :stock
		)
	`

	stmt, err := p.db.PrepareNamed(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(req)

	return
}

func NewPostgresSQLXRepository(db *sqlx.DB) PostgresSQLXRepository {
	return PostgresSQLXRepository{
		db: db,
	}
}
