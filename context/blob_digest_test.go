package context_test

import (
	"context"
	"testing"

	titancontext "github.com/atlaskerr/titan/context"
)

func TestBlobDigestFromContext(t *testing.T) {
	expected := "foo"
	ctx := context.Background()
	ctx = context.WithValue(ctx, titancontext.BlobDigestKey, expected)
	got := titancontext.BlobDigestFromContext(ctx)
	if got != expected {
		t.Fail()
	}
}

func TestContextWithBlobDigest(t *testing.T) {
	expected := "foo"
	ctx := context.Background()
	ctx = titancontext.WithBlobDigest(ctx, expected)
	got := titancontext.BlobDigestFromContext(ctx)
	if got != expected {
		t.Fail()
	}
}

func TestContextWithBlobDigestNilContext(t *testing.T) {
	expected := "foo"
	//lint:ignore SA1012 we want to make sure that a context is initialized if
	//no context is given.
	ctx := titancontext.WithBlobDigest(nil, expected)
	got := titancontext.BlobDigestFromContext(ctx)
	if got != expected {
		t.Fail()
	}
}
