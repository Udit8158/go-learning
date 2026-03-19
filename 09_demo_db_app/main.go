package main

import (
	"fmt"

	"github.com/Udit8158/go-learning/09_demo_db_app/api"
	"github.com/Udit8158/go-learning/09_demo_db_app/db"
)

func main() {
	api.AddUserToDB(db.UserDetails{
		Name:  "udit",
		Email: "udit@email.com",
		Todos: []string{"work", "sleep"},
	})
	api.AddUserToDB(db.UserDetails{
		Name:  "jhon",
		Email: "jhon@email.com",
		Todos: []string{"sleep", "eat"},
	})
	api.AddUserToDB(db.UserDetails{
		Name:  "youy",
		Email: "jhon@email.com",
		Todos: []string{"sleep", "eat"},
	})
	fmt.Printf("%#v", db.UserTable["jhon@email.com"])
}
