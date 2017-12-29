package types

import (
	"github.com/graphql-go/graphql"
	"github.com/Arkangel12/curso/graphql-gorm/models"
	"github.com/Arkangel12/curso/graphql-gorm/modules/fields"
	"github.com/Arkangel12/curso/graphql-gorm/modules/resolvers"
)

// ArticleType
var ArticleType = graphql.NewObject(graphql.ObjectConfig{
	Name:   "Article",
	Fields: fields.ArticleFields,
})

func init() {
	ArticleType.AddFieldConfig("user", &graphql.Field{
		Type: graphql.NewNonNull(UserType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			if article, ok := params.Source.(*models.Article); ok == true {
				return resolvers.GetUserByID(article.UserID)
			}
			return nil, nil
		},
	})
}
