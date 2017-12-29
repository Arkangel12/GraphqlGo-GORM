package fields

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/Arkangel12/curso/graphql-gorm/models"
)

// UserFields group all user information
var UserFields = graphql.Fields{
	"id": &graphql.Field{
		Type: graphql.NewNonNull(graphql.ID),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			// log.Println(params)
			user, ok := params.Source.(*models.User)

			if !ok {
				return nil, errors.New("No user received by user resolver")
			}
			return user.ID, nil
		},
	},
	"username": &graphql.Field{
		Type: graphql.NewNonNull(graphql.String),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			if user, ok := params.Source.(*models.User); ok == true {
				return user.Username, nil
			}
			return nil, nil
		},
	},
	"email": &graphql.Field{
		Type: graphql.NewNonNull(graphql.String),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			if user, ok := params.Source.(*models.User); ok == true {
				return user.Email, nil
			}
			return nil, nil
		},
	},
	"full_name": &graphql.Field{
		Type: graphql.NewNonNull(graphql.String),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			if user, ok := params.Source.(*models.User); ok == true {
				return user.FullName, nil
			}
			return nil, nil
		},
	},
}
