package context_test

import (
	"context"
	"testing"

	titancontext "github.com/atlaskerr/titan/context"
)

func TestRepoFromContext(t *testing.T) {
	expected := "foo"
	ctx := context.Background()
	ctx = context.WithValue(ctx, titancontext.RepoKey, expected)
	got := titancontext.RepoFromContext(ctx)
	if got != expected {
		t.Fail()
	}
}

func TestContextWithRepo(t *testing.T) {
	expected := "foo"
	ctx := context.Background()
	ctx = titancontext.WithRepo(ctx, expected)
	got := titancontext.RepoFromContext(ctx)
	if got != expected {
		t.Fail()
	}
}

func TestContextWithRepoNilContext(t *testing.T) {
	expected := "foo"
	//lint:ignore SA1012 we want to make sure that a context is initialized if
	//no context is given.
	ctx := titancontext.WithRepo(nil, expected)
	got := titancontext.RepoFromContext(ctx)
	if got != expected {
		t.Fail()
	}
}
