package graph

//go:generate go run github.com/99designs/gqlgen

import "github.com/avi94/newapp/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos   []*model.Todo
	authors []*model.Author
	posts   []*model.Post
}
