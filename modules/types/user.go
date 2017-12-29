package types

import (
	"github.com/graphql-go/graphql"
	"github.com/Arkangel12/curso/graphql-gorm/models"
	"github.com/Arkangel12/curso/graphql-gorm/modules/resolvers"
	"github.com/Arkangel12/curso/graphql-gorm/modules/fields"
)

// UserType
var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:   "User",
		Fields: fields.UserFields,
	},
)

func init() {
	UserType.AddFieldConfig("article", &graphql.Field{
		Type: ArticleType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Article ID",
				Type:        graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			if user, ok := params.Source.(*models.User); ok == true {
				id := params.Args["id"].(int)
				return resolvers.GetArticleByIDAndUser(id, user.ID)
			}
			return nil, nil
		},
	})
	UserType.AddFieldConfig("articles", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(ArticleType))),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			if user, ok := params.Source.(*models.User); ok == true {
				return resolvers.GetArticlesForUser(user.ID)
			}
			return nil, nil
		},
	})
}
