package schema

import (
	"github.com/Arkangel12/curso/graphql-gorm/modules"
	"github.com/graphql-go/graphql"
)

// Schema for graphql
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    modules.RootQuery,
	Mutation: modules.RootMutation,
})
