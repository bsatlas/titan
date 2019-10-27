package context_test

import (
	"context"
	"testing"

	titancontext "github.com/atlaskerr/titan/context"
)

func TestBlobUploadIDFromContext(t *testing.T) {
	expected := "foo"
	ctx := context.Background()
	ctx = context.WithValue(ctx, titancontext.BlobUploadIDKey, expected)
	got := titancontext.BlobUploadIDFromContext(ctx)
	if got != expected {
		t.Fail()
	}
}

func TestContextWithBlobUploadID(t *testing.T) {
	expected := "foo"
	ctx := context.Background()
	ctx = titancontext.WithBlobUploadID(ctx, expected)
	got := titancontext.BlobUploadIDFromContext(ctx)
	if got != expected {
		t.Fail()
	}
}

func TestContextWithBlobUploadIDNilContext(t *testing.T) {
	expected := "foo"
	//lint:ignore SA1012 we want to make sure that a context is initialized if
	//no context is given.
	ctx := titancontext.WithBlobUploadID(nil, expected)
	got := titancontext.BlobUploadIDFromContext(ctx)
	if got != expected {
		t.Fail()
	}
}
