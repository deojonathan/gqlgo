package core

import (
	"context"
	"fmt"
	"log"

	"github.com/dgmo/gqlgo/server/data"
	"github.com/dgmo/gqlgo/server/graph/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateProduct(ctx context.Context, cpi models.CreateProductInput) (models.CreateProductResult, error) {
	coll := data.MongoClient.GetCollection("gqlgo", "products")

	dpt := data.Product{
		ID:          uuid.New().String(),
		Name:        cpi.Name,
		Price:       cpi.Price,
		Description: cpi.Description,
		Image:       cpi.Image,
		Category:    cpi.Category,
	}

	res, err := coll.InsertOne(ctx, dpt)
	if err != nil {
		err = fmt.Errorf("Error inserting product: %v", err)
		log.Fatalf("%v", err)
		return nil, err
	}

	pt, err := Product(ctx, res.InsertedID.(string))
	if err != nil {
		err = fmt.Errorf("Error finding inserted product: %v", err)
		log.Fatalf("%v", err)
		return nil, err
	}
	return pt, nil
}

func UpdateProduct(ctx context.Context, id string, in models.UpdateProductInput) (models.UpdateProductResult, error) {
	return nil, fmt.Errorf("not implemented: UpdateProduct - updateProduct")
}

func DeleteProduct(ctx context.Context, id string) (models.DeleteProductResult, error) {
	return nil, fmt.Errorf("not implemented: DeleteProduct - deleteProduct")
}

func Product(ctx context.Context, id string) (*models.Product, error) {
	coll := data.MongoClient.GetCollection("gqlgo", "products")

	var result data.Product
	filter := bson.D{{"_id", id}}
	err := coll.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// this is weird because we just inserted it.
			log.Fatalf("Error finding successfully inserted product: %v", err)
			return nil, err
		} else {
			log.Fatalf("Error finding product: %v", err)
			return nil, err
		}
	}

	return &models.Product{
		ID:          result.ID,
		Name:        result.Name,
		Price:       result.Price,
		Description: result.Description,
		Image:       result.Image,
	}, nil
}

func Products(ctx context.Context, first *int, after *string, last *int, before *string, filter *models.ProductFilter, sorting *models.ProductSort) (*models.ProductConnection, error) {

	return nil, fmt.Errorf("not implemented: Products - products")
}
