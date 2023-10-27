package product

type CreateProductRequest struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
}

type UpdateProductByIdRequest struct {
	ID       int    `params:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
}

type DeleteProductByIdRequest struct {
	ID int `params:"id"`
}
