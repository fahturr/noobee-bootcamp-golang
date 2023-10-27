package product

import (
	"context"
	"log"
)

type Repository interface {
	Create(ctx context.Context, req Product) (err error)
	GetAll(ctx context.Context) ([]Product, error)
	UpdateById(ctx context.Context, req Product) (err error)
	DeleteById(ctx context.Context, id int) (err error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) CreateProduct(ctx context.Context, req Product) (err error) {
	if err = req.Validate(); err != nil {
		log.Println("error when try to validate request with error :", err.Error())
		return
	}

	if err = s.repo.Create(ctx, req); err != nil {
		log.Println("error when try to Create to database with error :", err.Error())
		return

	}

	return
}

func (s Service) GetAllProduct(ctx context.Context) (products []Product, err error) {
	products, err = s.repo.GetAll(ctx)

	return
}

func (s Service) UpdateProductById(ctx context.Context, req Product) (err error) {
	if err = s.repo.UpdateById(ctx, req); err != nil {
		log.Println("error when try to Update to database with error :", err.Error())
		return

	}

	return
}

func (s Service) DeleteProductById(ctx context.Context, id int) (err error) {
	if err = s.repo.DeleteById(ctx, id); err != nil {
		log.Println("error when try to Update to database with error :", err.Error())
		return

	}

	return
}
