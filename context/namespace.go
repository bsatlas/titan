package context

import (
	"context"
)

// NamespaceKey is a context namespace...for namespaces.
var NamespaceKey Key = "namespace"

// NamespaceFromContext returns a namespace defined in a context.Context.
func NamespaceFromContext(ctx context.Context) string {
	namespace := ctx.Value(NamespaceKey).(string)
	return namespace
}

// WithNamespace returns a new context with the provided namespace.
func WithNamespace(ctx context.Context, namespace string) context.Context {
	if ctx == nil {
		ctx = context.TODO()
	}
	return context.WithValue(ctx, NamespaceKey, namespace)
}
