package gqlsaresolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.60

import (
	"context"
)

// Ping is the resolver for the ping field.
func (r *mutationResolver) Ping(ctx context.Context) (string, error) {
	return "pong", nil
}

// Ping is the resolver for the ping field.
func (r *queryResolver) Ping(ctx context.Context) (string, error) {
	return "pong", nil
}
