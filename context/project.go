package context

import (
	"context"
)

// ProjectKey is a context namespace for projects.
var ProjectKey Key = "project"

// ProjectFromContext returns a project defined in a context.Context.
func ProjectFromContext(ctx context.Context) string {
	project := ctx.Value(ProjectKey).(string)
	return project
}

// WithProject returns a new context with the provided project.
func WithProject(ctx context.Context, project string) context.Context {
	if ctx == nil {
		ctx = context.TODO()
	}
	return context.WithValue(ctx, ProjectKey, project)
}
