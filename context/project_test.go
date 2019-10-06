package context_test

import (
	"context"
	"testing"

	titancontext "github.com/atlaskerr/titan/context"
)

func TestProjectFromContext(t *testing.T) {
	expected := "foo"
	ctx := context.Background()
	ctx = context.WithValue(ctx, titancontext.ProjectKey, expected)
	got := titancontext.ProjectFromContext(ctx)
	if got != expected {
		t.Fail()
	}
}

func TestContextWithProject(t *testing.T) {
	expected := "foo"
	ctx := context.Background()
	ctx = titancontext.WithProject(ctx, expected)
	got := titancontext.ProjectFromContext(ctx)
	if got != expected {
		t.Fail()
	}
}

func TestContextWithProjectNilContext(t *testing.T) {
	expected := "foo"
	//lint:ignore SA1012 we want to make sure that a context is initialized if
	//no context is given.
	ctx := titancontext.WithProject(nil, expected)
	got := titancontext.ProjectFromContext(ctx)
	if got != expected {
		t.Fail()
	}
}
