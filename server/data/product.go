package data

type Product struct {
	ID          string  `bson:"_id"`
	Name        string  `bson:"name"`
	Price       float64 `bson:"price"`
	Description string  `bson:"description"`
	Image       string  `bson:"image"`
	Category    string  `bson:"category"`
}
