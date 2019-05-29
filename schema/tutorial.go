package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/hykuan/golesson/model"
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

var tutorialType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Tutorial",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"author_id": &graphql.Field{
				Type: graphql.Int,
			},
			"comments": &graphql.Field{
				Type: graphql.NewList(commentType),
			},
		},
	},
)

func init() {
	tutorialType.AddFieldConfig("author", &graphql.Field{
		Type: authorType,
	})
}

func SingleTutorialSchema() *graphql.Field {
	return &graphql.Field{
		Type:        tutorialType,
		Description: "Get Tutorial By ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var tutorial model.Tutorial
			db, _ := gorm.Open("sqlite3", "sqlite.db")
			db.Set("gorm:auto_preload", true).First(&tutorial, params.Args["id"].(int))
			return tutorial, nil
		},
	}
}

func ListTutorialSchema() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(tutorialType),
		Description: "Get Tutorial List",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var tutorials []model.Tutorial
			db, _ := gorm.Open("sqlite3", "sqlite.db")
			db.Set("gorm:auto_preload", true).Find(&tutorials)
			return tutorials, nil
		},
	}
}

func CreateTutorialMutation() *graphql.Field {
	return &graphql.Field{
		Type:        tutorialType,
		Description: "Create a new Tutorial",
		Args: graphql.FieldConfigArgument{
			"title": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"author_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			tutorial := model.Tutorial{Title: params.Args["title"].(string), AuthorID: params.Args["author_id"].(int)}
			db, _ := gorm.Open("sqlite3", "sqlite.db")
			db.Save(&tutorial)
			return tutorial, nil
		},
	}
}