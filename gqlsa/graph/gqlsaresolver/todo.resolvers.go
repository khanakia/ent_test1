package gqlsaresolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.48

import (
	"context"
	"gqlsa/graph/generated"
	"saas/gen/ent"
	"saas/pkg/middleware/adminauthmiddleware"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input ent.CreateTodoInput) (*ent.Todo, error) {
	cuser, err := adminauthmiddleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}

	return r.Plugin.EntDB.Client().Todo.Create().SetInput(input).Save(ctx)
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, id string, input ent.UpdateTodoInput) (*ent.Todo, error) {
	cuser, err := adminauthmiddleware.GetUserFromGqlCtx(ctx)
	if cuser == nil {
		return nil, err
	}

	return r.Plugin.EntDB.Client().Todo.UpdateOneID(id).SetInput(input).Save(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
