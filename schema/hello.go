package schema

import (
	"fmt"

	"github.com/graphql-go/graphql"

	"github.com/hykuan/golesson/model"
)

var helloType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Hello",
	Description: "Hello Model",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"echo": &graphql.Field{
			Type: graphql.String,
			Description: "Echo what you enter",
			Args: graphql.FieldConfigArgument{
				"message": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// get arg value
				message, _ := p.Args["message"].(string)
				source, _ := p.Source.(model.Hello)

				return fmt.Sprintf("echo for %s with message: %s", source.Name, message), nil
			},
		},
	},
})

func HelloSchema() *graphql.Field {
	return &graphql.Field{
		Name:        "QueryHello",
		Description: "Query Hello",
		Type:        graphql.NewList(helloType),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			id, _ := p.Args["id"].(int)
			name, _ := p.Args["name"].(string)

			return (&model.Hello{}).Query(id, name)
		},
	}
}
