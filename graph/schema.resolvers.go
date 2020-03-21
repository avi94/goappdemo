package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"github.com/avi94/newapp/graph/generated"
	"github.com/avi94/newapp/graph/model"
)

func (r *authorResolver) Posts(ctx context.Context, obj *model.Author) ([]*model.Post, error) {
	var posts []*model.Post
	for _, post := range r.posts {
		if post.AuthorID == obj.ID {
			posts = append(posts, post)
		}
	}
	return posts, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		ID:     fmt.Sprintf("T%d", rand.Int()),
		Text:   input.Text,
		Done:   false,
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) CreateAuthor(ctx context.Context, input model.NewAuthor) (*model.Author, error) {
	author := &model.Author{
		ID:    fmt.Sprintf("A%d", rand.Int()),
		Name:  input.Name,
		Age:   input.Age,
		Posts: nil,
	}
	r.authors = append(r.authors, author)
	return author, nil
}

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	post := &model.Post{
		ID:       fmt.Sprintf("P%d", rand.Int()),
		Title:    input.Title,
		AuthorID: input.AuthorID,
	}
	r.posts = append(r.posts, post)
	return post, nil
}

func (r *postResolver) Author(ctx context.Context, obj *model.Post) (*model.Author, error) {
	author := new(model.Author)
	for _, a := range r.authors {
		if a.ID == obj.AuthorID {
			author = a
			break
		}
	}
	if author == nil {
		return nil, errors.New("Author with id not found.")
	}
	return author, nil
	//return &model.Author{
	//	ID:    obj.AuthorID,
	//	Name:  "author " + obj.AuthorID,
	//	Age:   0,
	//	Posts: nil,
	//}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	return r.authors, nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	return r.posts, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{
		ID:   obj.UserID,
		Name: "user " + obj.UserID,
	}, nil
}

// Author returns generated.AuthorResolver implementation.
func (r *Resolver) Author() generated.AuthorResolver { return &authorResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type authorResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
