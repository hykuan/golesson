package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/hykuan/golesson/model"
	"github.com/jinzhu/gorm"
)

var authorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"tutorials": &graphql.Field{
				Type: graphql.NewList(tutorialType),
			},
		},
	},
)

func SingleAuthorSchema() *graphql.Field {
	return &graphql.Field{
		Type:        authorType,
		Description: "Get Author By ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var author model.Author
			db, _ := gorm.Open("sqlite3", "sqlite.db")
			db.Debug().Set("gorm:auto_preload", true).First(&author, params.Args["id"].(int))
			return author, nil
		},
	}
}

func ListAuthorSchema() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(authorType),
		Description: "Get Author List",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var authors []model.Author
			db, _ := gorm.Open("sqlite3", "sqlite.db")
			db.Set("gorm:auto_preload", true).Find(&authors)
			return authors, nil
		},
	}
}

func CreateAuthorMutation() *graphql.Field {
	return &graphql.Field{
		Type:        authorType,
		Description: "Create a new author",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			author := model.Author{Name: params.Args["name"].(string)}
			db, _ := gorm.Open("sqlite3", "sqlite.db")
			err := db.Debug().Save(&author).Error
			return author, err
		},
	}
}
