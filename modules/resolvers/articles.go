package resolvers

import (
	"github.com/graphql-go/graphql"
	"github.com/Arkangel12/curso/graphql-gorm/db"
	"github.com/Arkangel12/curso/graphql-gorm/models"
)

// CreateArticle create new article
func CreateArticle(params graphql.ResolveParams) (interface{}, error) {
	title, _ := params.Args["title"].(string)
	content, _ := params.Args["content"].(string)
	id, _ := params.Args["id"].(int)

	article := models.Article{
		Title:   title,
		Content: content,
		UserID:  id,
	}
	db := db.GetConnection()

	// Create new article
	db.Create(&article)

	return &article, nil

}

// GetArticleByID create new article
func GetArticleByID(id int) (interface{}, error) {

	article := models.Article{}

	db := db.GetConnection()

	// Create new article
	db.Where("id = ?", id).First(&article)
	return &article, nil

}

// GetArticleByIDAndUser ...
func GetArticleByIDAndUser(id, user int) (interface{}, error) {

	article := models.Article{}

	db := db.GetConnection()

	db.Where("id = ? AND user_id = ?", id, user).First(&article)

	return &article, nil
}

// GetArticlesForUser ...
func GetArticlesForUser(id int) ([]*models.Article, error) {

	var article []*models.Article

	db := db.GetConnection()

	db.Where("user_id = ?", id).Find(&article)

	return article, nil
}
