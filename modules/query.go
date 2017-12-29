package modules

import (
	"github.com/graphql-go/graphql"
	"github.com/Arkangel12/curso/graphql-gorm/modules/types"
	"github.com/Arkangel12/curso/graphql-gorm/modules/resolvers"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RootQuery",
	Description: "A detailed options for querying scrous api. Feel free to navigate for each option.",
	Fields: graphql.Fields{
		"me": &graphql.Field{
			Type:        types.UserType,
			Description: "Gets the current user logged",
			Resolve:     resolvers.Me,
		},
		"user": &graphql.Field{
			Type:        types.UserType,
			Description: "Start point where we can search for one user by ID and inspect everything related to the user",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "User ID",
					Type:        graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id := params.Args["id"].(int)
				return resolvers.GetUserByID(id)
			},
		},
		"article": &graphql.Field{
			Type:        types.ArticleType,
			Description: "Retrieves a single article by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "Article ID that we are looking for",
					Type:        graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := params.Args["id"].(int)
				return resolvers.GetArticleByID(id)
			},
		},
	},
})
