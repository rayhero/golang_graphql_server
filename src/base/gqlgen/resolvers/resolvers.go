package resolvers

import (
	"github.com/rayhero/gqlgen-todos/src/base/gorm"
	"github.com/rayhero/gqlgen-todos/src/base/gqlgen"
)

// Resolver is a modifable struct that can be used to pass on properties used
// in the resolvers, such as DB access
type Resolver struct {
	ORM *gorm.ORM
}

// Mutation exposes mutation methods
func (r *Resolver) Mutation() gqlgen.MutationResolver {
	return &mutationResolver{r}
}

// Query exposes query methods
func (r *Resolver) Query() gqlgen.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
