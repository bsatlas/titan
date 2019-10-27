package context

import (
	"context"
)

// BlobDigestKey is a context namespace for OCI blob digests.
var BlobDigestKey Key = "blob-digest"

// BlobDigestFromContext returns a blob digest defined in a context.Context.
func BlobDigestFromContext(ctx context.Context) string {
	blobDigest := ctx.Value(BlobDigestKey).(string)
	return blobDigest
}

// WithBlobDigest returns a new context with the provided blob digest.
func WithBlobDigest(ctx context.Context, digest string) context.Context {
	if ctx == nil {
		ctx = context.TODO()
	}
	return context.WithValue(ctx, BlobDigestKey, digest)
}
