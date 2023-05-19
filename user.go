package main

import (
	"github.com/graphql-go/graphql"
)

type User struct {
	ID    string
	Name  string
	Email string
	//Birthday *time.Time
	//Gender   string
	//Location    *Location
	//Timezone string
	//Job      string
	//Medications []*Medication
	// ... and so on for the rest of your fields
}

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
	},
})

func getUser(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(string)
	if ok {
		row := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)
		var user User
		err := row.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, nil
}
