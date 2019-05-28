package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/hykuan/golesson/model"
	"github.com/jinzhu/gorm"
)

var commentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"body": &graphql.Field{
				Type: graphql.String,
			},
			"author_id": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

func CreateCommentMutation() *graphql.Field {
	return &graphql.Field{
		Type:        commentType,
		Description: "Create a new Comment",
		Args: graphql.FieldConfigArgument{
			"body": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"tutorial_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"author_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			comment := model.Comment{
				Body: params.Args["body"].(string),
				TutorialID: params.Args["tutorial_id"].(int),
				AuthorID: params.Args["author_id"].(int),
			}
			db, _ := gorm.Open("sqlite3", "sqlite.db")
			db.Save(&comment)
			return comment, nil
		},
	}
}