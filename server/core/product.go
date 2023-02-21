package core

import (
	"context"
	"fmt"
	// "log"

	"github.com/dgllrmo/gqlgo/server/graph/models"

	// "go.mongodb.org/mongo-driver/mongo"
    // "go.mongodb.org/mongo-driver/mongo/options"
    // "go.mongodb.org/mongo-driver/mongo/readpref"
)

func CreateProduct(ctx context.Context, in models.CreateProductInput) (models.CreateProductResult, error) {
	return nil, fmt.Errorf("not implemented: CreateProduct - createProduct")
}


func UpdateProduct(ctx context.Context, id string, in models.UpdateProductInput) (models.UpdateProductResult, error) {
	return nil, fmt.Errorf("not implemented: UpdateProduct - updateProduct")
}

func DeleteProduct(ctx context.Context, id string) (models.DeleteProductResult, error) {
	return nil, fmt.Errorf("not implemented: DeleteProduct - deleteProduct")
}

func Product(ctx context.Context, id string) (*models.Product, error) {
	return nil, fmt.Errorf("not implemented: Product - product")
}

func Products(ctx context.Context, first *int, after *string, last *int, before *string, filter *models.ProductFilter, sorting *models.ProductSort) (*models.ProductConnection, error) {
	return nil, fmt.Errorf("not implemented: Products - products")
}
