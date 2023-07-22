package dao

import (
	"GoBlog/models"
	"log"
)

func GetAllCategory() ([]models.Category, error) {
	rows, err := DB.Query("select * from blog_category")
	if err != nil {
		log.Println("GetAllCategory查询出错:", err)
		return nil, err
	}
	categories := make([]models.Category, 0)
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Println("GetAllCategory取值出错:", err)
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
