package resolvers

import (
	"context"

	"github.com/rayhero/gqlgen-todos/src/base/gqlgen"
	"github.com/rayhero/gqlgen-todos/src/base/gqlgen/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() gqlgen.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() gqlgen.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input models.UserInput) (*models.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context, id *string) (*models.Users, error) {
	panic("not implemented")
}
