package fields

import (
	"github.com/graphql-go/graphql"
	"github.com/Arkangel12/curso/graphql-gorm/models"
)

// ArticleFields ...
var ArticleFields = graphql.Fields{
	"id": &graphql.Field{
		Type: graphql.NewNonNull(graphql.Int),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			if article, ok := params.Source.(*models.Article); ok == true {
				return article.ID, nil
			}
			return nil, nil
		},
	},
	"title": &graphql.Field{
		Type: graphql.NewNonNull(graphql.String),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			if article, ok := params.Source.(*models.Article); ok == true {
				return article.Title, nil
			}
			return nil, nil
		},
	},
	"content": &graphql.Field{
		Type: graphql.NewNonNull(graphql.String),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			if article, ok := params.Source.(*models.Article); ok == true {
				return article.Content, nil
			}
			return nil, nil
		},
	},
}
