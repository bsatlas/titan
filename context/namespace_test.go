package context_test

import (
	"context"
	"testing"

	titancontext "github.com/atlaskerr/titan/context"
)

func TestNamespaceFromContext(t *testing.T) {
	expected := "foo"
	ctx := context.Background()
	ctx = context.WithValue(ctx, titancontext.NamespaceKey, expected)
	got := titancontext.NamespaceFromContext(ctx)
	if got != expected {
		t.Fail()
	}
}

func TestContextWithNamespace(t *testing.T) {
	expected := "foo"
	ctx := context.Background()
	ctx = titancontext.WithNamespace(ctx, expected)
	got := titancontext.NamespaceFromContext(ctx)
	if got != expected {
		t.Fail()
	}
}

func TestContextWithNamespaceNilContext(t *testing.T) {
	expected := "foo"
	//lint:ignore SA1012 we want to make sure that a context is initialized if
	//no context is given.
	ctx := titancontext.WithNamespace(nil, expected)
	got := titancontext.NamespaceFromContext(ctx)
	if got != expected {
		t.Fail()
	}
}
