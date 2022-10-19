package repositories

import (
	"database/sql"
	"fmt"
	"github.com/sanzharanarbay/golang_gRPC_crud_eample/internal/models"
	"log"
	"time"
)

type CategoryRepository struct {
	dbClient *sql.DB
}

func NewCategoryRepository(dbClient *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		dbClient: dbClient,
	}
}

type CategoryRepositoryInterface interface {
	GetCategoryById(ID int) (*models.Category, error)
	GetAllCategories() ([]*models.Category, error)
	SaveCategory(*models.Category) (bool, error)
	DeleteCategory(ID int) (bool, error)
	UpdateCategory(*models.Category) (bool, error)
}

func (c *CategoryRepository) GetCategoryById(ID int) (*models.Category, error) {
	var category models.Category
	err := c.dbClient.QueryRow(`SELECT * FROM categories WHERE id=$1`, ID).Scan(&category.ID, &category.Name, &category.Keyword,
		&category.CreatedAt, &category.UpdatedAt)
	switch err {
	case sql.ErrNoRows:
		log.Printf("No rows were returned!")
		return nil, nil
	case nil:
		return &category, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}
	return &category, nil
}

func (c *CategoryRepository) GetAllCategories() (*[]models.Category, error) {
	rows, err := c.dbClient.Query("SELECT * FROM categories")
	if err != nil {
		fmt.Printf("ERROR SELECT QUERY - %s", err)
		return nil, err
	}
	var catList []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.ID, &category.Name, &category.Keyword,
			&category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			fmt.Printf("ERROR QUERY SCAN - %s", err)
			return nil, err
		}
		catList = append(catList, category)
	}
	return &catList, nil
}

func (c *CategoryRepository) SaveCategory(category *models.Category) (bool, error) {
	category.CreatedAt = time.Now().Local()
	category.UpdatedAt = time.Now().Local()

	sqlStatement := `INSERT into categories (name, keyword, created_at, updated_at) VALUES ($1, $2, $3, $4)`
	_, err := c.dbClient.Exec(sqlStatement, category.Name, category.Keyword, category.CreatedAt, category.UpdatedAt)
	if err != nil {
		panic(err)
	}
	return true, nil
}

func (c *CategoryRepository) UpdateCategory(category *models.Category, catID int) (bool, error) {
	category.UpdatedAt = time.Now().Local()
	sqlStatement := `UPDATE categories SET name=$1, keyword=$2, updated_at=$3 WHERE id=$4`
	_, err := c.dbClient.Exec(sqlStatement, category.Name, category.Keyword, category.UpdatedAt, catID)
	if err != nil {
		panic(err)
	}
	return true, nil
}

func (c *CategoryRepository) DeleteCategory(ID int) (bool, error) {
	_, err := c.dbClient.Exec(`DELETE FROM categories WHERE id=$1`, ID)
	if err != nil {
		panic(err)
	}

	return true, nil
}







