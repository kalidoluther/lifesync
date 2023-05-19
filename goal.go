package main

import (
	"github.com/graphql-go/graphql"
)

type Goal struct {
	ID          string
	Description string
	// ... and so on for the rest of your fields
}

var goalType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Goal",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		// ... and so on for the rest of your fields
	},
})

func getGoal(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(string)
	if ok {
		row := db.QueryRow("SELECT id, description FROM goals WHERE id = ?", id)
		var goal Goal
		err := row.Scan(&goal.ID, &goal.Description)
		if err != nil {
			return nil, err
		}
		return goal, nil
	}
	return nil, nil
}
