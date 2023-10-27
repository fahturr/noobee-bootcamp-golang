package product

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type HandlerInterface interface {
	CreateProduct(ctx context.Context, req CreateProductRequest) error
}

type Handler struct {
	svc Service
}

func NewHandler(svc Service) Handler {
	return Handler{
		svc: svc,
	}
}

func (h Handler) CreateProduct(c *fiber.Ctx) error {
	var req = CreateProductRequest{}

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	model := Product{
		Name:     req.Name,
		Category: req.Category,
		Price:    req.Price,
		Stock:    req.Stock,
	}

	err = h.svc.CreateProduct(c.UserContext(), model)
	if err != nil {
		var payload fiber.Map
		httpCode := 400

		switch err {
		case ErrEmptyName, ErrEmptyCategory, ErrEmptyPrice, ErrEmptyStock:
			payload = fiber.Map{
				"success": false,
				"message": "ERR BAD REQUEST",
				"error":   err.Error(),
			}
			httpCode = http.StatusBadRequest
		default:
			payload = fiber.Map{
				"success": false,
				"message": "ERR INTERNAL",
				"error":   err.Error(),
			}
			httpCode = http.StatusInternalServerError
		}
		return c.Status(httpCode).JSON(payload)
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "CREATE SUCCESS",
	})

}

func (h Handler) GetAllProduct(c *fiber.Ctx) error {
	products, err := h.svc.GetAllProduct(c.UserContext())
	if err != nil {
		var payload fiber.Map
		httpCode := 400

		switch err {
		case ErrEmptyName, ErrEmptyCategory, ErrEmptyPrice, ErrEmptyStock:
			payload = fiber.Map{
				"success": false,
				"message": "ERR BAD REQUEST",
				"error":   err.Error(),
			}
			httpCode = http.StatusBadRequest
		default:
			payload = fiber.Map{
				"success": false,
				"message": "ERR INTERNAL",
				"error":   err.Error(),
			}
			httpCode = http.StatusInternalServerError
		}
		return c.Status(httpCode).JSON(payload)
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "GET ALL SUCCESS",
		"data":    products,
	})

}

func (h Handler) UpdateProductById(c *fiber.Ctx) error {
	var req = UpdateProductByIdRequest{}

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	err = c.ParamsParser(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	model := Product{
		Id:       req.ID,
		Name:     req.Name,
		Category: req.Category,
		Price:    req.Price,
		Stock:    req.Stock,
	}

	err = h.svc.UpdateProductById(c.UserContext(), model)
	if err != nil {
		var payload fiber.Map
		httpCode := 400

		switch err {
		case ErrEmptyName, ErrEmptyCategory, ErrEmptyPrice, ErrEmptyStock:
			payload = fiber.Map{
				"success": false,
				"message": "ERR BAD REQUEST",
				"error":   err.Error(),
			}
			httpCode = http.StatusBadRequest
		default:
			payload = fiber.Map{
				"success": false,
				"message": "ERR INTERNAL",
				"error":   err.Error(),
			}
			httpCode = http.StatusInternalServerError
		}
		return c.Status(httpCode).JSON(payload)
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "UPDATE SUCCESS",
	})

}

func (h Handler) DeleteProductById(c *fiber.Ctx) error {
	var req = DeleteProductByIdRequest{}

	err := c.ParamsParser(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	err = h.svc.DeleteProductById(c.UserContext(), req.ID)
	if err != nil {
		var payload fiber.Map
		httpCode := 400

		switch err {
		case ErrEmptyName, ErrEmptyCategory, ErrEmptyPrice, ErrEmptyStock:
			payload = fiber.Map{
				"success": false,
				"message": "ERR BAD REQUEST",
				"error":   err.Error(),
			}
			httpCode = http.StatusBadRequest
		default:
			payload = fiber.Map{
				"success": false,
				"message": "ERR INTERNAL",
				"error":   err.Error(),
			}
			httpCode = http.StatusInternalServerError
		}
		return c.Status(httpCode).JSON(payload)
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "DELETE SUCCESS",
	})

}
