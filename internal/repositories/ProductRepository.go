package repositories

import (
	"database/sql"
	"fmt"
	"github.com/sanzharanarbay/golang_gRPC_crud_eample/internal/models"
	"log"
	"time"
)

type ProductRepository struct {
	dbClient *sql.DB
}

func NewProductRepository(dbClient *sql.DB) *ProductRepository {
	return &ProductRepository{
		dbClient: dbClient,
	}
}

type ProductRepositoryInterface interface {
	GetProductById(ID int) (*models.Product, error)
	GetAllCategories() ([]*models.Product, error)
	SaveProduct(*models.Product) (bool, error)
	DeleteProduct(ID int) (bool, error)
	UpdateProduct(*models.Product) (bool, error)
}


func (p *ProductRepository) GetProductById(ID int) (*models.Product, error) {
	var product models.Product
	err := p.dbClient.QueryRow(`SELECT * FROM products WHERE id=$1`, ID).Scan(&product.ID, &product.Name, &product.Description,
		&product.CategoryId, &product.Price, &product.CreatedAt, &product.UpdatedAt)
	switch err {
	case sql.ErrNoRows:
		log.Printf("No rows were returned!")
		return nil, nil
	case nil:
		return &product, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}
	return &product, nil
}

func (p *ProductRepository) GetAllProducts() (*[]models.Product, error) {
	rows, err := p.dbClient.Query("SELECT * FROM products")
	if err != nil {
		fmt.Printf("ERROR SELECT QUERY - %s", err)
		return nil, err
	}
	var productList []models.Product
	for rows.Next() {
		var product models.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Description,
			&product.CategoryId, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			fmt.Printf("ERROR QUERY SCAN - %s", err)
			return nil, err
		}
		productList = append(productList, product)
	}
	return &productList, nil
}

func (p *ProductRepository) SaveProduct(product *models.Product) (bool, error) {
	product.CreatedAt = time.Now().Local()
	product.UpdatedAt = time.Now().Local()

	sqlStatement := `INSERT into products (name, description, category_id, price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := p.dbClient.Exec(sqlStatement, product.Name, product.Description, product.CategoryId, product.Price, product.CreatedAt, product.UpdatedAt)
	if err != nil {
		panic(err)
	}
	return true, nil
}

func (p *ProductRepository) UpdateProduct(product *models.Product, catID int) (bool, error) {
	product.UpdatedAt = time.Now().Local()
	sqlStatement := `UPDATE products SET name=$1, description=$2, category_id=$3, price=$4, updated_at=$5 WHERE id=$6`
	_, err := p.dbClient.Exec(sqlStatement, product.Name, product.Description, product.CategoryId, product.Price, product.UpdatedAt, catID)
	if err != nil {
		panic(err)
	}
	return true, nil
}

func (p *ProductRepository) DeleteProduct(ID int) (bool, error) {
	_, err := p.dbClient.Exec(`DELETE FROM products WHERE id=$1`, ID)
	if err != nil {
		panic(err)
	}

	return true, nil
}