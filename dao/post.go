package dao

import (
	"GoBlog/models"
	"log"
)

// 统计博客数量
func CountGetAllPost() (int, error) {
	rows, err := DB.Query("select count(*) from blog_post ")
	if err != nil {
		log.Println(err)
		return -1, err
	}
	rows.Next()
	var postCnt int
	_ = rows.Scan(&postCnt)
	return postCnt, nil
}

func GetPostInfo(page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post limit ?,?", page, pageSize)
	if err != nil {
		return nil, err
	}

	posts := make([]models.Post, 0)
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
