package schema

import (
	"github.com/graphql-go/graphql"
)

var aggregateSchema = graphql.Fields{
	"hello":     HelloSchema(),
	"author":    SingleAuthorSchema(),
	"authors":   ListAuthorSchema(),
	"tutorial":  SingleTutorialSchema(),
	"tutorials": ListTutorialSchema(),
}

var aggregateMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"create_author":   CreateAuthorMutation(),
		"create_tutorial": CreateTutorialMutation(),
		"create_comment": CreateCommentMutation(),
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RootQuery",
	Description: "Root Query",
	Fields:      aggregateSchema,
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: aggregateMutations,
})
