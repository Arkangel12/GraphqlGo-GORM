package modules

import (
	"github.com/graphql-go/graphql"
	"github.com/Arkangel12/curso/graphql-gorm/modules/resolvers"
	"github.com/Arkangel12/curso/graphql-gorm/modules/types"
)

// RootMutation for the api
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RootMutation",
	Description: "Root scrous mutations where we can find a detail info about the actions that we can do",
	Fields: graphql.Fields{
		"createUser": &graphql.Field{
			Type:        types.UserType,
			Description: "Creates a new user account",
			Args: graphql.FieldConfigArgument{
				"username": &graphql.ArgumentConfig{
					Description: "New username for user",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"email": &graphql.ArgumentConfig{
					Description: "New email for user",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Description: "New password for user",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"full_name": &graphql.ArgumentConfig{
					Description: "User Fullname",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: resolvers.CreateUser,
		},
		"createArticle": &graphql.Field{
			Type:        types.ArticleType,
			Description: "Creates a new article",
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Description: "New article title",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"content": &graphql.ArgumentConfig{
					Description: "New article body",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"id": &graphql.ArgumentConfig{
					Description: "Current user id",
					Type:        graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: resolvers.CreateArticle,
		},
		"authenticateUser": &graphql.Field{
			Type:        types.UserType,
			Description: "User authentication",
			Args: graphql.FieldConfigArgument{
				"username": &graphql.ArgumentConfig{
					Description: "Username to log in",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Description: "Password to log in",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: resolvers.AuthenticateUser,
		},
		"passwordRecovery": &graphql.Field{
			Type:        types.UserType,
			Description: "Password recovery link",
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Description: "Email to send recovery link",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: resolvers.PasswordRecovery,
		},
	},
})
