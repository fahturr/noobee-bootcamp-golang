package product

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

func RegisterServiceProduct(router fiber.Router, db *gorm.DB, dbSqlx *sqlx.DB, dbNative *sql.DB) {
	// repo := NewPostgresGormRepository(db)
	repo := NewPostgresSQLXRepository(dbSqlx)

	svc := NewService(repo)
	handler := NewHandler(svc)

	var productRouter = router.Group("products")
	{
		productRouter.Post("", handler.CreateProduct)
		productRouter.Get("", handler.GetAllProduct)
		productRouter.Put("/:id", handler.UpdateProductById)
		productRouter.Delete("/:id", handler.DeleteProductById)
	}
}
