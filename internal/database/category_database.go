package database

import (
	"ecommerce-app/config"
	"ecommerce-app/internal/model"
)

func GetAllCategory() ([]model.Category, error) {
	rows, err := config.DB.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var category model.Category
		if err := rows.Scan(
			&category.CategoryID,
			&category.CategoryName,
			&category.ParentID,
			&category.Description,
		); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func InsertCategory(category *model.Category) error {
	sqlStatement := `INSERT INTO categories  (
   
                                            category_name,
                                            parent_id,
                                            description

                                        )
                    VALUES  (
                                $1, $2, $3
                            ) 
                    RETURNING category_id`
	return config.DB.QueryRow(
		sqlStatement,
		&category.CategoryName,
		&category.ParentID,
		&category.Description,
	).Scan(&category.CategoryID)
}

func UpdateCategory(id string, category *model.Category) error {
	sqlStatement := `UPDATE categories SET
  
                        category_name=$1,
                        parent_id=$2,
                        description=$3
                    
                    WHERE category_id=$4 RETURNING category_id`
	return config.DB.QueryRow(
		sqlStatement,

		&category.CategoryName,
		&category.ParentID,
		&category.Description,

		id,
	).Scan(&category.CategoryID)
}

func DeleteCategory(id string) error {
	sqlStatement := `DELETE FROM categories WHERE category_id =$1`
	_, err := config.DB.Exec(sqlStatement, id)
	return err
}
