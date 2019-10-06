package context

import (
	"context"
)

// RepoKey is a context namespace for repositories.
var RepoKey Key = "repository"

// RepoFromContext returns a repository defined in a context.Context.
func RepoFromContext(ctx context.Context) string {
	repo := ctx.Value(RepoKey).(string)
	return repo
}

// WithRepo returns a new context with the provided repository.
func WithRepo(ctx context.Context, repo string) context.Context {
	if ctx == nil {
		ctx = context.TODO()
	}
	return context.WithValue(ctx, RepoKey, repo)
}
