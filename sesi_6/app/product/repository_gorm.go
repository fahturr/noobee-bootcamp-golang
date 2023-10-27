package product

import (
	"context"

	"gorm.io/gorm"
)

type PostgresGormRepository struct {
	db *gorm.DB
}

// DeleteById implements Repository
func (p PostgresGormRepository) DeleteById(ctx context.Context, id int) (err error) {
	return p.db.Delete(&Product{}, id).Error
}

// UpdateById implements Repository
func (p PostgresGormRepository) UpdateById(ctx context.Context, model Product) (err error) {
	return p.db.Save(model).Error
}

// GetAll implements Repository
func (p PostgresGormRepository) GetAll(ctx context.Context) ([]Product, error) {
	var products []Product
	err := p.db.Find(&products).Error
	return products, err
}

func (p PostgresGormRepository) Create(ctx context.Context, model Product) (err error) {
	return p.db.Create(&model).Error
}

func NewPostgresGormRepository(db *gorm.DB) PostgresGormRepository {
	return PostgresGormRepository{
		db: db,
	}
}
